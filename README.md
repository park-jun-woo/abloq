# abloq

**v0.1.0** · MIT

**Agentic blog Quest** — 에이전트가 운용하는 블로그.

사람은 인사이트(주제·주장·관점)만 결정한다. 자료 수집·집필·퇴고·발행·번역·GEO 운용·갱신은 에이전트가 **퀘스트**로 대행하고, 품질은 결정적 **게이트**가 보증한다.

GEO 대행사·SaaS는 체크리스트와 대시보드(사람 노동)를 판다. abloq은 노동을 에이전트가 하고 판정을 게이트가 하는 박스를 MIT로 공개한다.

## 핵심 개념

- **blog.yaml = SSOT 한 장.** 블로그 하나의 전 선언(사이트·언어·섹션·글의 정규 구조·GEO 임계·배포)을 한 파일에 담는다. 여기서 `hugo.toml`·`robots.txt`·`llms.txt`·`sitemap`(hreflang)·JSON-LD·게이트 룰 파라미터가 파생된다. blog.yaml이 바뀌지 않는 한 어떤 글도 게이트를 우회할 수 없다 — 제약은 계약이다.

- **결정적 게이트 vs 비결정 퀘스트.** 분업 원칙은 단순하다. 에이전트는 산문을 만지는 비결정 노동에만 쓴다. 생성·스캔·측정·외부 API 호출·판정은 전부 결정적 코드([reins](https://github.com/park-jun-woo/reins) 게이트 + 운용 백엔드)다. 코드로 되는 일에 에이전트를 쓰면 비용·비결정성·치즈 공격면만 늘어난다.

- **운용 백엔드 abloqd (옵션).** 산문을 만지지 않는 일 — 파생물 생성·아카이브·스캔·측정·큐 — 은 상시 실행 백엔드 서비스([yongol](https://github.com/park-jun-woo) 기반 Go+Gin, 셀프호스트)로 분리한다. 단, 전 모듈은 `abloq` CLI 서브커맨드로도 단독 실행 가능하다. 백엔드는 스케줄과 상태를 더한 것뿐이며, 정적 사이트의 단순성을 해치지 않는다.

- **가시성 3계층 측정 + 래칫.** AI 인용은 직접 관측이 안 되므로 프록시 3계층으로 측정한다 — (1) 크롤 계층: CloudFront 로그에서 AI봇 히트 집계(결정적), (2) 색인 계층: GSC 노출·클릭 추이(API, 결정적), (3) 인용 계층: 표준 질의 셋을 주기 실행해 AI 응답 내 인용을 추세로 기록(비결정, 게이트화하지 않음). 측정 결과는 우선순위 큐의 가중치가 되어 다음 퀘스트의 입력을 지정한다. 측정이 다음 노동을 지정하는 이 구조가 운용 루프의 래칫이다.

## 구조 개요

```
abloq/
  cmd/abloq/         CLI 엔트리포인트 (init·generate·gate·check·scan·quest·archive·report 등)
  pkg/
    blogyaml/        blog.yaml 스키마·로더·검증 (SSOT)
    gen/             파생물 생성기 (hugo.toml·robots.txt·llms.txt·sitemap·JSON-LD)
    gate/            구조·근거·정책 게이트 룰셋 (reins 위에)
    content/         글 파싱 (front matter·섹션 구조)
    scan/            신선도·주장-출처·클러스터 스캐너 (큐 생성)
    visibility/      크롤 로그 집계·GSC 폴링·인용 샘플링·월간 리포트
    insight/         인사이트 명세 ↔ 본문 대조
    queueio/         quests/queue 읽기·쓰기
    quests/          퀘스트 팩 로딩
    archive/         Wayback·IndexNow·GSC 색인 요청
    bots/ img/ ...   AI봇 UA·이미지 변환 등 도메인 유틸
  quests/            에이전트 퀘스트 팩 (writing·translation·refresh·evidence·cluster)
  template/          abloq init이 복제하는 블로그 템플릿 (임베드)
  backend/           abloqd — yongol SSOT 선언 (manifest·OpenAPI·DDL·SSaC·Rego·Hurl)
  deploy/backend/    셀프호스트 한 장 (Dockerfile + docker-compose.yaml)
  docs/              insight-spec·operations 문서
```

## 빠른 시작

```bash
# 1. CLI 설치
go install github.com/park-jun-woo/abloq/cmd/abloq@latest

# 2. 새 블로그 스캐폴드 (blog.yaml + 템플릿 + CLAUDE.md, 게이트 클린 상태로 생성)
abloq init my-blog
cd my-blog

# 3. blog.yaml에서 파생물 생성 (hugo.toml·robots.txt·llms.txt·jsonld)
abloq generate .

# 4. 게이트 실행 — 위반 시 exit 1 (커밋·CI 훅에 건다)
abloq validate .     # blog.yaml 자체 검증
abloq generate .     # 파생물 재생성
abloq check .        # 파생물이 SSOT와 일치하는지 (드리프트 시 exit 1)
abloq gate .         # 글 본문에 구조·근거·정책 룰 적용
```

스캐폴드된 `CLAUDE.md`는 abloq이 생성한다 — 에이전트가 1급 사용자라는 선언이고, 설치 직후부터 Claude Code 같은 에이전트가 운영을 인수할 수 있다. 게시 절차는 곧 퀘스트 호출이다.

### 운용 백엔드 (옵션)

스케줄·상태(시계열·큐·영수증)가 필요하면 abloqd를 셀프호스트한다. PostgreSQL + abloqd 한 장이다.

```bash
# backend/ SSOT에서 Go+Gin 코드를 projection으로 생성
yongol generate backend/specs backend/arts

# 자격증명 작성 후 기동 (POSTGRES_PASSWORD·JWT_SECRET·BLOG_REPO_PATH 등은 env로만 주입)
docker compose -f deploy/backend/docker-compose.yaml up -d --build
```

백엔드를 켜지 않아도 모든 검출은 CLI로 돌아간다 — `abloq scan freshness`, `abloq scan evidence`, `abloq scan cluster`, `abloq report monthly` 등. blog.yaml과 abloqd manifest는 별개의 SSOT이며, abloqd는 마운트된 저장소의 blog.yaml을 읽어 언어·섹션·URL 계약과 GEO 임계를 파라미터로 쓴다.

## 에이전트 퀘스트 (5종)

산문을 만지는 비결정 노동만 퀘스트로 남긴다. 검출·생성·측정·외부 API는 스캐너·백엔드가 하고, 갱신·근거·클러스터 퀘스트는 스캐너가 떨어뜨린 `quests/queue/`의 큐를 소비한다. 각 퀘스트는 게이트로 닫힌다.

| # | 퀘스트 | 트리거 | 게이트(핵심) |
|---|---|---|---|
| 0 | **writing** (집필) | 사람의 인사이트 명세 | 인사이트 명세 각 항목이 본문에 대응(REVIEW) · 인용 실재 검증(URL 200 + 메타 일치) · 출처 ≥ `min_sources` · 무출처 수치 주장 0 |
| 1 | **translation** (번역) | 새 글 + 본문 실변경 | 구조 게이트 룰셋 + 전 언어 slug 일치 + front matter 스키마 + 빌드 0 |
| 2 | **refresh** (갱신) | 신선도 스캐너 큐 | 본문 실변경 + 구조 게이트 유지 · 본문 실변경 없는 lastmod 갱신 차단(`honest-lastmod`) |
| 3 | **evidence** (근거 보강) | 주장-출처 스캐너 큐 | 글당 출처 ≥ `min_sources` · 신규 인용 실재 검증 · 큐 밖 주장 변경 금지 |
| 4 | **cluster** (클러스터 큐레이션) | 클러스터 스캐너 큐 | 태그가 taxonomy SSOT에 존재 · 고아 태그 0 · 글당 내부 링크 ≥ `min_internal_links` · 고립 글 0 |

치즈 방어 원칙(전 퀘스트 공통): front matter 보존, 게이트 판정과 저장소 반영의 바이트 일치, 큐 아이템 범위 밖 파일 변경 금지, 본문 변경 퀘스트(2·3)는 기존 주장 삭제 금지(인사이트 명세 보존), 구조 변환 퀘스트(1)는 본문 무손실(multiset). 외부 부수효과(아카이브·색인)는 백엔드 영수증으로 처리하며, **에이전트는 외부 API를 직접 치지 않는다.**

GEO는 상태가 아니라 운용이다. 큐 소비 퀘스트(2·4)가 본체다.

## 라이선스

abloq은 MIT 무료 오픈소스다. 프레임워크 본체·게이트 룰셋·운용 백엔드·퀘스트 팩 모두 같은 라이선스로 공개한다. 수익화 계획은 없다.

- 전문: [LICENSE](./LICENSE)
- 서드파티 저작권 고지: [NOTICE](./NOTICE)
- 게이트 엔진: [reins](https://github.com/park-jun-woo/reins) (MIT)
