# 큐 소비 공통 프로토콜 (Phase018)

백엔드 exporter가 떨어뜨린 `quests/queue/*.yaml`을 소비하는 퀘스트
(refresh·evidence·cluster) 전부에 적용되는 규약. 큐 파일 1개 = 퀘스트 아이템 1개이고,
Seed는 priority 내림차순으로 정렬한다(백엔드 스코어러가 매긴 운용 순서).

## 순서 박제 (제출·커밋 순서 — 절대 바꾸지 마라)

1. **작업트리에서 글 수정** — 커밋하지 않은 상태로 둔다.
2. **submit → PASS** — 게이트는 작업트리(더티) vs git HEAD 기준선으로 판정한다.
   글 수정을 먼저 커밋하면 작업트리==HEAD가 되어 기준선 룰 전체(honest-lastmod·
   claim 룰·queue-scope 변경 집합)가 공허 통과한다 — 게이트 무력화이므로 금지.
3. **① 글 수정 커밋** — PASS 이후에만.
4. **(해당 시) 번역 재동기화** — lastmod가 갱신됐다면 translation quest로 전 언어
   재동기화 후 커밋들. 큐 파일의 `keys:`가 전 언어 키를 동반 발급하므로 번역 커밋도
   honest-lastmod 큐 등재 검사를 통과한다.
5. **② 큐 파일 삭제 커밋(소비 신호)** — 반드시 마지막. ②를 번역 재동기화보다
   앞당기면 repo/CI honest-lastmod가 번역 커밋을 재차단한다(Phase017 D4).
   푸시 후 다음 `POST /queue/export` 회전이 삭제를 status=consumed로 동기화한다.

## queue-scope 규약

- 작업트리 변경 파일 집합(`git status --porcelain`, untracked 포함)은 허용 집합에
  들어가야 한다: **대상 글 + 그 전 언어 번역본 경로 + insight 사이드카** (+ cluster
  kind는 payload candidates의 글 경로 추가).
- **큐 파일 자신은 허용 집합 밖이다** — 게이트 시점에 큐 파일은 무변경이어야 하고,
  삭제는 ② 커밋에서만 한다. 게이트 전에 지우거나 고치면 queue-scope FAIL이다.
- blog.yaml·레이아웃·다른 글 수정은 전부 범위 밖 — 사람(운영자)의 몫이다.

## 주장(claim) 라인 규약

- claim-scope/claim-preserved는 **라인 단위** 비교다 — 주장이 든 문단을 리랩(줄바꿈
  재배치)하면 보존된 주장도 변경으로 검출돼 거짓 FAIL이 난다. **주장 라인은
  리랩하지 마라.**
- 큐 payload의 claims 해시 목록에 **없는** 주장은 한 글자도 바꾸지 마라(claim-scope).
- 주장을 삭제로 "해소"하지 마라 — refresh는 건수 하한(claim-preserved)이 막고,
  evidence는 출처 추가가 본질이다.

## 치즈 방어 공통 원칙

1. **payload는 Seed 시점에 고정된다** — 작업트리 큐 파일을 고쳐도 게이트가 보는
   payload는 변하지 않는다.
2. **기준선은 git HEAD다** — 대상 글이 HEAD에 없으면 제출 자체가 중단된다(에러,
   try 미소진). 인스턴스를 먼저 커밋 상태로 만들어라.
3. **빈 diff 무작업 통과는 막혀 있다** — 각 퀘스트의 작업 완수 강제 룰(lastmod 전진,
   해소 재검)이 "아무것도 안 하고 PASS"를 차단한다.
4. **MaxTries=3** — FAIL 3회 누적 시 DONE으로 영구 잠금된다. Fact를 정확히 반영해라.
