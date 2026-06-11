# abloqd 운영 문서 (Phase019·020)

대상: `deploy/backend/docker-compose.yaml` 한 장으로 셀프호스트되는 abloqd
(postgres + abloqd + cron 프로필 9종). 시크릿 템플릿은
`deploy/backend/.env.example`, 사이트 선언 SSOT 템플릿은
`deploy/backend/sites.yaml.example`, 통합 리허설 증거는 `docs/rehearsal/`.

**v0.2.0 멀티사이트**: abloqd 한 인스턴스가 사이트 N개를 운용한다. 사이트
목록은 deploy 측 `sites.yaml`(env `SITES_YAML_PATH`)이 SSOT — 기동 시
strict 파싱·검증해 `sites` 테이블에 upsert하고, 파일에서 사라진 사이트는
`active=false`(행 삭제 없음). 도메인 API는 전부 `/sites/<name>/…` 하위로
이동했고, 글로벌로 남는 것은 `POST /auth/login`·`GET /sites`·`/health`·
`/metrics` 뿐이다. 운용 절차의 `<s>`는 사이트 이름(슬러그)이다.

## 1. 구성 요약

| 서비스 | 프로필 | 권고 주기 | 호출 |
|---|---|---|---|
| `archiver-backstop` | `backstop` | 15분 | login → 사이트 순회 → `POST /sites/<s>/receipts/retry` → `POST /sites/<s>/archive/process` |
| `crawl-ingest` | `crawl` | 일간 | login → 사이트 순회 → `POST /sites/<s>/ingest/crawl` |
| `gsc-ingest` | `gsc` | 일간 | login → 사이트 순회 → `POST /sites/<s>/ingest/gsc` (inspect:false) |
| `citation-sample` | `citation` | 주간 | login → 사이트 순회 → `POST /sites/<s>/sample/citations` |
| `report-monthly` | `report` | 월간 | login → 사이트 순회 → `POST /sites/<s>/reports/monthly` (ym:"") |
| `freshness-scan` | `freshness` | 월간 | login → 사이트 순회 → `POST /sites/<s>/scans/freshness` (ym:"") |
| `evidence-scan` | `evidence` | 분기 | login → 사이트 순회 → `POST /sites/<s>/scans/evidence` |
| `cluster-scan` | `cluster` | 분기 | login → 사이트 순회 → `POST /sites/<s>/scans/cluster` |
| `queue-export` | `queue` | 주간 | login → 사이트 순회 → `POST /sites/<s>/queue/export` |

전 cron 공통 규약: **시크릿은 operator 자격증명(토큰 아님)** — access token
TTL 15분이라 매 주기 login으로 발급받는다. login rate limit는 5회/분·IP.
**모든 호출은 멱등** — 실패한 주기는 다음 주기가 흡수한다.
**사이트 순회는 셸이 한다**: 매 주기 `GET /sites?active_filter=true`로
목록을 떠서 사이트별 op를 순차 호출하고, 한 사이트의 실패는 로그만 남기고
다음 사이트로 넘어간다 (fan-out은 결정 로직이 아니므로 셸 소관).
operator 비밀번호에 `"`·`\`는 금지(JSON body에 그대로 끼워 넣는 패턴 —
`backend/scripts/compose-cron-smoke`가 셸 전개를 검증한다).

배포 직후 훅(CI → login → `POST /sites/<s>/hooks/deployed` →
`POST /sites/<s>/archive/process`)은 cron이 아니라 각 블로그 저장소 CI에
산다 — `template/files/deploy/archiver.md` (`ABLOQD_SITE` 시크릿 추가).

## 2. 모니터링 최소선

- **`GET /health`** — 인증 없음. 외부 모니터(uptime 체커 등)는 이것 하나를
  1분 주기로 본다. 비정상이면 아래 순서로 본다:
  1. `docker compose -f deploy/backend/docker-compose.yaml ps` — 컨테이너 상태
  2. `docker compose logs abloqd --tail 100` — 기동 실패는 대부분 env 누락
     (`.env.example`의 필수 3종) 또는 postgres 미기동
  3. `docker compose logs postgres --tail 50`
- cron 동작 확인은 로그로: 각 서비스가 실패 시 `<이름>: ... failed` 한 줄을
  남긴다. `docker compose logs --since 48h | grep failed`가 0줄이면 정상.
- 데이터 신선도 점검(주간 권고, operator 토큰 필요):
  사이트별 `GET /sites/<s>/receipts?status=failed` 0행, `GET /sites/<s>/queue?status=open`이 계속 쌓이기만
  하지 않는지(아래 §3.3), `GET /sites/<s>/reports/monthly/<직전월>` 존재. 사이트 목록은 `GET /sites?active_filter=true`.

## 3. 장애 시나리오별 대응

### 3.1 cron 실패 (login failed / 호출 failed 로그)

- `login failed` 반복: operator 자격증명 불일치 또는 rate limit. `.env`의
  `ABLOQD_OPERATOR_EMAIL/PASSWORD` 확인(비밀번호에 `"`·`\` 금지), 같은 IP에서
  분당 5회를 넘는 login이 없는지 확인.
