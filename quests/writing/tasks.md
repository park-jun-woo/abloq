# 집필 퀘스트 — 태스크 트리 (T1~T4)

이 트리는 래칫 상태가 아니다. 래칫(게이트)은 최종 제출물(글 + 작업 로그 + REVIEW 기록)만
판정한다. T1~T3은 집필 에이전트의 작업 순서, T4는 별도 컨텍스트 검토자의 작업이다.

## T1 — 자료 수집

- insight.yaml의 claims 중 `requires_source: true` 항목마다 출처 후보를 수집한다:
  실존 URL + 그 페이지에서 인용할 예정 문구.
- 출처 후보는 안정적인 공개 문서(표준 문서, 공식 문서, 원 논문)를 우선한다 —
  게이트(citation-exists)가 URL의 HTTP 200과 페이지 제목·인용 표기 겹침을 실검증한다.
- 수집 결과를 작업 로그에 기록한다: claim ID → 후보 URL → 인용 예정 문구.

## T2 — 초안

- insight.yaml의 topic/stance/audience/claims를 본문으로 전개한다.
  claims 순서 ≠ 본문 순서 — 글의 논리 흐름이 우선이다.
- `requires_source` claim을 뒷받침하는 문장에는 출처를 인라인 마크다운 링크로 배치한다
  (형식은 context.md 참조).
- blog.yaml `structure.order`의 구조(메인 이미지·섹션 순서·출처 섹션)를 따른다.
- front matter는 front-matter-schema가 요구하는 필드(title/date/lastmod/tags)를 채운다.

## T3 — 퇴고 (자기 점검)

- 논지 누락: insight.yaml의 모든 claim이 본문에서 다뤄졌는지 훑는다.
  `abloq insight match <insight.yaml> <글>`로 anchors 스크리닝을 돌려 미출현 목록을 확인한다.
- non_goals 범위 이탈: non_goals에 선언된 주제를 본문이 다루면 삭제한다
  (게이트 룰이 아니라 REVIEW 검토 대상 — T4에 넘기지 말고 여기서 끝낸다).
- 구조 순서: 섹션 헤딩 레벨(##)과 상대 순서를 점검한다.
- 작업 로그에 퇴고에서 바꾼 내용을 추가한다.

## T4 — REVIEW (별도 컨텍스트 검토자)

- **집필 에이전트는 자신의 글을 REVIEW할 수 없다** (context.md 금지 사항).
- 집필 에이전트는 검토자에게 다음을 제시한다:
  ⑴ 글 본문 ⑵ insight.yaml ⑶ `abloq insight match` 스크리닝 결과(미출현 claim 목록)
  ⑷ 작업 로그.
- 검토자는 별도 컨텍스트(다른 세션/프로세스)에서 판정하고 REVIEW 기록을 작성한다:
  - `reviewer:` 라인 — 검토 컨텍스트 식별자 (필수, 집필 컨텍스트와 달라야 한다)
  - 미출현 claim **전건**에 대한 disposition 라인 (필수):
    - `- <claim-id>: addressed — <본문 어디서 어떻게 대응했는지>`
    - `- <claim-id>: revised — <수정을 요구해 반영된 내용>`
    - `- <claim-id>: excluded — <제외 사유>`
  - 출현 claim의 지지 여부, non_goals 이탈 여부 소견 (자유 형식).
- 게이트의 review-record 룰은 기록 존재 + reviewer 라인 + 미출현 claim 전건의
  disposition 커버리지를 결정적으로 검사한다. 지지 판정 자체는 비결정이라 룰이 아니다.

## 제출

`abloq quest writing submit --key <key> --in <submission.json>` — submission.json:

```json
{
  "article": "content/<lang>/<section>/<slug>.md",
  "worklog": "quests/writing/logs/<slug>.md",
  "review": "quests/writing/reviews/<slug>.md"
}
```

경로는 모두 블로그 루트(blog.yaml 위치) 기준 상대 경로. article은 시드된 대상 경로와
일치해야 한다. FAIL이면 Fact(위치·기대·실제)가 돌아온다 — 수정 후 재제출.
**MaxTries=3**: FAIL 3회 누적 시 DONE으로 영구 잠금되니 Fact를 정확히 반영해라.
