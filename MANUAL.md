# abloq — 에이전트 운용 매뉴얼

에이전트가 이 문서만 읽고 abloq 블로그 인스턴스를 운용할 수 있게 하는 매뉴얼이다.
명령·룰 ID·키는 코드와 1:1이며, 상세 스키마는 `docs/`로 연결한다.

- 개념·동기·비교: `README.md`
- blog.yaml 전 스키마: `docs/blog-yaml.md`
- insight.yaml 작성: `docs/insight-spec.md`
- abloqd 운영: `docs/operations.md`
- 게이트 엔진: [reins](https://github.com/park-jun-woo/reins) `MANUAL.md`

---

## 1. 핵심 모델 (짧게)

```
SSOT(blog.yaml) → 파생물(generate) → 드리프트 검증(check) → 게이트(gate/quest) → 측정(scan·ingest·report) → 래칫
```

- **SSOT 한 장.** `blog.yaml`이 사이트·언어·섹션·글의 정규 구조·GEO 임계·배포를
  전부 선언한다. `hugo.toml`·`robots.txt`·`llms.txt`·sitemap(hreflang)·JSON-LD·
  게이트 룰 파라미터가 전부 여기서 파생된다. blog.yaml이 바뀌지 않는 한 어떤 글도
  게이트를 우회할 수 없다.
- **결정적 게이트 vs 비결정 퀘스트.** 산문을 만지는 비결정 노동만 에이전트(퀘스트)가
  한다. 생성·스캔·측정·외부 API·판정은 전부 결정적 코드다.
- **게이트는 위반 검출기.** 룰이 발동(fire)하면 `Fact{위치·기대·실제}`를 돌려준다.
  FAIL은 의견이 아니라 위치·수치가 박힌 사실이다 — 그것만 고쳐서 재제출한다.
- **래칫.** 측정 결과가 우선순위 큐의 가중치가 되어 다음 퀘스트 입력을 지정한다.
  PASS는 단방향이고 남은 일은 단조 감소한다. `MaxTries=3` — FAIL 3회 누적 시
  DONE으로 영구 잠금된다.

권한 비대칭: **기계(L1)만 PASS를 잠근다. 에이전트(L2)는 REVIEW만, 사람(L3)은 나머지.**

---

## 2. 설치 · 인스턴스 생성 · 일상 루프

### 설치

```bash
go install github.com/park-jun-woo/abloq/cmd/abloq@latest
```

### 인스턴스 생성 (init)

`init`은 비대화형(에이전트 경로)이 기본이다. 플래그로 전부 지정한다.

```bash
abloq init my-blog \
  --title "My Blog" \
  --baseurl https://example.com \
  --author "Author" \
  --languages en,ja,ko \   # 콤마 구분 BCP-47, 첫 항목 = 기본 언어
  --sections posts,tech    # 콤마 구분 섹션
cd my-blog
```

생성물: `blog.yaml` + 임베드 템플릿 + `CLAUDE.md`(에이전트 운영 매뉴얼) — **게이트
클린 상태**로 나온다. `--interactive`를 주면 플래그 대신 프롬프트로 묻는다.

### 일상 루프

```bash
abloq validate .     # blog.yaml 자체 검증 (스키마·룰)
abloq generate .     # blog.yaml → 파생물 재생성 (hugo.toml·robots.txt·llms.txt·jsonld.json)
abloq check .        # 파생물이 SSOT와 일치하는지 (드리프트 시 exit 1)
abloq gate .         # 글 본문에 구조·근거·정책 14룰 적용 (위반 시 exit 1)
```

빌드 후처리:

```bash
hugo                 # 정적 사이트 빌드
abloq postbuild md . # 글마다 노이즈 제로 .md를 public/에 병행 산출 (AI 컨텍스트 포맷)
```

전 명령 공통: 진단이 있으면 **exit code 1**, 깨끗하면 0. 진단 형식은
`파일:라인 [룰ID] 메시지` 한 줄씩. `--json`이 있는 명령은 진단을 JSON 배열로 낸다.

---

## 3. CLI 레퍼런스

`[dir]`은 blog.yaml이 있는 블로그 루트로 기본값 `.`이다. exit 1 = 진단/위반 존재.

### 코어 (SSOT · 파생 · 게이트)

| 명령 | 인자·플래그 | 출력 / 의미 |
|---|---|---|
| `abloq validate [dir]` | `--json` | blog.yaml 스키마·룰 검증. 진단 시 exit 1 |
| `abloq generate [dir]` | — | blog.yaml → hugo.toml·robots.txt·llms.txt·jsonld.json 생성 |
| `abloq check [dir]` | — | 파생물 vs 새 재생성 대조. 드리프트 시 exit 1 |
| `abloq gate [dir]` | `--rule <id>` `--json` `--offline` | 글에 14룰 적용. 위반 시 exit 1. `--offline`은 네트워크 룰(citation-exists) 스킵 |
| `abloq init <dir>` | `--title --baseurl --author --languages --sections --interactive` | 새 블로그 스캐폴드 (게이트 클린) |
| `abloq claudemd [dir]` | — | blog.yaml에서 CLAUDE.md 재생성 |
| `abloq postbuild md [dir]` | — | hugo 빌드 후 글마다 클린 .md를 public/에 산출 |

### 이미지 (local은 순수 Go — cgo·외부 바이너리 없음, AI provider는 opt-in HTTP)

| 명령 | 인자·플래그 | 의미 |
|---|---|---|
| `abloq image og <slug> <title>` | `--brand --font --out` `--provider --model --overlay` `--variant --all-variants --count` | 1200×630 OG WebP 생성. 기본(플래그 없음)은 결정론 local 카드를 `static/images/{slug}.webp`에 직행 — 현행 그대로. provider 해석: `--provider` > blog.yaml `image.og` > `local` |
| `abloq image convert <src>` | `--slug --max-width --out` | 이미지를 WebP로 변환 (흰 배경 평탄화·리사이즈). `--max-width` 기본 1400(0=유지) |

#### AI OG (provider `gemini`) — 생성→검토→채택 절차

AI 이미지는 비결정·네트워크·유료다 — **`generate`/`check` 루프에 절대 들어가지
않는다**(바이트 멱등 파생 가정 파괴 금지). 절차상 위치는 인용 샘플링과 같다:
**명시 호출 1회 생성 → 검토 → `mv` 채택 → 커밋**, 이후 평범한 정적 자산.

```bash
abloq image og <slug> <title> --provider gemini            # 기본 설정 1안 1장 → static/images/{slug}.webp 직행
abloq image og <slug> <title> --variant minimal,photo      # 지정 안들 → files/og/{slug}/{variant}-{n}.webp 드래프트
abloq image og <slug> <title> --all-variants --count 2     # 선언 전 안 × 2장 샘플링 (드래프트)
mv files/og/<slug>/minimal-1.webp static/images/<slug>.webp  # 검토 후 채택 (후보가 이미 최종 규격) → 커밋
```

- 다중 안(`--variant`/`--all-variants`/`--count>1`)은 마지막 생성본의 묵시적 채택을
  막기 위해 드래프트 디렉토리(`files/og/` — 빌드 경로 밖)로만 나간다. variants
  미선언 + `--count>1`의 기본 안 파일명은 예약명 `default-{n}.webp`.
- `--overlay`는 AI 배경 위에 제목/브랜드를 **결정론적으로** 합성한다(local 카드와
  같은 합성 코드). 프롬프트는 blog.yaml `image.og.prompt`(`{title}`/`{summary}`/`{brand}`
  치환, 미선언 시 no text·safe margin 포함 내장 템플릿) — 글자는 overlay 담당.
- 실행 전 "생성 예정 N건"과 안별 모델을 echo, 성공 경로마다 사용 모델 echo.
  안 일부 실패 시 성공분은 보존하고 exit 1.
- API 키는 **env 전용**: `GEMINI_API_KEY`(또는 `GOOGLE_API_KEY`). 부재 시 명확한
  진단 + exit 1. `GEMINI_API_BASE`로 테스트 스텁 오버라이드. blog.yaml·인자에 키 금지.

### 스캐너 (큐 생성 — `quests/queue/`에 직접 기록)

| 명령 | 인자 | 의미 |
|---|---|---|
| `abloq scan freshness [dir]` | — | `geo.freshness_days` 초과 글 검출 → kind=refresh 큐 |
| `abloq scan evidence [dir]` | — | 무출처 수치 주장 검출 → kind=evidence 큐 + 인용 link rot 1회 점검 보고 |
| `abloq scan cluster [dir]` | — | 태그·내부링크 그래프 위반 검출 → kind=cluster 큐 (연결 후보 동봉) |

### 측정 (가시성 — 전부 무상태·DB 없음)

| 명령 | 인자·플래그 | 의미 |
|---|---|---|
| `abloq ingest crawl` | `--source <dir\|s3://...>` (필수) `--repo` | CloudFront 로그에서 AI봇 히트 집계 (1회) |
| `abloq ingest gsc` | `--site --days --repo` | GSC Search Analytics 최근 N일 조회 (1회). `--days` 기본 7 |
| `abloq sample citations` | `--queries <yaml\|json>` (필수) `--repo` | 질의 셋으로 AI 엔진 인용 1회 샘플링. **게이트화하지 않음 (비결정)** |
| `abloq report monthly` | `--ym <YYYY-MM>` `--source <logs>` `[dir]` | 로그+저장소만으로 부분 월간 리포트 (DB 없으면 인용/색인 계층 누락 명시) |

### 외부 부수효과 · 인사이트

| 명령 | 인자 | 의미 |
|---|---|---|
| `abloq archive <url>` | — | Wayback·IndexNow·GSC Indexing API 제출 (자격증명은 env). 백엔드 아카이버와 동일 코드경로 |
| `abloq insight match <insight.yaml> <article>` | — | insight claims를 본문 anchors로 스크리닝 (REVIEW 보조, 기본 언어 글 전용) |

### 퀘스트 (reins 게이트 — 5종)

각 퀘스트는 `scan / next / submit / status / export / rules` 하위 명령을 가진다.

```bash
abloq quest <writing|translation|refresh|evidence|cluster> <scan|next|submit|status|export|rules> ...
```

| 하위 명령 | 의미 |
|---|---|
| `scan <args>` | 입력 → 퀘스트 아이템 시드. writing=insight.yaml, translation=기본 언어 글, refresh/evidence/cluster=`[instance-dir]`(큐 소비) |
| `next` | TODO 1개 + 집필/검증 프롬프트 출력 |
| `submit --key <k> --in <file>` | 제출 → 게이트 판정 → PASS 잠금 / FAIL 시 Fact 피드백 |
| `status` | 진행 집계 (PASS/REVIEW/DONE/TODO/...) |
| `export` | 종료 아이템을 JSONL로 1회 방출 |
| `rules` | 게이트 룰 카탈로그 (감사용) |

> 환경 변수: `archive`·`ingest`·`sample`·`image og`(AI provider)는 자격증명을
> **env로만** 받는다 — `AWS_*`(S3 로그), `GSC_SA_JSON`/`GSC_SA_JSON_PATH`(GSC),
> IndexNow/Wayback 키, `GEMINI_API_KEY`/`GOOGLE_API_KEY`(OG gemini).
> 토큰·키를 인자나 blog.yaml에 넣지 않는다.

---

## 4. 퀘스트 실행 절차 (5종)

원칙: 퀘스트는 산문을 만지는 비결정 노동만 한다. **외부 API를 직접 치지 않는다**
(아카이브·색인은 백엔드 영수증). `submit` 결과가 FAIL이면 Fact를 정확히 반영해
재제출한다 — `MaxTries=3`.

### 4.1 writing (집필) — 사람의 인사이트 → 글 1편

입력은 사람이 작성한 `insight.yaml`(글 1편당 1장, 기본 언어에만 둔다). 작성법은
`docs/insight-spec.md`.

```bash
abloq quest writing scan <insight.yaml> [insight.yaml...]   # claims → 아이템 시드
abloq quest writing next                                    # 집필 프롬프트
abloq insight match <insight.yaml> <article>                # (보조) 미출현 claim 스크리닝
abloq quest writing submit --key <key> --in submission.json
```

`submission.json`:

```json
{
  "article": "content/<lang>/<section>/<slug>.md",
  "worklog": "quests/writing/logs/<slug>.md",
  "review": "quests/writing/reviews/<slug>.md"
}
```

경로는 전부 블로그 루트 기준 상대 경로. article은 시드된 대상 경로와 일치해야 한다.

**REVIEW 격리 규약 (필수):**
- 집필 에이전트는 **자기 글을 REVIEW할 수 없다.** REVIEW 기록은 반드시 별도
  컨텍스트(다른 세션/프로세스)의 검토자가 작성한다.
- REVIEW 기록은 `reviewer: <컨텍스트 식별자>` 라인(집필 컨텍스트와 달라야 한다) +
  `abloq insight match` 미출현 claim **전건**의 disposition 라인을 담아야 한다:
  - `- <claim-id>: addressed — ...` / `revised — ...` / `excluded — ...`
- `review-record` 룰이 기록 존재·reviewer 라인·미출현 claim 전건 disposition
  커버리지를 결정적으로 검사한다. 지지 판정 자체는 비결정이라 룰이 아니다.

### 4.2 translation (번역) — 기본 언어 글 → 전 언어

```bash
abloq quest translation scan <기본언어 글.md> [글.md...]   # 글 × (기본 제외 전 언어) 매트릭스 시드
abloq quest translation next
abloq quest translation submit --key <key> --in submission.json   # {"article": "content/<lang>/<section>/<slug>.md"}
```

아이템은 글 × 언어이고 언어마다 독립이다(한 언어 FAIL이 다른 언어 PASS에 무영향).
**`translation-parity` 룰이 보는 것** — 원문 대비 양방향 비교: 헤딩 레벨 시퀀스, 문단
블록 수(빈 줄 경계), 이미지 경로, 코드블록(바이트 동일), 외부 링크 URL, 내부 글 링크
(대상 언어 프리픽스로 치환), `date`·`lastmod`(원문 이식 = fm-mirror). `slug`는 전
언어 동일(`slug-consistency` 스코프드). 의미 동등성은 비결정이라 게이트하지 않는다.

> hugo 빌드는 인스턴스 전체를 대상으로 매 제출 실행된다(`hugo-build` 룰). hugo가
> PATH에 없으면 제출 자체가 중단된다(스킵 아님).

### 4.3 큐 소비 3종 (refresh · evidence · cluster)

스캐너(또는 백엔드)가 `quests/queue/*.yaml`에 떨어뜨린 큐를 소비한다. 큐 파일 1개 =
아이템 1개, priority 내림차순으로 시드된다.

```bash
abloq quest <refresh|evidence|cluster> scan [instance-dir]   # 큐 디렉토리 스캔
abloq quest <refresh|evidence|cluster> next
abloq quest <refresh|evidence|cluster> submit --key <key> --in submission.json   # {"article": "..."}
```

**순서 박제 — 절대 바꾸지 마라 (어기면 차단된다):**

1. **작업트리에서 글 수정** — 커밋하지 않은 상태로 둔다.
2. **submit → PASS** — 게이트는 작업트리(더티) vs git HEAD 기준선으로 판정한다.
   글을 먼저 커밋하면 작업트리==HEAD가 되어 기준선 룰(honest-lastmod·claim 룰·
   queue-scope)이 **공허 통과**한다 → 게이트 무력화 = 금지.
3. **① 글 수정 커밋** — PASS 이후에만.
4. **(해당 시) 번역 재동기화** — lastmod가 갱신됐다면 translation 퀘스트로 전 언어
   재동기화 후 커밋한다. 큐 파일의 `keys:`가 전 언어 키를 동반 발급하므로 번역
   커밋도 honest-lastmod 큐 등재 검사를 통과한다.
5. **② 큐 파일 삭제 커밋(소비 신호)** — 반드시 **마지막**. ②를 번역 재동기화보다
   앞당기면 repo/CI honest-lastmod가 번역 커밋을 재차단한다.

차단 메커니즘:
- 큐 파일을 게이트 전에 지우거나 고치면 → `queue-scope` FAIL (큐 파일은 게이트
  시점에 무변경이어야 한다 — 허용 집합 밖).
- payload는 Seed 시점에 고정된다 — 작업트리 큐 파일을 고쳐도 게이트가 보는 payload는
  불변.
- 빈 diff 무작업 통과는 각 퀘스트의 작업 완수 강제 룰(`lastmod-advance`·
  `claims-resolved`·`rot-resolved`·`cluster-resolved`)이 막는다.

**queue-scope 허용 집합:** 대상 글 + 그 전 언어 번역본 경로 + insight 사이드카
(+cluster는 payload candidates의 글 경로). blog.yaml·레이아웃·다른 글은 전부 범위 밖
(사람의 몫).

**lastmod 정직 규약:** refresh는 lastmod를 **전진**시켜야 한다(빈 갱신 금지).
evidence·cluster는 `geo.min_meaningful_diff` **미달** 변경(출처 한두 개, 앵커 한 줄)에
lastmod를 **올리지 않는다** — 미미한 diff + lastmod 갱신은 honest-lastmod FAIL.

**claim 라인 규약:** 주장 라인은 **리랩(줄바꿈 재배치) 금지** — 라인 단위 비교라
보존된 주장도 변경으로 검출돼 거짓 FAIL이 난다. 큐 payload에 없는 주장은 한 글자도
바꾸지 마라.

---

## 5. 게이트 진단 읽는 법

진단: `파일:라인 [룰ID] 메시지`. FAIL Fact = `{Where(위치), Expected(기대), Actual(실제)}`.
고칠 것은 Fact 하나다. 단일 룰만 돌리려면 `abloq gate --rule <id> .`.

### 5.1 repo 레벨 14룰 (`abloq gate`, 실행 순서)

| 룰 ID | 의미 | 흔한 FAIL | 수정법 |
|---|---|---|---|
| `image-first` | 첫 본문 라인이 메인 이미지 `![..](..)` | 본문이 텍스트로 시작 | 글 맨 위에 메인 이미지 라인 배치 |
| `image-attribution` | 메인 이미지 다음에 이탤릭 출처 라인 | 출처 라인 누락 | 이미지 바로 아래 `*출처...*` 추가 |
| `section-order` | 인식 섹션이 정규 상대 순서 | 섹션 순서 뒤바뀜 | `structure.order` 순서로 재배치 |
| `section-preserved` | 기준선 대비 인식 섹션 누락 없음 | 섹션 삭제됨 | 떨어뜨린 섹션 복원 |
| `body-lossless` | 모든 기준선 본문 라인 생존(멀티셋 부분집합) | 라인 삭제/변형 | 원 라인 보존 (구조 변환 퀘스트 전용) |
| `front-matter-intact` | front matter 무변경 (lastmod만 허용) | title/tags/slug 변조 | lastmod 외 키 원복 |
| `heading-canonical` | 인식 섹션은 `##` 레벨 헤딩 | 헤딩 레벨 불일치 | `##`로 통일 |
| `front-matter-schema` | 필수 필드 존재·타입 유효 | title/date/lastmod/tags 누락 | 필드 채우기 |
| `slug-consistency` | 전 언어 동일 slug, 누락 언어 없음 | slug 불일치·언어 누락 | 전 언어 같은 유효 slug |
| `honest-lastmod` | lastmod 갱신 = 의미 있는 본문 diff(`min_meaningful_diff`) + 큐 등재 | 빈 갱신·큐 미등재 | 실변경 동반 또는 lastmod 원복 |
| `hreflang-complete` | 빌드 페이지가 전 언어를 hreflang으로 상호참조 | hreflang 누락 | 전 언어 alternate 생성 (빌드 후) |
| `min-sources` | 출처 섹션 ≥ `geo.min_sources` | 출처 부족 | 출처 항목 추가 |
| `numeric-claim-sourced` | HEAD 이후 추가된 수치 주장이 같은 문단에 출처 링크 | 무출처 수치 | 같은 문단에 인라인 출처 링크 |
| `citation-exists` | 신규 인용 URL이 HTTP 200 + 제목 일치 (오프라인 스킵) | URL 404/403·제목 불일치 | 실재 URL로 교체 |

### 5.2 퀘스트 전용 룰

| 룰 ID | 채택 퀘스트 | 의미 |
|---|---|---|
| `review-record` | writing | REVIEW 기록 존재 + reviewer 식별자 + 미출현 claim 전건 disposition |
| `translation-parity` | translation | 번역이 원문 구조를 미러 (헤딩·문단블록·이미지·코드·링크·date/lastmod) |
| `hugo-build` | translation | hugo가 인스턴스 전체를 0 에러로 빌드 |
| `lastmod-advance` | refresh | lastmod가 기준선을 엄격히 전진 (실 갱신 강제) |
| `claim-preserved` | refresh | 수치 주장 건수 ≥ 기준선 (삭제로 해소 금지) |
| `claims-resolved` | evidence | 큐 claim 해시가 무출처로 남지 않음 |
| `rot-resolved` | evidence | 큐 rot URL이 인용에서 사라짐 |
| `claim-scope` | evidence | 큐 밖 수치 주장 한 글자도 불변 |
| `cluster-resolved` | cluster | 큐 클러스터 위반 종류가 재스캔에서 소멸 |
| `queue-scope` | refresh·evidence·cluster | 작업트리 변경이 큐 아이템 허용 집합 안 |

> 룰 채택/배제는 퀘스트마다 다르다 (단일 글 제출엔 `slug-consistency`·`hreflang-complete`
> 부적합, Base-nil 신규엔 기준선 비교 룰 무력 등). 정확한 채택표는
> `abloq quest <name> rules` 또는 `pkg/quests/<name>/rules.go` 주석 참조.

---

## 6. 치즈 방어 — 하지 말 것

| 하지 말 것 | 잡는 룰 / 메커니즘 |
|---|---|
| **자가 REVIEW** (집필 컨텍스트가 자기 글 REVIEW 작성) | `review-record` — reviewer 라인이 집필 컨텍스트와 같으면 격리 위반 (완전 차단은 비결정, 기록 형식이 감사 가능성을 만든다) |
| **범위 밖 변경** (blog.yaml·다른 글·레이아웃 수정) | `queue-scope` (큐 퀘스트), writing/translation은 산출물 1편 + 로그/리뷰로 한정 |
| **lastmod 위조** (본문 무변경 + lastmod 전진) | `honest-lastmod` (repo), `lastmod-advance`(refresh 강제), translation은 `fm-mirror`로 원문 이식 강제 |
| **출처 날조** (없는 URL·제목 무관 표기) | `citation-exists` — 실 HTTP 200 + 제목/og:title 토큰 겹침 검증 |
| **주장 삭제로 해소** | `claim-preserved`(refresh 건수 하한), `claim-scope`(evidence 큐 밖 보존), evidence는 출처 추가가 본질 |
| **주장 리워딩으로 검출 회피** | `numeric-claim-sourced` — 리워딩은 기준선 대비 신규 주장으로 검출 → 출처 필수 |
| **게이트 무력화** (글 선커밋으로 기준선 공허화) | 순서 박제 §4.3 — 더티 작업트리 vs HEAD 판정. 선커밋 시 기준선 룰 공허 통과 = 금지 |
| **큐 선삭제** (게이트 전 큐 파일 삭제) | `queue-scope` — 큐 파일은 게이트 시점 무변경, 삭제는 ② 커밋에서만 |
| **태그 전삭제로 위반 회피** | 스캔 재실행이 `no-orphan-tag`·링크 위반을 다시 잡는다 |

공통: 외부 부수효과(아카이브·색인)는 백엔드 영수증으로 처리하고 **에이전트는 외부
API를 직접 치지 않는다.** payload는 Seed 시점 고정 — 작업트리 큐 파일을 고쳐도 무효.

---

## 7. 백엔드 연동 (옵션) — abloqd

스케줄·상태(시계열·큐·영수증)가 필요하면 abloqd를 셀프호스트한다. 백엔드를 켜지
않아도 모든 검출은 위 CLI로 돈다 — 백엔드는 스케줄과 상태를 더한 것뿐이다.

```bash
yongol generate backend/specs backend/arts                    # SSOT → Go+Gin 코드
docker compose -f deploy/backend/docker-compose.yaml up -d --build
```

필수 env 3종: `POSTGRES_PASSWORD`·`JWT_SECRET`·`BLOG_REPO_PATH`
(템플릿 `deploy/backend/.env.example`). 자격증명은 env로만 주입한다.

주요 endpoint (전체·바디는 `backend/specs/api/openapi.yaml`):

| endpoint | 역할 |
|---|---|
| `GET /health` | 인증 없는 헬스체크 (1분 주기 외부 모니터) |
| `POST /auth/login` | operator 자격증명 → access token (TTL 15분) |
| `POST /scans/{freshness,evidence,cluster}` | 스캐너 → 큐 적재 |
| `POST /queue/export` | 큐 회전 (consumed 동기화 + 신규 발급 + push, 멱등) |
| `POST /ingest/{crawl,gsc}` | 크롤·색인 계층 수집 |
| `POST /sample/citations` | 인용 샘플링 |
| `POST /reports/monthly` | 월간 리포트 |
| `POST /archive/process` · `POST /receipts/retry` | 아카이버 영수증 처리·재무장 |
| `POST /hooks/deployed` | 배포 직후 훅 (CI에서 호출) |

cron 프로필 9종(주기·호출 순서)·장애 대응(영수증 retry→process 순서, 큐 적체
동기화)은 `docs/operations.md` 참조. 전 호출은 멱등이며 실패 주기는 다음 주기가
흡수한다.

---

## 8. 키 레퍼런스 (요약)

### blog.yaml (상세: `docs/blog-yaml.md`)

strict 파싱 — 스키마에 없는 키는 `unknown-key` 에러.

```yaml
site:
  baseURL: https://example.com        # 필수, 절대 http(s), query/fragment 금지
  title: My Blog
  author: Author
  default_lang_in_subdir: true        # false면 기본 언어를 사이트 루트에 서빙

languages: [en, ja, ko]               # 필수, BCP-47, 첫 항목 = 기본 언어
sections: [posts, tech]               # 필수, 1개 이상

structure:
  order: [image, attribution, body, related, sources, changelog]   # 정규 섹션 순서 = 구조 게이트 입력
  headings:                           # 헤딩 키 → 언어 → 현지화 텍스트 (기본 언어 필수)
    sources: { en: "Sources", ja: "出典", ko: "출처" }

geo:
  crawlers: { training: allow, bytespider: block }   # 봇/분류 → allow|block
  llms_txt: auto                      # auto|manual|off 단축형 또는 객체
  # llms_txt:                         # 객체 폼 (상세: docs/blog-yaml.md)
  #   mode: auto                      # manual|off면 generate·check가 llms.txt 불간섭
  #   languages: base                 # base(기본 언어 1개)|all|[en, ko]
  #   header: |                       # 사이트 포지셔닝 블록 (자유 마크다운)
  #   pinned: [{title, url, desc, group}]   # 선두 고정 엔트리
  #   section_labels: { posts: Pattern }    # 섹션 → 사람 라벨
  #   max_summary: 200                # 설명문 rune 상한 (0=무제한)
  jsonld: [Article, Person]
  freshness_days: 90                  # refresh 임계 (≥1)
  min_sources: 1                      # 근거 게이트 임계 (≥0)
  min_internal_links: 2               # 클러스터 게이트 임계 (≥0)
  min_meaningful_diff: 10             # honest-lastmod 토큰 diff 임계 (≥1)

# image:                              # 선택 — AI OG 선언 (상세: docs/blog-yaml.md)
#   og:
#     provider: gemini                # local(기본)|gemini — 키는 env(GEMINI_API_KEY)
#     overlay: true                   # AI 배경 위 제목/브랜드 결정론 합성
#     prompt: |                       # {title}/{summary}/{brand} 치환
#       Minimal abstract tech background for "{title}". No text.
#     variants:                       # 안 프리셋 — 미지정 필드는 상위 상속
#       - { name: minimal, prompt: Flat geometric shapes. }

deploy:
  provider: s3-cloudfront
  terraform: false
  indexnow: true
```

검증 룰 ID: `yaml-syntax`·`unknown-key`·`lang-bcp47`·`heading-default-lang`·
`sections-empty`·`threshold-range`·`baseurl-format`·`crawlers-policy`·
`llmstxt-mode`·`llmstxt-languages`·`llmstxt-pinned`·`llmstxt-labels`·`llmstxt-max-summary`·
`og-provider`·`og-variant-name` (+ 비차단 경고 `og-local-variants`).

### insight.yaml (상세: `docs/insight-spec.md`)

글 1편당 1장, **기본 언어에만** 둔다. 짝짓기는 파일 이름 기준.
번들 `index.md` → 같은 디렉토리 `insight.yaml`, 플랫 `{이름}.md` → `{이름}.insight.yaml`.

```yaml
topic: ...            # 주제 한 줄
stance: ...           # 관점
audience: ...         # 독자
section: tech         # 필수, 글 실위치와 일치
tags: [...]           # taxonomy SSOT 어휘
claims:               # 필수, ≥1
  - id: claim-a       # 필수, 유니크
    text: ...         # 필수, 사람 언어
    kind: claim       # 필수: claim|rebuttal|prediction|definition
    requires_source: true   # 출처 필수 여부
    anchors: [...]    # 본문 대응 확인용 어휘 (글 언어, 동의어 허용)
non_goals: [...]      # 다루지 않을 것 (범위 이탈 방지)
tone: ...             # 힌트, 게이트 비대상
```

검증 룰: `insight-claims-min`·`insight-claim-id-unique`·`insight-claim-kind`(에러),
`insight-claim-anchors-empty`(경고). `abloq insight match`로 미출현 claim을 스크리닝
(REVIEW 보조, 기본 언어 글 전용 — 출현이 대응을 보장하지 않는다).

---

## 9. 예정 변경 (v0.2.0 — 미구현, 참고만)

아래는 계획 확정·구현 전 단계다. **현 시점 코드에는 없다** — 이 절 외의 본 매뉴얼은 전부 현행 코드와 1:1이다. 구현이 끝나면 이 절은 해당 본문으로 흡수된다.

- abloqd 멀티사이트: 사이트 목록 `sites.yaml` SSOT, 도메인 endpoint가 `/sites/{site}/…`로 이동(**경로 호환 깨짐**), 사이트 단위 env 8종이 sites.yaml로 이관.

---

## 10. 한눈에

```
사람: blog.yaml + insight.yaml 작성 (인사이트 결정)
에이전트: quest writing → translation, 큐 소비(refresh/evidence/cluster)
기계: validate → generate → check → gate / quest 게이트가 PASS를 잠근다
측정: scan·ingest·report → 큐 가중치 → 다음 퀘스트 입력 (래칫)
```

라이선스: MIT (`LICENSE`·`NOTICE`).
