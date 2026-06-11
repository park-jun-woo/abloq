# 갱신 퀘스트 — 컨텍스트 규약

freshness 큐(kind=refresh)를 소비하는 에이전트에게 주입되는 규약.
공통 소비 프로토콜(`_queue-protocol.md`)이 우선한다 — 여기는 갱신 고유 규약만.

## 갱신의 본질

- 갱신은 **낡은 수치·사실의 교체**다. lastmod만 올리는 빈 갱신(빈 diff)은
  lastmod-advance + honest-lastmod 조합이 차단하고, 주장 삭제로 분량을 줄이는
  우회는 claim-preserved(건수 하한)가 차단한다.
- 새로 단 수치 주장은 같은 문단에 출처 링크 필수(numeric-claim-sourced — 기준선
  대비 신규 주장만 검사). 인용 URL은 실재해야 한다(citation-exists).

## lastmod 규약

- `lastmod`는 기준선(git HEAD)보다 **미래**여야 한다(lastmod-advance). 오늘 날짜로
  올리는 게 정상이다.
- lastmod 갱신은 본문 실변경(`geo.min_meaningful_diff` 토큰 diff 이상)과 큐 등재를
  요구한다(honest-lastmod) — 이 퀘스트의 아이템은 큐 파일에서 왔으므로 등재는
  이미 충족돼 있다. 큐 파일을 게이트 전에 지우면 등재가 깨진다 — 삭제는 ② 커밋.

## 주장(claim) 처리

- 낡은 수치 주장은 **그 라인에서 새 수치로 교체**한다 — 문단 리랩 금지
  (claim 룰은 라인 단위 비교다).
- 주장 건수는 기준선 이상 유지(claim-preserved). 교체는 1:1이 자연스럽고,
  보강(신규 주장 추가)은 자유다 — 단 출처 필수.

## 금지 사항 (치즈 방어)

1. **빈 갱신 금지** — 본문 무변경 + lastmod 전진은 honest-lastmod FAIL이다.
2. **주장 삭제로 해소 금지** — claim-preserved가 건수 감소를 FAIL로 막는다.
3. **front matter 변조 금지** — lastmod 외의 키 변경은 front-matter-intact FAIL.
4. **범위 밖 파일 변경 금지** — queue-scope가 변경 집합을 기계 검사한다.
5. **본문 통삭제·재작성 금지** — 갱신은 부분 교체다. 인식 섹션을 떨어뜨리면
   section-preserved FAIL. (body-lossless는 이 퀘스트에서 배제 — 기존 라인 수정이
   갱신의 본질이라 무손실 비교와 모순이다.)
