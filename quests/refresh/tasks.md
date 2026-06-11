# 갱신 퀘스트 — 태스크 트리 (T1~T3)

이 트리는 래칫 상태가 아니다. 래칫(게이트)은 최종 제출물(갱신된 글 1편)만 판정한다.
아이템은 freshness_days를 초과한 글 1편 — payload의 `lastmod`(현재 값)와
`freshness_days`(임계)가 발급 근거다.

## T1 — 낡음 진단

- 글의 수치 주장·날짜 언급·제품 버전·외부 링크를 훑고 **시점 의존 정보** 목록을
  뽑는다 (예: "2024년 기준", 가격, 점유율, API 버전).
- 각 항목의 현재 사실을 재수집한다 — 출처 URL은 실재(HTTP 200)해야 한다
  (citation-exists가 신규 인용을 기계 검증한다).

## T2 — 갱신

- 낡은 수치는 **교체**한다 — 삭제가 아니다. 수치 주장 건수는 기준선 이상이어야
  한다(claim-preserved). 새 수치 주장에는 같은 문단에 출처 링크를 단다
  (numeric-claim-sourced).
- 본문에 의미 있는 변경(`geo.min_meaningful_diff` 토큰 이상)을 만든다 —
  honest-lastmod가 lastmod 갱신과 본문 실변경의 짝을 검사한다.
- front matter는 **lastmod만** 전진시킨다(front-matter-intact·lastmod-advance).
  date·title·tags·slug는 그대로.
- 섹션 구조(인식 섹션의 존재·순서)는 보존한다(section-preserved·section-order).

## T3 — 자기 점검과 제출

- lastmod가 기준선보다 미래인지, 본문 diff가 임계 이상인지, 주장 건수가 줄지
  않았는지, 변경 파일이 대상 글뿐인지 확인한다.
- (선택) 의미 수준 보존 REVIEW: 주장 교체의 타당성(낡은 수치 → 올바른 새 수치)은
  비결정이라 게이트 룰이 아니다. 원하면 별도 컨텍스트 검토자에게 기준선·갱신본
  쌍을 제시해 소견을 받아라 — 기록 강제 없음.

## 제출

`abloq quest refresh submit --key <key> --in <submission.json>` — submission.json:

```json
{
  "article": "content/<lang>/<section>/<slug>.md"
}
```

경로는 블로그 루트 기준 상대 경로이고 시드된 대상 경로와 일치해야 한다.
제출은 **글 수정 커밋 이전**(작업트리 더티 상태)에 한다 — `_queue-protocol.md`의
순서 박제를 지켜라. FAIL이면 Fact(위치·기대·실제)가 돌아온다 — 수정 후 재제출.
**MaxTries=3**: FAIL 3회 누적 시 DONE으로 영구 잠금되니 Fact를 정확히 반영해라.