- `... failed (next cycle retries)`: 호출 자체는 멱등이라 다음 주기가
  흡수한다. 즉시 복구가 필요하면 운영자가 같은 endpoint를 손으로 1회
  호출하면 된다(컨테이너 재시작 불요):
  ```bash
  TOKEN=$(curl -sf -X POST "$ABLOQD_URL/auth/login" -H 'Content-Type: application/json' \
    -d "{\"email\":\"$ABLOQD_OPERATOR_EMAIL\",\"password\":\"$ABLOQD_OPERATOR_PASSWORD\"}" | jq -r .access_token)
  curl -sf -m 600 -X POST "$ABLOQD_URL/sites/<s>/<endpoint>" -H "Authorization: Bearer $TOKEN" \
    -H 'Content-Type: application/json' -d '<해당 cron의 body>'
  ```
- 멀티사이트에서 한 사이트만 계속 실패하면(`<이름>: <사이트> ... failed`
  로그) 그 사이트의 행 값(repo_path 마운트, queue_export.repo_url,
  cf_log_source)을 먼저 본다 — 다른 사이트는 영향 없이 계속 돈다.
- 주기를 당기고 싶으면 `.env`의 `*_PERIOD_SECONDS`를 낮추고
  `docker compose up -d` (해당 서비스만 재생성된다).

### 3.2 영수증 failed → retry

아카이버 영수증(`wayback`/`indexnow`/`gsc_index`)이 `failed`로 남는 경우
(외부 API 5xx, 자격증명 누락 등):

```bash
# ① 원인 확인 — kind별 failed 행과 마지막 에러
curl -sf "$ABLOQD_URL/sites/$SITE/receipts?status=failed" -H "Authorization: Bearer $TOKEN"
# ② 재무장: failed/deferred → pending (deploy_id·kind_filter로 좁힐 수 있다, ""=전체)
curl -sf -X POST "$ABLOQD_URL/sites/$SITE/receipts/retry" -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' -d '{"deploy_id":"","kind_filter":""}'
# ③ 실행: pending 일괄 처리 (멱등, limit로 쿼터 보호)
curl -sf -m 600 -X POST "$ABLOQD_URL/sites/$SITE/archive/process" -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' -d '{"limit":200}'
```

`archiver-backstop` 프로필이 켜져 있으면 같은 ②→③을 15분마다 자동 수행한다.
**retry 없이 process만 돌리면 failed/deferred는 영원히 남는다** — 순서 필수.
자격증명 누락이 원인이면(`INDEXNOW_KEY`/`GSC_SA_JSON_FILE`/`WAYBACK_*` 빈 값)
`.env`를 채우고 `docker compose up -d abloqd` 후 ②→③.

