# 시범 운행 기록 — Phase018 큐 소비 퀘스트 3종 (2026-06-11)

git 초기화 임시 인스턴스(`/tmp/abloq-trial-018`, en/posts, freshness_days=30,
min_internal_links=1)에서 **발급(backend exporter)→소비(퀘스트 3종)→consumed 동기화**
루프를 1회 완주한 기록. 인용 URL은 로컬 스텁(127.0.0.1:18123, 경로→title 에코)으로
기계 검증 가능하게 구성. 발행 없음, 인스턴스·postgres는 기록 추출 후 폐기.

## 구성

- 글 3편(전부 git 커밋): `stale.md`(lastmod 2026-04-01 — freshness 초과),
  `sourcing.md`(무출처 수치 주장 1건), `thin.md`(고립 — min-internal-links +
  **no-isolated-post**).
- 임시 postgres(initdb, 127.0.0.1:55432 trust) + abloqd(arts 빌드, abloq v0.0.12)
  + bare origin. `POST /sync`(3) → `POST /scans/{freshness,evidence,cluster}`(각
  detected=1) → `POST /queue/export` → **exported=3**.
- bare 클론 = 에이전트 인스턴스. 큐 파일 3종이 `keys:` 블록을 동반 발급
  (`queue-file-sample.yaml`). CLI `abloq scan` 산출과 **diff -r 0(바이트 동일)** 확인.

## 타임라인

1. **refresh** — `scan`(1 item, Key `en/posts/stale`) → `next` 프롬프트(발급 근거
   lastmod 2026-04-01·30일 임계 + 프로토콜·tasks·context). 작업트리에서 수치 교체
   (40%/2025 → 62%/2026, 동일 출처 유지)·본문 갱신·lastmod 2026-06-11 전진.
2. **submit #1 — 의도적 queue-scope FAIL** (`thin.md`에 범위 밖 한 줄 추가 상태):
   ```
   en/posts/stale -> FAIL (state TODO)
     - queue-scope: content/en/posts/thin.md
       expected="changes limited to the target article, its language companions and insight sidecars"
       actual="out-of-scope change: content/en/posts/thin.md"
   ```
   → `queue-scope-fail.txt`. thin.md 원복(git checkout).
3. **submit #2 — PASS** (`en/posts/stale -> PASS`). 순서 박제대로
   **① 글 수정 커밋 → ② 큐 파일 삭제 커밋**(단일 언어라 번역 재동기화 해당 없음).
4. **evidence** — `scan`(1 item) → 큐 주장(해시 d481fd8aea2c4f77)에 같은 문단
   인라인 출처 추가(스텁 URL — citation-exists가 실제 GET으로 200+title 대조).
   **lastmod 불변경**(임계 미달 변경 정직 규약). submit → **PASS** → ①→②.
5. **cluster** — `scan`(1 item, 위반 2종 + 후보 2건) → 대상 글에 out 링크
   + **payload 후보(stale)에 in 앵커 한 줄**(lastmod 불변경). 게이트가
   pkg/scan/cluster를 재실행해 위반 소멸 확인. submit → **PASS** → ①→② → push.
6. **consumed 동기화** — `POST /queue/export` 2회전:
   ```
   {"consumed":3,"exported":0}
   GET /queue?status=consumed → refresh/evidence/cluster 3행 status=consumed
   ```

## 파일

| 파일 | 내용 |
|---|---|
| `blog.yaml` | 시범 인스턴스 선언 |
| `queue-file-sample.yaml` | 발급된 refresh 큐 파일 (key + keys: 블록) |
| `queue-scope-fail.txt` | 의도적 범위 밖 변경의 FAIL Fact 원문 |
| `{refresh,evidence,cluster}-session.json` | reins 세션 — 시도 이력(refresh는 FAIL→PASS) |
| `{refresh,evidence,cluster}-results.jsonl` | export 산출물 (emit-once) |
| `agent-commits.txt` | 에이전트 커밋 순서(수정 커밋 → 소비 커밋, 퀘스트별) |

## 메모

- 인스턴스 `.gitignore`에 `.abloq/` 필요 — citation-exists 영수증 캐시가 untracked로
  남으면 다음 퀘스트의 queue-scope가 거짓 FAIL한다. 템플릿 `.gitignore`를
  `.abloq/cache/` → `.abloq/`로 같이 갱신했다(Phase018).
