# 시범 운행 기록 — Phase016 집필 퀘스트 (2026-06-11)

일회용 인스턴스(`abloq init /tmp/abloq-trial-016`, en/posts, 기본 구조 body→sources,
min_sources=1)에서 명세→집필→게이트 루프를 1회 완주한 기록. 발행 없음, 인스턴스는
기록 추출 후 폐기. 사람 개입은 인사이트 결정(insight.yaml)뿐이고, REVIEW는 별도
컨텍스트 에이전트가 대행했다.

## 타임라인

1. **scan** — `abloq quest writing scan content/en/posts/robots-exclusion-protocol.insight.yaml`
   → `seeded 1 item(s)` (Key `en/posts/robots-exclusion-protocol`).
2. **next** — 저작 프롬프트(insight 원문 + tasks.md T1~T4 + context.md) 출력. 집필
   컨텍스트(claude-code-phase016-writer)가 T1 수집(실 URL 200 확인) → T2 초안 →
   T3 퇴고(match 4/5, planted-claim 미출현 확인) 수행. → `worklog.md`
3. **submit #1 — 의도적 FAIL** (Sources 섹션 누락 상태로 제출):
   ```
   en/posts/robots-exclusion-protocol -> FAIL (state TODO)
     - min-sources: content/en/posts/robots-exclusion-protocol.md:1
       expected="the sources section lists at least geo.min_sources entries"
       actual="sources section missing — geo.min_sources requires >= 1 source(s)"
     - citation-exists: content/en/posts/robots-exclusion-protocol.md:19
       expected="new citation URLs answer HTTP 200 and match the cited title (skipped offline)"
       actual="citation URL https://www.robotstxt.org/orig.html is not reachable (HTTP 403)"
     - review-record: quests/writing/reviews/robots-exclusion-protocol.md
       expected="REVIEW record file (reviewer: line + claim dispositions)"
       actual="missing or empty"
   ```
   citation-exists 건은 **의도하지 않은 실검출**: robotstxt.org가 브라우저 UA에만 200을
   주고 Go 클라이언트에 403 — 출처 날조가 아니어도 기계 검증 불가 인용을 게이트가 차단.
4. **수정** — Fact 반영: ① `## Sources` 섹션 추가 ② robotstxt.org 인용을 RFC 9309
   자체의 1994 출처 서술로 교체(robotstxt.org는 비링크 항목으로 명기) ③ T4 의뢰.
   → `worklog.md`의 "FAIL #1 반영" 절.
5. **T4 REVIEW — 컨텍스트 격리** — 집필 컨텍스트가 기록을 쓰지 않고, 별도 프로세스
   `claude -p`(headless, sonnet)에 글·insight·match 결과·로그를 제시. 검토자가
   `reviewer: claude-headless-reviewer-20260611` 식별자와 미출현 claim 전건의
   disposition(`- planted-claim: excluded — …`)을 포함한 기록을 작성. → `review.md`
6. **submit #2 — PASS**:
   ```
   en/posts/robots-exclusion-protocol -> PASS (state PASS)
   ```
   tries=1(FAIL 1회) 후 2번째 제출에서 PASS — MaxTries=3 내 수렴. 래칫 잠금·export
   (`results.jsonl`) 자동 수행. 시도 이력은 `session.json`의 log에 보존.

## 파일

| 파일 | 내용 |
|---|---|
| `insight.yaml` | 사람 입력(인사이트 결정) — claims 5건, requires_source 2건, planted 1건 |
| `article.md` | 게이트 PASS 최종 글 |
| `worklog.md` | T1~T3 + FAIL #1 반영 기록 (집필 컨텍스트) |
| `review.md` | 별도 컨텍스트 검토자의 REVIEW 기록 (reviewer 식별자 + disposition) |
| `session.json` | reins 세션 — FAIL→PASS 시도 이력 |
| `results.jsonl` | export 산출물 (emit-once) |
