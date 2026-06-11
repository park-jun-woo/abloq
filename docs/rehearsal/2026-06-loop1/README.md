# 리허설 기록 — Phase019-A 운용 루프 1회전 (2026-06-11)

`backend/scripts/rehearsal/run.sh`가 임시 인스턴스(git)+임시 postgres(127.0.0.1:55432
trust)+아카이브 스텁 위에서 **[측정→큐→퀘스트→게이트→발행→측정] 루프를 기계 판정으로
1회 완주**한 기록. 사람 개입 0 — 글 수정(에이전트 노동)은 스크립트에 박제된 결정적
편집이고, 판정은 전부 게이트 룰·HTTP·git assert다. 인스턴스·postgres·스텁은 종료 후
폐기, 이 디렉토리만 남는다. 재실행: `backend/scripts/rehearsal/run.sh [record-dir]`.

## 구성

- 인스턴스: ko/tech 2편, baseURL `https://fixture.example.com` —
  `backend/fixtures/cflogs`(CF 로그)와 아카이브 스텁의 GSC Search Analytics 행이
  같은 URL(`/tech/post-a/`, `/tech/post-b/`)로 조인되도록 설계.
  `post-a` lastmod 2026-04-01(freshness_days 30 초과 — 항상 stale),
  `post-b` lastmod는 실행일로 생성(재실행해도 항상 fresh — detected=1 고정).
- 토폴로지: bare origin 1개 = 큐 발급처 = 에이전트 클론 원점 = 리포트 발행처.
  별도 "deploy" 클론이 본번 체크아웃(BLOG_REPO_PATH) 역할 — ① 커밋 후
  `git pull`이 CI 배포를 대신한다.

## 타임라인 (전부 하니스 1회 실행, exit 0)

1. `POST /sync`(synced=2) → `POST /ingest/crawl`(CF 픽스처) → `POST /ingest/gsc`(스텁)
   → `POST /scans/freshness`(detected=1) → `POST /queue/export`(exported=1).
2. 에이전트 클론 → `abloq quest refresh scan`/`next`(발급 근거: lastmod 2026-04-01,
   freshness_days 30 — `next-prompt.txt`). 편집: 수치 교체(40%/2025→55%/2026, 동일
   출처 유지) + 본문 갱신 + lastmod 전진(2026-06-11).
3. **submit #1 — 의도적 queue-scope FAIL**(post-b.md에 범위 밖 한 줄):
   `queue-scope-fail.txt`. 원복 후 **submit #2 — PASS**(lastmod-advance·
   honest-lastmod·queue-scope 등 14룰 통과).
4. 순서 박제대로 **① 글 수정 커밋** → deploy pull → 재-`POST /sync`(post-a lastmod
   2026-06-11 인식 확인) → `POST /hooks/deployed`(planned=3) →
   `POST /archive/process`(failed=0) → 영수증 3행 done(스텁).
5. **② 큐 파일 삭제 커밋** → `POST /queue/export` 2회전: consumed=1, exported=0.
6. `POST /reports/monthly`(ym=2026-06, published=true, articles=2) — bare origin에
   `reports/2026-06.md` 발행 커밋 확인.

## 증거물 4종 (완료판정 A)

| # | 파일 | 내용 | 실측값 |
|---|---|---|---|
| ① | `refresh-session.json` | reins 게이트 세션 — try1 FAIL(queue-scope)→try2 PASS 추적 | state=PASS, log=[FAIL, PASS] |
| ② | `receipts-done.json` | `GET /receipts?deploy_id=rehearsal-loop1&status=done` | wayback·indexnow·gsc_index 3행 done |
| ③ | `report-publish.txt` | bare origin의 `reports/2026-06.md` 발행 커밋 해시 | `8e41fae6e1459d86c0a3cd6c251ac1d33246203e` |
| ④ | `export-consumed.json` | export consumed 동기화 결과 + consumed 행 조회 | `{"consumed":1,"exported":0}` + 1행 status=consumed |

## 보조 기록

| 파일 | 내용 |
|---|---|
| `queue-file.yaml` | 발급된 큐 파일 (key + keys: 게이트 계약) |
| `next-prompt.txt` | `quest refresh next`가 에이전트에게 보여준 프롬프트 원문 |
| `queue-scope-fail.txt` | 의도적 범위 밖 변경의 FAIL Fact 원문 |
| `agent-commits.txt` | 에이전트 커밋 순서 박제 (① 수정 커밋 → ② 소비 커밋) |
| `report-2026-06.md` | 발행된 월간 가시성 리포트 사본 (Queue intake에 consumed 1행) |
| `steps.log` | 단계별 API 응답 전문 |

## 메모

- reins 세션 파일(`--session`)은 **인스턴스 밖**에 둬야 한다 — 클론 안의 untracked
  파일은 작업트리 변경으로 잡혀 queue-scope가 (정당하게) FAIL한다. Phase018 시범의
  `.abloq/` 메모와 같은 계열 — 하니스 첫 실행에서 실제로 재현됐고 스크립트에 박제했다.
- `submit`은 FAIL에도 exit 0이다(verdict는 `key -> OUTCOME` 출력 라인) — 하니스는
  출력 grep으로 판정한다.