### 3.3 큐 적체 → export consumed 동기화

증상: `GET /sites/<s>/queue?status=open` 행이 줄지 않거나, 에이전트가 큐 파일을
지웠는데 `status=exported`가 그대로다.

대응: **`POST /sites/<s>/queue/export` 재호출이 동기화 그 자체다** — 전용 endpoint는
없다. 한 사이클이 작업 클론 pull → 삭제된 큐 파일 감지(exported→consumed)
→ open 행 신규 발급 → push를 전부 수행하며 멱등이다(변경 없으면 no-op).

```bash
curl -sf -m 600 -X POST "$ABLOQD_URL/sites/$SITE/queue/export" -H "Authorization: Bearer $TOKEN"
# → {"consumed":N,"exported":M}
```

- export가 500이면: 사이트 행의 `queue_export.repo_url` 미설정(단일 사이트 env
  배포면 `QUEUE_EXPORT_REPO_URL`)이거나 deploy key로 push
  불가(`QUEUE_DEPLOY_KEY_FILE`/`QUEUE_KNOWN_HOSTS_FILE` 확인).
  `docker compose logs abloqd | grep -i export`.
- open이 계속 쌓이기만 하면 소비 측(에이전트) 문제다 — 큐 자체는 정상.
  발급량을 줄이려면 스캔 주기를 늘린다(적체는 장애가 아니라 운용 신호 —
  설계서 §5).

### 3.4 커서 리셋 (재집계) — 구간 DELETE와 반드시 한 동작

크롤 수집 커서(`ingest_cursors`, source=`cf_logs`)는 닫힌 시간대 경계
(`YYYY-MM-DD-HH`)다. 어떤 구간을 다시 집계해야 하면 **커서 후퇴와 해당
구간 crawl_hits DELETE를 반드시 한 트랜잭션으로** 실행한다 — 커서만 되돌리면
같은 로그가 이중 누적된다 (`backend/specs/db/ingest_cursors.sql` 주석).

```sql
-- 예: 사이트 <name>의 2026-06-05 00시 이후를 재집계 (psql, 수동 — API 없음).
-- 멀티사이트: 커서·행 모두 site_id 스코프 — 반드시 같은 사이트로 한정한다.
BEGIN;
DELETE FROM crawl_hits
 WHERE site_id = (SELECT id FROM sites WHERE name = '<name>')
   AND hit_date >= '2026-06-05';
UPDATE ingest_cursors SET cursor_hour = '2026-06-04-23', updated_at = NOW()
 WHERE site_id = (SELECT id FROM sites WHERE name = '<name>')
   AND source = 'cf_logs';
COMMIT;
-- 이후 POST /sites/<name>/ingest/crawl (또는 다음 crawl-ingest 주기)이 구간을 다시 채운다
```

GSC 커서는 별도 테이블이 아니라 `MAX(gsc_snapshots.snap_date)`에서 유도된다
— 재집계는 구간 DELETE만 하면 커서가 자동 후퇴한다:

```sql
DELETE FROM gsc_snapshots
 WHERE site_id = (SELECT id FROM sites WHERE name = '<name>')
   AND snap_date >= '2026-06-05';
-- 이후 POST /sites/<name>/ingest/gsc — GSC_LOOKBACK_DAYS 한도 안에서 다시 채운다
```

CF 배달 지연(드물게 24h)이 의심되면 DELETE 대신 `CF_LOG_MARGIN_HOURS`를
올리는 것이 먼저다.

---

## 4. 본번 가동 절차서 (B 인계물 — 사용자 승인 후)

A(이 문서·compose·리허설)는 코드 판정까지다. 아래는 parkjunwoo.com 본번
가동(B)의 절차 — **자격증명 주입과 parkjunwoo.com 저장소 CI 수정은 사용자
승인 필요**. 리허설 1회전(`docs/rehearsal/2026-06-loop1/`)이 같은 루프의
픽스처 증명이다.

