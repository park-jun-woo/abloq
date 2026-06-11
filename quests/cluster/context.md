# 클러스터 큐레이션 퀘스트 — 컨텍스트 규약

cluster 큐(kind=cluster)를 소비하는 에이전트에게 주입되는 규약.
공통 소비 프로토콜(`_queue-protocol.md`)이 우선한다 — 여기는 큐레이션 고유 규약만.

## 큐레이션의 본질

- topical authority는 태그 어휘의 정합과 내부 링크 그래프의 연결성에서 온다.
  이 퀘스트는 **글 쪽**(front matter tags·본문 내부 링크)만 만진다 — blog.yaml
  `geo.taxonomy` 선언의 수정(어휘 정책 변경)은 사람 몫이고 queue-scope가 차단한다.
- 게이트는 pkg/scan/cluster를 대상 인스턴스에 재실행해 payload가 지정한 위반
  종류가 대상 글에서 소멸했는지 본다(cluster-resolved). 타 글의 위반은 무관하다.
- 스캔 경합으로 이미 해소된 아이템(다른 글의 커밋이 위반을 지운 경우)은 빈 diff로
  제출해도 PASS다 — 교착이 아니라 정상 소비이며, 큐 파일만 ② 커밋으로 지운다.

## 내부 링크 규약

- 링크는 **문맥에 맞는 앵커 문장**으로 추가한다 — 링크 덤프(목록 나열)는 그래프는
  채워도 리뷰에서 드러나는 저품질이다. 본문 흐름에 한 문장으로 끼워 넣어라.
- 내부 링크 경로는 기본 언어 서빙 규약을 따른다 (`/<section>/<slug>/` —
  default_lang_in_subdir=false 기준; 인스턴스 설정을 확인하라).
- 고립 해소(no-isolated-post)는 **들어오는** 링크가 필요하다 — payload candidates의
  글에 대상 글로 가는 링크를 추가한다. candidates 밖의 글은 queue-scope FAIL이다.

## lastmod 규약

- **후보 글의 앵커 한 줄 추가는 lastmod 불변경** — min_meaningful_diff 임계 미달
  변경에 lastmod를 올리면 honest-lastmod FAIL이다(미미 diff + lastmod 갱신 조합).
- 대상 글의 변경이 임계를 넘으면(드물다 — 태그 정리 + 링크 여러 개) lastmod 갱신
  후 번역 재동기화 절차를 따른다.

## 금지 사항 (치즈 방어)

1. **blog.yaml(taxonomy 포함) 수정 금지** — queue-scope가 차단한다.
2. **candidates 밖 글에 링크 심기 금지** — 허용 집합은 payload 유래뿐이다.
3. **태그를 전부 지워 위반 회피 금지** — tag-taxonomy는 통과해도 클러스터 그래프가
   빈약해진다; 스캔 재실행이 no-orphan-tag·링크 위반을 다시 잡는다.
4. **큐 파일 선삭제 금지** — 삭제는 ② 커밋(소비 신호)에서만.
