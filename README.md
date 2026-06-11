# abloq

[![version](https://img.shields.io/badge/version-v0.1.0-blue)](https://github.com/park-jun-woo/abloq/releases)
[![license: MIT](https://img.shields.io/badge/license-MIT-green)](./LICENSE)

**Agentic blog Quest** — 에이전트가 운용하는 블로그.

사람은 인사이트(주제·주장·관점)만 결정한다. 자료 수집·집필·퇴고·발행·번역·GEO 운용·갱신은 에이전트가 **퀘스트**로 대행하고, 품질은 결정적 **게이트**가 보증한다.

## 왜 abloq인가

에이전트에게 블로그를 맡기면 글은 나온다. 문제는 **믿을 수 없다는 것**이다 — 출처를 날조하고, 고치지 않은 글의 lastmod를 올리고, 시키지 않은 파일을 건드린다. 사람이 전부 검수할 거면 맡긴 의미가 없다.

abloq의 답은 분업이다: **생성은 확률적, 검증은 결정론적.** 사람이 쓰는 것은 인사이트 명세 한 장뿐이다.

```yaml
# insight.yaml — 사람이 쓰는 전부
topic: "robots.txt와 Robots Exclusion Protocol — 30년 묵은 관행이 표준이 되기까지"
stance: "robots.txt는 접근 제어 장치가 아니라 신호다"
claims:
  - id: rep-standardized-2022
    text: "robots.txt 관행은 1994년에 시작됐지만 IETF 표준(RFC 9309)이 된 것은 2022년이다"
    requires_source: true
    anchors: ["RFC 9309", "1994"]
```

에이전트가 자료를 수집하고 집필해 제출하면, 게이트가 판정한다. 아래는 실제 운행 기록이다 — 에이전트가 출처 섹션을 빠뜨리고 기계 검증이 불가능한 인용을 넣자:

```
en/posts/robots-exclusion-protocol -> FAIL
  - min-sources: content/en/posts/robots-exclusion-protocol.md:1
    actual="sources section missing — geo.min_sources requires >= 1 source(s)"
  - citation-exists: content/en/posts/robots-exclusion-protocol.md:19
    actual="citation URL https://www.robotstxt.org/orig.html is not reachable (HTTP 403)"
```

FAIL은 의견이 아니라 **위치와 기대값이 박힌 사실(Fact)**이다. 에이전트는 이 피드백으로 수렴하고, 수정 제출이 전 룰을 통과해야만 기계가 PASS를 잠근다. 잠긴 PASS는 불가역이다 — 에이전트는 일회용이어도 진행은 누적된다.

집필만이 아니다. 같은 구조로 번역(N개 언어, 구조 무손실 검증), 갱신(낡은 글 검출 → 큐 → lastmod 정직성 강제), 근거 보강(죽은 링크·무출처 주장 검출 → 큐 밖 주장 변경 금지), 클러스터 큐레이션(내부 링크 그래프)이 돌고, CloudFront 로그·GSC·AI 응답 샘플링으로 **AI 인용 가시성을 측정해 다음 노동의 우선순위를 기계가 지정**한다. 측정이 노동을 지정하는 래칫 — GEO는 상태가 아니라 운용이다.

## 퀵스타트

에이전트(Claude Code 등)를 쓴다면 스킬 설치 한 번이면 끝이다:

```bash
npx skills add park-jun-woo/abloq
```

이후 에이전트에게 "abloq으로 블로그 만들어줘"라고 하면 된다 — [SKILL.md](./SKILL.md)가 라우팅하고 [MANUAL.md](./MANUAL.md)가 절차를 안내한다.

직접 쓴다면:

```bash
go install github.com/park-jun-woo/abloq/cmd/abloq@latest

abloq init my-blog      # blog.yaml + 템플릿 + CLAUDE.md, 게이트 클린 상태로 스캐폴드
cd my-blog
abloq generate .        # blog.yaml → hugo.toml·robots.txt·llms.txt·jsonld 파생
abloq check .           # 파생물 드리프트 검사 (CI 훅)
abloq gate .            # 글 본문에 구조·근거·정책 룰 적용 (위반 시 exit 1)
```

## 핵심 개념

- **blog.yaml = SSOT 한 장.** 블로그 하나의 전 선언(사이트·언어·섹션·글의 정규 구조·GEO 임계·배포)을 한 파일에 담는다. 여기서 `hugo.toml`·`robots.txt`·`llms.txt`·`sitemap`(hreflang)·JSON-LD·게이트 룰 파라미터가 파생된다. blog.yaml이 바뀌지 않는 한 어떤 글도 게이트를 우회할 수 없다 — 제약은 계약이다.

- **결정적 게이트 vs 비결정 퀘스트.** 에이전트는 산문을 만지는 비결정 노동에만 쓴다. 생성·스캔·측정·외부 API 호출·판정은 전부 결정적 코드([reins](https://github.com/park-jun-woo/reins) 게이트 + 운용 백엔드)다. 코드로 되는 일에 에이전트를 쓰면 비용·비결정성·치즈 공격면만 늘어난다.

- **운용 백엔드 abloqd (옵션).** 산문을 만지지 않는 일 — 파생물 생성·아카이브·스캔·측정·큐 — 은 상시 실행 백엔드 서비스(yongol 기반 Go+Gin, 셀프호스트)로 분리한다. 단, 전 모듈은 `abloq` CLI 서브커맨드로도 단독 실행 가능하다. 백엔드는 스케줄과 상태를 더한 것뿐이며, 정적 사이트의 단순성을 해치지 않는다.

- **가시성 3계층 측정 + 래칫.** AI 인용은 직접 관측이 안 되므로 프록시 3계층으로 측정한다 — (1) 크롤 계층: CloudFront 로그에서 AI봇 히트 집계(결정적), (2) 색인 계층: GSC 노출·클릭 추이(API, 결정적), (3) 인용 계층: 표준 질의 셋을 주기 실행해 AI 응답 내 인용을 추세로 기록(비결정, 게이트화하지 않음). 측정 결과는 우선순위 큐의 가중치가 되어 다음 퀘스트의 입력을 지정한다.

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
    quests/          퀘스트 정의 (reins Definition 5종)
    archive/         Wayback·IndexNow·GSC 색인 요청
    bots/ img/ ...   AI봇 UA·이미지 변환 등 도메인 유틸
  quests/            에이전트 퀘스트 팩 (writing·translation·refresh·evidence·cluster)
  template/          abloq init이 복제하는 블로그 템플릿 (임베드)
  backend/           abloqd — yongol SSOT 선언 (manifest·OpenAPI·DDL·SSaC·Rego·Hurl)
  deploy/backend/    셀프호스트 한 장 (Dockerfile + docker-compose.yaml)
  docs/              blog-yaml·insight-spec·operations 문서
```

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
| 1 | **translation** (번역) | 새 글 + 본문 실변경 | 구조 무손실(translation-parity) + 전 언어 slug 일치 + front matter 미러 + 빌드 0 |
| 2 | **refresh** (갱신) | 신선도 스캐너 큐 | 본문 실변경 + 구조 게이트 유지 · 본문 실변경 없는 lastmod 갱신 차단(`honest-lastmod`) |
| 3 | **evidence** (근거 보강) | 주장-출처 스캐너 큐 | 글당 출처 ≥ `min_sources` · 신규 인용 실재 검증 · 큐 밖 주장 변경 금지 |
| 4 | **cluster** (클러스터 큐레이션) | 클러스터 스캐너 큐 | 태그가 taxonomy SSOT에 존재 · 고아 태그 0 · 글당 내부 링크 ≥ `min_internal_links` · 고립 글 0 |

치즈 방어 원칙(전 퀘스트 공통): front matter 보존, 게이트 판정과 저장소 반영의 바이트 일치, 큐 아이템 범위 밖 파일 변경 금지, 본문 변경 퀘스트(2·3)는 기존 주장 삭제 금지(인사이트 명세 보존), 구조 변환 퀘스트(1)는 본문 무손실(multiset). 외부 부수효과(아카이브·색인)는 백엔드 영수증으로 처리하며, **에이전트는 외부 API를 직접 치지 않는다.**

## 로드맵 — v0.2.0 (계획)

도그푸드 #0(parkjunwoo.com 역이식) 검토에서 나온 채택 차단 요소와 운용 확장 요구를 묶은 다음 마일스톤이다.

- **llms.txt 큐레이션 등급화 + 옵트아웃.** `geo.llms_txt`를 `auto | manual | off` 또는 객체 선언으로 확장 — `manual`이면 손큐레이션본이 `generate`/`check`와 깨끗이 공존하고, `auto`는 기본 언어 단일 스코프·포지셔닝 헤더·핀 엔트리·섹션 라벨·요약 절단으로 큐레이션 등급이 된다.
- **OG 이미지 provider 선택형.** `abloq image og`에 AI 이미지 provider(Gemini 등)를 opt-in으로 추가 — blog.yaml에 사이트 공통 프롬프트·모델과 안(variant)별 프리셋을 선언하고, 기본 1안·선택 시 N개 안을 후보로 생성해 검토 후 채택한다. AI 생성물은 파생물이 아닌 1회성 자산 — `generate`/`check`에는 절대 들어가지 않는다(결정론 경계 유지).
- **abloqd 멀티사이트.** 백엔드 한 인스턴스가 여러 블로그를 운용 — 사이트 목록은 `sites.yaml` 선언 SSOT, API는 `/sites/{site}/…` 스코프, 전 도메인 데이터가 사이트 차원으로 격리된다. **이때 API 경로가 변경된다(호환 깨짐).**

## 문서

- 에이전트 운용 매뉴얼: [MANUAL.md](./MANUAL.md)
- 에이전트 스킬: [SKILL.md](./SKILL.md)
- blog.yaml 스키마: [docs/blog-yaml.md](./docs/blog-yaml.md)
- 인사이트 명세 작성: [docs/insight-spec.md](./docs/insight-spec.md)
- 백엔드 운영: [docs/operations.md](./docs/operations.md)

## 라이선스

abloq은 MIT 무료 오픈소스다. 프레임워크 본체·게이트 룰셋·운용 백엔드·퀘스트 팩 모두 같은 라이선스로 공개한다. 수익화 계획은 없다.

- 전문: [LICENSE](./LICENSE)
- 서드파티 저작권 고지: [NOTICE](./NOTICE)
- 게이트 엔진: [reins](https://github.com/park-jun-woo/reins) (MIT)