### 4.1 자격증명 체크리스트

전부 백엔드(.env)에만 둔다 — 에이전트·영수증·큐 파일에 절대 없다(설계서 §3.3).

| # | 자격증명 | .env 키 | 발급/권한 |
|---|---|---|---|
| 1 | postgres 비밀번호 | `POSTGRES_PASSWORD` | 생성: `openssl rand -hex 16` |
| 2 | JWT 시크릿 (≥32자) | `JWT_SECRET` | 생성: `openssl rand -hex 32` |
| 3 | operator 계정 | `ABLOQD_OPERATOR_EMAIL/PASSWORD` | §4.2에서 직접 시드. 비밀번호에 `"`·`\` 금지 |
| 4 | **CF 로그 RO** — CloudFront 표준 로그 버킷 읽기 전용 IAM | `CF_LOG_SOURCE`(s3://…), `AWS_ACCESS_KEY_ID/SECRET_ACCESS_KEY/REGION` | s3:GetObject+ListBucket만 (로그 prefix 한정) |
| 5 | **GSC 서비스 계정 JSON** | `GSC_SA_JSON_FILE` | GCP SA 1개 — Indexing API(`indexing`)와 Search Console(`webmasters.readonly`) scope, Search Console 속성에 SA 이메일을 사용자로 추가 |
| 6 | IndexNow 키 | `INDEXNOW_KEY` | 임의 hex — 사이트 루트에 `<key>.txt`로도 배포돼야 한다 |
| 7 | Wayback SPN2 키 쌍 | `WAYBACK_ACCESS_KEY/SECRET_KEY` | https://archive.org/account/s3.php |
| 8 | **엔진 API 키** (옵트인) | `PERPLEXITY/OPENAI/ANTHROPIC_API_KEY` | 인용 샘플링용 — blog.yaml `geo.citation_budget`>0일 때만 의미 |
| 9 | **deploy key** — 블로그 저장소 쓰기 | `QUEUE_EXPORT_REPO_URL`(ssh), `QUEUE_DEPLOY_KEY_FILE`, `QUEUE_KNOWN_HOSTS_FILE` | GitHub deploy key(write) — 큐 발급·리포트 발행 push용 |
| 10 | CI 시크릿 (블로그 저장소 측) | `ABLOQD_URL`, `ABLOQD_OPERATOR_EMAIL/PASSWORD` | §4.4 배포 훅용 — abloqd 것과 동일 계정이어도 된다 |

### 4.2 기동 순서

```bash
# ① 생성물 투영 (backend/arts는 일회용 — 커밋되지 않는다)
backend/scripts/local-goproxy.sh   # pkg 변경이 있었으면 버전 범프 후
export GOPROXY="file:///tmp/abloq-goproxy,https://proxy.golang.org,direct"
yongol generate backend/specs backend/arts

# ② 시크릿 작성
cp deploy/backend/.env.example deploy/backend/.env   # §4.1 값 기입

# ③ 코어 기동 (postgres가 최초 기동에서 migrations를 적용한다)
docker compose -f deploy/backend/docker-compose.yaml up -d --build

# ④ operator 계정 시드 (1회 — 시드 API 없음, bcrypt 해시 직접 INSERT.
#    hashgen은 backend/scripts/rehearsal/run.sh와 동일 스니펫 — arts 모듈의
#    x/crypto 의존성을 빌려 쓴다)
cat > /tmp/hashgen.go <<'EOF'
//go:build ignore

package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	h, err := bcrypt.GenerateFromPassword([]byte(os.Args[1]), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(h))
}
EOF
HASH=$(cd backend/arts/backend && go run /tmp/hashgen.go "$ABLOQD_OPERATOR_PASSWORD")
docker compose -f deploy/backend/docker-compose.yaml exec -T postgres \
  psql -U abloqd -d abloqd -c \
  "INSERT INTO users (email, password_hash, role) VALUES ('<email>', '$HASH', 'operator');"

