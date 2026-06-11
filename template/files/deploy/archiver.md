# deploy/archiver — 배포 직후 아카이버 결선 (abloqd Phase008)

배포 파이프라인 끝에 **3단계**를 추가한다. 아카이버는 변경 URL마다
Wayback 저장 + IndexNow 제출 + GSC 색인 요청을 실행하고 영수증을 남긴다 —
Wayback 타임스탬프가 원저자 시점 증거이므로 **배포 직후 실행이 곧 증거
가치**다. 에이전트는 외부 API를 직접 치지 않는다: CI도 abloqd webhook만
호출한다.

## 시크릿

CI 시크릿에는 **operator 자격증명**(이메일/비밀번호)을 보관한다 — 토큰이
아니다. access token TTL은 15분이라 토큰을 저장하면 첫 배포 15분 뒤부터
전부 401이 된다. 매 실행 ①에서 새로 발급받는다 (login rate limit는
5회/분·IP — 실행당 1회라 무관).

| secret | 값 |
|---|---|
| `ABLOQD_URL` | abloqd 베이스 URL (예: `https://ops.example.com`) |
| `ABLOQD_SITE` | 이 블로그의 사이트 이름 — abloqd sites.yaml의 `name` (단일 사이트 env 배포는 `default`) |
| `ABLOQD_OPERATOR_EMAIL` / `ABLOQD_OPERATOR_PASSWORD` | operator 계정 |

abloqd v0.2.0(멀티사이트)부터 도메인 API는 전부 `/sites/<name>/…` 하위다 —
이 블로그가 등록된 사이트 이름을 `ABLOQD_SITE`로 주입한다.

## 파이프라인 3단계 (빌드·업로드 뒤에)

```bash
# ① login — 매 실행 토큰 발급
TOKEN=$(curl -sf -X POST "$ABLOQD_URL/auth/login" \
  -H 'Content-Type: application/json' \
  -d "{\"email\":\"$ABLOQD_OPERATOR_EMAIL\",\"password\":\"$ABLOQD_OPERATOR_PASSWORD\"}" \
  | jq -r .access_token)

# ② webhook — 변경 URL을 pending 영수증으로 적재 (멱등, 202 즉시)
#    DEPLOY_ID는 커밋 SHA 등 배포 1회당 유일값. CHANGED_URLS는 이번 배포에서
#    본문이 실변경된 글의 정규 URL JSON 배열 (예: git diff → slug → URL).
curl -sf -X POST "$ABLOQD_URL/sites/$ABLOQD_SITE/hooks/deployed" \
  -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' \
  -d "{\"deploy_id\":\"$DEPLOY_ID\",\"changed\":$CHANGED_URLS}"

# ③ process — 직접 호출로 즉시 실행 (실패 비치명: 멱등 + limit 보호,
#    cron 백스톱이 흡수). Wayback SPN2는 수십 초 걸릴 수 있어 타임아웃을 준다.
curl -s -m 600 -X POST "$ABLOQD_URL/sites/$ABLOQD_SITE/archive/process" \
  -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' \
  -d '{"limit":200}' || echo "archive process deferred to the cron backstop"
```

③이 실패해도 배포는 성공으로 둔다 — pending 영수증은 cron 백스톱
(abloqd compose의 `archiver-backstop`, 권고 주기 15분: login → 사이트 순회
→ retry → process)이 처리한다. 결과 확인은
`GET /sites/<site>/receipts?deploy_id=<id>`.
