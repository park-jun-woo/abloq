# worklog — en/posts/robots-exclusion-protocol (집필 컨텍스트: claude-code-phase016-writer)

## T1 자료 수집 (2026-06-11)

- rep-standardized-2022 (requires_source):
  - 후보 URL: https://www.rfc-editor.org/rfc/rfc9309.html — HTTP 200, title "RFC 9309: Robots Exclusion Protocol"
  - 인용 예정 문구: "This document standardizes and extends the 'Robots Exclusion Protocol' ... originally defined by Martijn Koster in 1994"
  - 후보 URL: https://www.robotstxt.org/orig.html — HTTP 200, title "The Web Robots Pages" (1994년 원 규약 문서)
  - 인용 예정 문구: "A Standard for Robot Exclusion" (1994 합의 문서)
- rep-not-access-control (requires_source):
  - 후보 URL: https://www.rfc-editor.org/rfc/rfc9309.html — 동일 문서
  - 인용 예정 문구: "these rules are not a form of access authorization"
- rep-definition / ai-crawlers-geo: requires_source 아님 — 출처 불요, 본문 전개만.
- planted-claim: 본문에서 다루지 않기로 결정(주제 흐름과 무관, non_goals에 가까움) —
  미출현으로 남겨 REVIEW disposition 판단에 넘긴다.

## T2 초안 (2026-06-11)

- 구조: body → sources (blog.yaml structure.order).
- claims 순서와 무관하게 정의(rep-definition) → 역사(rep-standardized-2022) →
  한계(rep-not-access-control) → AI 시대 운용(ai-crawlers-geo) 순으로 전개.
- 출처 인라인 배치: RFC 9309, robotstxt.org 원 문서.

## T3 퇴고 (2026-06-11)

- anchors 자기 점검: RFC 9309 / 1994 / not a form of access authorization / voluntary /
  robots.txt / crawler / AI crawler / GEO — 본문 출현 확인.
- planted-claim anchors(caching TTL cost saving)는 의도적으로 미출현 — REVIEW로 이관.
- non_goals 점검: 업체별 차단 가이드·법적 논의 안 다룸.
- (기록) 1차 제출은 Sources 섹션 누락 상태로 제출 — 게이트 FAIL Fact 피드백 확인용 의도적 결함.

## FAIL #1 반영 (2026-06-11)

게이트 FAIL Fact 3건 수신:
1. min-sources — Sources 섹션 누락 (의도적 결함): "## Sources" 섹션 추가, RFC 9309 항목.
2. citation-exists — robotstxt.org/orig.html이 Go HTTP 클라이언트에 403 (브라우저 UA만 허용):
   본문 인용을 RFC 9309 자체의 1994 출처 서술("originally defined by Martijn Koster in 1994")로
   교체, robotstxt.org는 Sources에 비링크 항목으로 명기. 게이트가 잡은 실질 결함 — 의도하지 않은 발견.
3. review-record — 기록 부재: T4 REVIEW를 별도 컨텍스트 검토자에게 의뢰 예정.
