# 근거 보강 퀘스트 — 컨텍스트 규약

evidence 큐(kind=evidence)를 소비하는 에이전트에게 주입되는 규약.
공통 소비 프로토콜(`_queue-protocol.md`)이 우선한다 — 여기는 보강 고유 규약만.

## 보강의 본질

- 보강은 **출처 추가/교체**다 — 주장을 지우거나 검출을 피하게 비틀어 "해소"하는
  것이 아니다. 게이트는 Doc에 무출처 주장 검출을 재실행해 payload claims 해시와
  교집합이 남으면 FAIL(claims-resolved), rot URL이 인용에 남으면 FAIL(rot-resolved).
- 주장 리워딩(수치를 바꿔 해시를 벗어나는 우회)은 numeric-claim-sourced가 잡는다 —
  기준선(git HEAD) 대비 신규 주장으로 검출되므로 출처 없는 리워딩은 FAIL이다.

## 출처 규약

- 인라인 마크다운 링크(`[표기](https://...)`)를 주장과 **같은 문단**에 단다 —
  문단 단위 출처 보유 판정(Phase010 검출기와 동일)이다.
- 신규 인용 URL은 HTTP 200 + 제목 대조를 통과해야 한다(citation-exists).
  접근 차단(403)·리다이렉트 만료 URL은 출처로 못 쓴다.
- sources 섹션 최소 건수(`geo.min_sources`)를 유지한다(min-sources).

## claim-scope (큐에 없는 주장 변경 금지)

- payload `claims`의 해시 목록이 **변경 허가 목록**이다 — 목록 밖 주장 라인은
  한 글자도 바꾸지 마라(텍스트 multiset 비교, 라인 단위).
- 주장 라인 리랩 금지 — 문단을 재배치하면 보존된 주장도 변경으로 검출된다.

## lastmod 정책 (이 퀘스트는 lastmod를 강제하지 않는다)

- 실변경(`geo.min_meaningful_diff` 이상)이면 lastmod 갱신 + ① 커밋 후 번역
  재동기화(translation quest). 큐 파일 `keys:`가 전 언어를 동반 발급한다.
- 임계 미달 변경(출처 링크 한두 개)이면 lastmod **불변경** — keys 불필요,
  번역 재동기화 미발생. honest-lastmod의 정직 규약과 정합이다.

## 금지 사항 (치즈 방어)

1. **주장 삭제·은닉 금지** — claims-resolved는 해소를 검증하지만 claim-scope가
   큐 밖 주장 보존을, numeric-claim-sourced가 신규/리워딩 주장 출처를 잡는다.
2. **claims_ignore 남용 금지** — 예외는 사유 필수이고 repo 게이트·리뷰에 드러난다.
3. **rot URL 방치 금지** — 인용에서 제거(교체)해야 rot-resolved가 통과한다.
4. **범위 밖 파일 변경 금지** — queue-scope가 변경 집합을 기계 검사한다.
