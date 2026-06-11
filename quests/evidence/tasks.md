# 근거 보강 퀘스트 — 태스크 트리 (T1~T3)

이 트리는 래칫 상태가 아니다. 래칫(게이트)은 최종 제출물(보강된 글 1편)만 판정한다.
아이템은 무출처 수치 주장(payload `claims`) 또는 확정 link rot(payload `rot_urls`)이
검출된 글 1편이다.

## T1 — 검출 내역 확인

- payload `claims`의 각 항목(hash·loc·text)을 글에서 찾는다 — `loc`(파일:라인)은
  힌트일 뿐이고 `text`(주장 원문)가 기준이다.
- payload `rot_urls`의 각 URL이 글의 어느 인용에 쓰였는지 찾는다.

## T2 — 보강

- **무출처 주장**: 주장을 뒷받침하는 실재 출처(HTTP 200)를 찾아 **같은 문단**에
  인라인 링크로 단다. 출처를 못 찾으면 주장을 **사실로 다듬어 재서술**하거나
  출처 있는 다른 수치로 교체한다 — 단 큐에 없는 주장은 건드리지 마라(claim-scope).
- **rot URL**: 살아 있는 대체 출처(같은 내용의 새 URL, 아카이브 사본 등)로
  교체한다. rot URL이 글에 남아 있으면 rot-resolved FAIL이다.
- 신규 인용 URL은 게이트가 실재·제목 일치를 검증한다(citation-exists).

## T3 — 자기 점검과 제출

- payload claims의 주장 전부가 출처를 갖거나 정당하게 교체됐는지, rot URL이
  본문 인용에서 사라졌는지, 큐 밖 주장이 한 글자도 변하지 않았는지 확인한다.
- lastmod 정책: 보강이 `geo.min_meaningful_diff` 이상의 실변경이면 lastmod를
  갱신하고(이후 번역 재동기화), 임계 미달이면 lastmod를 **올리지 마라**
  (honest-lastmod 정직 규약 — 미미한 변경에 lastmod를 올리면 repo 게이트 FAIL).

## 제출

`abloq quest evidence submit --key <key> --in <submission.json>` — submission.json:

```json
{
  "article": "content/<lang>/<section>/<slug>.md"
}
```

경로는 블로그 루트 기준 상대 경로이고 시드된 대상 경로와 일치해야 한다.
제출은 **글 수정 커밋 이전**(작업트리 더티 상태)에 한다 — `_queue-protocol.md`의
순서 박제를 지켜라. FAIL이면 Fact(위치·기대·실제)가 돌아온다 — 수정 후 재제출.
**MaxTries=3**: FAIL 3회 누적 시 DONE으로 영구 잠금되니 Fact를 정확히 반영해라.