# ⑤ 초기 적재 (B 작업 3 — login 후 사이트마다 순서대로, 전부 멱등.
#    단일 사이트 env 배포의 사이트 이름은 default)
#    POST /sites/<s>/sync          : posts 인덱스
#    POST /sites/<s>/ingest/crawl  : CF 로그 백필 (보존 구간 — 커서가 따라온다)
#    POST /sites/<s>/ingest/gsc    : GSC 백필 (GSC_LOOKBACK_DAYS, 한도 16개월)
#    /sites/<s>/citation-queries   : 질의 셋 초기 작성 (에이전트 보조 1회)

# ⑥ cron 프로필 켜기 (전부)
COMPOSE_PROFILES=backstop,crawl,gsc,citation,report,freshness,evidence,cluster,queue \
  docker compose -f deploy/backend/docker-compose.yaml up -d
```

### 4.3 검증 항목 (기동 직후)

1. `GET /health` 200.
2. `docker compose config` 무에러 + `backend/scripts/compose-cron-smoke/run.sh`
   PASS (cron 명령 셸 전개 — 이스케이프 회귀 방지).
3. operator login 200 → `GET /sites`에 선언 사이트 전부 → 사이트마다
   `POST /sites/<s>/sync`의 synced 수 = 그 사이트의 발행 글 수.
4. `POST /sites/<s>/ingest/crawl` 후 `GET /sites/<s>/crawl-hits` 행 존재 (CF RO 자격 검증).
5. `POST /sites/<s>/ingest/gsc` 후 스냅샷 적재 (SA scope 검증).
6. `POST /sites/<s>/queue/export` 200 (deploy key push 검증 — 큐가 비면 no-op도 정상).
7. 첫 배포 후 `GET /sites/<s>/receipts?deploy_id=<id>` 전행 done (아카이버 3종 자격 검증).
8. 월말: `GET /sites/<s>/reports/monthly/<ym>` 존재 + 그 사이트 블로그 저장소
   `reports/<ym>.md` 커밋 ((site, ym) 독립 — 사이트마다 1개).

### 4.4 parkjunwoo.com CI 훅 추가 절차 (승인 후 — 원본 저장소 수정)

> **읽기 전용 제약 해제 필요** — 이 단계만 parkjunwoo.com 저장소를 수정한다.

1. 저장소 CI 시크릿 등록: `ABLOQD_URL`, `ABLOQD_SITE`(사이트 이름 — 단일
   사이트 env 배포는 `default`), `ABLOQD_OPERATOR_EMAIL`,
   `ABLOQD_OPERATOR_PASSWORD` (§4.1 #10).
2. 배포 파이프라인 끝(빌드·업로드 뒤)에
   `template/files/deploy/archiver.md`의 3단계를 그대로 추가:
   login → `POST /sites/$ABLOQD_SITE/hooks/deployed`(deploy_id=커밋 SHA,
   changed=실변경 글 URL 배열) → `POST /sites/$ABLOQD_SITE/archive/process`
   (실패 비치명 — backstop이 흡수).
3. 검증: 글 1편 수정 배포 → `GET /sites/<s>/receipts?deploy_id=<sha>` 3행 done +
   Wayback 타임스탬프 확인(원저자 시점 증거).
4. 이후 Phase008·012·013·014 이월 본번 판정 일괄 수행 + 갱신 루프 실 3편
   (Phase019 계획 B항).

---

## 5. 멀티사이트 운용 (Phase020, v0.2.0)

### 5.1 사이트 추가/제거 — sites.yaml이 유일한 손잡이

런타임 등록 API는 없다(선언→기동 반영이 abloq 철학). 절차:

```bash
# ① 블로그 저장소를 호스트에 두고 compose에 /blogs/<name> 마운트 추가
# ② deploy/backend/sites.yaml 에 사이트 항목 추가 (sites.yaml.example 참조)
# ③ 재기동 — 기동 동기화가 upsert + 누락 deactivate를 수행한다
docker compose -f deploy/backend/docker-compose.yaml up -d abloqd
# ④ 확인
curl -sf "$ABLOQD_URL/sites?active_filter=true" -H "Authorization: Bearer $TOKEN"
```

- **제거 = 선언 삭제**: sites.yaml에서 항목을 지우고 재기동하면 그 사이트는
  `active=false`가 되고 모든 `/sites/<name>/…` 호출이 404가 된다. 행과 FK
  이력(posts·receipts·reports·…)은 보존된다 — 다시 선언하면 그대로 복귀.
- **이름 변경 금지에 준함**: name은 FK 앵커의 키다. 바꾸면 새 사이트가
  생기고 옛 이름은 비활성으로 남는다(이력 분리).
- sites.yaml에 진단(이름 형식, 중복, repo_path 비절대 등)이 있으면 abloqd가
  **기동을 거부**한다 — `docker compose logs abloqd`에 진단이 한 줄씩 남는다.

### 5.2 기존 단일 사이트 배포의 v0.2.0 마이그레이션

1. **DB**: abloqd 중지 후 `deploy/backend/migrate-0.2.0-multisite.sql` 적용 —
   sites 테이블 생성 + 기존 데이터 전부 `name='default'` 행으로 백필.
   ```bash
   docker compose -f deploy/backend/docker-compose.yaml stop abloqd
   docker compose -f deploy/backend/docker-compose.yaml exec -T postgres \
     psql -U abloqd -d abloqd < deploy/backend/migrate-0.2.0-multisite.sql
   ```
2. **env는 그대로 동작**: `SITES_YAML_PATH` 없이 `BLOG_REPO_PATH`만 있으면
   기동 시 default 사이트 1개가 기존 사이트 단위 env 8종으로 합성된다.
3. **호출 경로만 갱신 (호환 깨짐)**: cron은 compose 갱신으로 자동 반영,
   블로그 저장소 CI 훅은 `/hooks/deployed`→`/sites/default/hooks/deployed`,
   `/archive/process`→`/sites/default/archive/process` (`ABLOQD_SITE=default`
   시크릿 추가 — `template/files/deploy/archiver.md`).
4. **sites.yaml로 이관(권장)**: 사이트가 2개 이상이 되는 시점에
   사이트 단위 env 8종(`BLOG_REPO_PATH`, `QUEUE_EXPORT_REPO_URL·AUTHOR·
   AUTHOR_EMAIL`, `GSC_SITE_URL`, `GSC_SA_JSON_PATH`, `CF_LOG_SOURCE`,
   `INDEXNOW_KEY`)을 sites.yaml의 사이트 행으로 옮긴다. 빈 행 값은 전역
   env로 fallback하므로 점진 이관이 가능하다(행 값이 항상 이긴다).

### 5.3 격리 보장 (Hurl 회귀가 고정하는 사실)

- 데이터: 도메인 테이블 10종이 `site_id` FK 스코프 — 같은 slug가 두
  사이트에 공존하고, A의 sync/scan/queue/report가 B에 비가시.
- 제출 자격: IndexNow 키·GSC SA 경로는 사이트 행 값 우선 — A의 제출이
  B의 키로 나가지 않는다.
- 작업 클론: 큐 export `/var/lib/abloqd/queue-export/<name>/`,
  리포트 발행 `<name>-reports` — 사이트끼리 체크아웃을 공유하지 않는다.
- 미등록·비활성 사이트: 전 op 404 (fail-closed — path 누락은 오염이 아니라
  404다).
- 전역 공유로 남는 것(v1): 계정 자격증명(AWS·Wayback·엔진 키·deploy key,
  `GIT_SSH_COMMAND`), operator 계정(전 운영자 전 사이트), 인용 샘플링
  비용 상한(사이트 blog.yaml `geo.citation_budget`이 각자 제한).
