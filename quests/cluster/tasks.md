# 클러스터 큐레이션 퀘스트 — 태스크 트리 (T1~T3)

이 트리는 래칫 상태가 아니다. 래칫(게이트)은 최종 제출물(큐레이션된 글들)만 판정한다.
아이템은 클러스터 위반(payload `violations`)이 검출된 기본 언어 글 1편이고, payload
`candidates`가 연결 후보(공유 태그·섹션·날짜 근접 순위)를 제안한다.

## T1 — 위반 확인

- payload `violations`의 각 항목(rule·detail)을 읽는다: `tag-taxonomy`(태그가 선언
  taxonomy 밖), `no-orphan-tag`(전 코퍼스에서 1회뿐인 태그), `min-internal-links`
  (내부 링크 부족), `no-isolated-post`(들어오는 링크 0).
- payload `candidates`에서 연결할 글을 고른다 — 순위가 제안 근거다(shared_tags·
  같은 섹션·날짜 근접).

## T2 — 큐레이션

- **태그 위반**: 대상 글 front matter의 `tags`를 고친다 — taxonomy 안의 태그로
  교체하거나 다른 글과 공유되는 태그로 정리한다. (blog.yaml의 taxonomy 선언 수정은
  사람 몫 — queue-scope가 차단한다.)
- **링크 부족(out)**: 대상 글 본문에 후보 글로 가는 내부 링크를 문맥에 맞는
  앵커 문장으로 추가한다.
- **고립(no-isolated-post, in)**: 후보 글 쪽에 대상 글로 들어오는 링크를 추가한다 —
  payload candidates의 글만 허용된다(queue-scope +candidates).
- **후보 글의 앵커 한 줄 추가는 lastmod 불변경** — `geo.min_meaningful_diff` 임계
  미달 변경은 lastmod를 올리지 않는 게 정직 규약이다(honest-lastmod 정합,
  번역 재동기화 미발생).

## T3 — 자기 점검과 제출

- 클러스터 스캔을 머리로 재실행한다: payload가 지정한 위반 종류가 대상 글에서
  사라졌는가? (게이트가 pkg/scan/cluster를 실제로 재실행해 검증한다.)
- 변경 파일이 대상 글(+전 언어 동반)과 candidates 글뿐인지 확인한다.

## 제출

`abloq quest cluster submit --key <key> --in <submission.json>` — submission.json:

```json
{
  "article": "content/<lang>/<section>/<slug>.md"
}
```

경로는 블로그 루트 기준 상대 경로이고 시드된 대상 경로와 일치해야 한다.
제출은 **글 수정 커밋 이전**(작업트리 더티 상태)에 한다 — `_queue-protocol.md`의
순서 박제를 지켜라. FAIL이면 Fact(위치·기대·실제)가 돌아온다 — 수정 후 재제출.
**MaxTries=3**: FAIL 3회 누적 시 DONE으로 영구 잠금되니 Fact를 정확히 반영해라.
