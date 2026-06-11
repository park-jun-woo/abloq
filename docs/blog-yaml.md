# blog.yaml — 스키마 레퍼런스 (v1)

블로그 하나의 전 선언을 담는 SSOT. 모든 파생물(hugo.toml · robots.txt · llms.txt · sitemap · JSON-LD · 게이트 파라미터)이 이 한 장에서 나온다. **blog.yaml이 바뀌지 않는 한 어떤 글도 게이트를 우회할 수 없다.**

파싱은 strict — 스키마에 없는 키는 `unknown-key` 에러다. 키 추가는 마이너 버전, 의미 변경은 메이저 버전으로만 한다.

## 검증

```bash
abloq validate [dir]          # dir/blog.yaml 검증 (기본: 현재 디렉토리)
abloq validate --json [dir]   # 진단을 JSON 배열로 출력
```

진단 형식: `파일:라인 [룰ID] 메시지`. 진단이 있으면 exit code 1.

## 키

| 키 | 타입 | 필수 | 기본값 | 설명 |
|---|---|---|---|---|
| `site.baseURL` | string | ✅ | — | 절대 http(s) URL, query/fragment 금지 |
| `site.title` | string | | — | 사이트 제목 |
| `site.author` | string | | — | 저자명 |
| `site.default_lang_in_subdir` | bool | | `true` | `false`면 기본 언어를 사이트 루트(`/...`)에 서빙 — hugo `defaultContentLanguageInSubdir`, llms.txt URL, hreflang 페이지 경로, `.md` 병행 서빙 경로가 함께 따른다 |
| `languages` | [string] | ✅ | — | BCP-47 언어 코드. **첫 항목 = 기본 언어** |
| `sections` | [string] | ✅ | — | 디렉토리 기반 섹션 (1개 이상) |
| `structure.order` | [string] | | — | 글의 정규 섹션 순서 = 구조 게이트 룰의 입력 |
| `structure.headings` | map | | — | 헤딩 키 → 언어 코드 → 현지화 헤딩. 기본 언어 항목 필수 |
| `geo.crawlers` | map | | `{}` | 크롤러 분류/봇 이름 → `allow` \| `block` |
| `geo.llms_txt` | string \| object | | `auto` | llms.txt 생성 선언 — 문자열 단축형(`auto` \| `manual` \| `off`) 또는 아래 객체 폼 |
| `geo.llms_txt.mode` | string | | `auto` | `auto`(자동 생성·드리프트 게이트) \| `manual`(손큐레이션 — generate/check 불간섭) \| `off`(미생성) |
| `geo.llms_txt.languages` | string \| [string] | | `base` | 언어 스코프 — `base`(기본 언어 1개) \| `all`(전 언어) \| 선언된 언어의 부분집합(예: `[en, ko]`) |
| `geo.llms_txt.header` | string | | — | 사이트 포지셔닝 블록(자유 마크다운) — 목록 위에 그대로 삽입 |
| `geo.llms_txt.pinned` | [object] | | — | 선두 고정 엔트리 — `title`·`url` 필수(절대 URL 또는 `/` 시작), `desc`·`group` 선택. `group`이 섹션 그룹 헤딩과 일치하면 그 그룹 선두에, 아니면 자체 헤딩 그룹으로, 미지정이면 최상단 무헤딩 |
| `geo.llms_txt.section_labels` | map | | — | 섹션 → 사람이 읽는 그룹 라벨 (키는 선언된 섹션만) |
| `geo.llms_txt.max_summary` | int | | `0` | 항목 설명문 길이 상한(rune) — 초과 시 절단 + `…`, `0` = 무제한 |
| `geo.jsonld` | [string] | | `[Article, Person]` | 생성할 JSON-LD 타입 |
| `geo.freshness_days` | int | | `90` | 신선도 퀘스트 임계 (≥ 1) |
| `geo.min_sources` | int | | `1` | 근거 게이트 임계 (≥ 0) |
| `geo.min_internal_links` | int | | `2` | 클러스터 게이트 임계 (≥ 0) |
| `geo.min_meaningful_diff` | int | | `10` | honest-lastmod 임계 — 공백·구두점 정규화 후 토큰 diff가 이 값 미만이면 lastmod 갱신 거부 (≥ 1) |
| `deploy.provider` | string | | `s3-cloudfront` | 배포 대상 |
| `deploy.terraform` | bool | | `false` | IaC 생성 여부 |
| `deploy.indexnow` | bool | | `true` | 배포 후 IndexNow 핑 |

## 검증 룰

| 룰ID | 판정 |
|---|---|
| `yaml-syntax` | YAML 문법 오류 · 빈 파일 · 타입 불일치 |
| `unknown-key` | 스키마에 없는 키 (strict 파싱) |
| `lang-bcp47` | `languages` 비어있지 않음 + 각 항목 BCP-47 유효 |
| `heading-default-lang` | `structure.headings`의 각 헤딩 키에 기본 언어 항목 존재 |
| `sections-empty` | `sections` 1개 이상 |
| `threshold-range` | `freshness_days ≥ 1`, `min_sources ≥ 0`, `min_internal_links ≥ 0`, `min_meaningful_diff ≥ 1` |
| `baseurl-format` | 절대 http(s) URL, host 존재, query/fragment 없음 |
| `crawlers-policy` | `geo.crawlers` 값이 `allow` \| `block` |
| `llmstxt-mode` | `geo.llms_txt` mode가 `auto` \| `manual` \| `off` |
| `llmstxt-languages` | `geo.llms_txt.languages`가 `base` \| `all` \| 선언된 `languages`의 부분집합 |
| `llmstxt-pinned` | pinned 각 엔트리에 `title`·`url` 존재, `url`은 절대 http(s) URL 또는 `/` 시작 |
| `llmstxt-labels` | `geo.llms_txt.section_labels` 키가 선언된 섹션 |
| `llmstxt-max-summary` | `geo.llms_txt.max_summary ≥ 0` |

## llms.txt 생성 모드

- `auto`(기본) — `abloq generate`가 `static/llms.txt`를 렌더하고 `abloq check`가 드리프트를 잡는다. 기본 언어 스코프는 `base`(정준 언어 가이드) — 전 언어 덤프는 `languages: all`로 옵트인.
- `manual` — 손큐레이션 모드. generate가 파일을 건드리지 않고 check도 강제하지 않는다 (`static/llms.txt`는 사람/에이전트 소유).
- `off` — llms.txt를 만들지 않는다. **주의:** `auto`에서 `manual`/`off`로 전환해도 기존에 생성된 `static/llms.txt`는 abloq가 지우지 않는다 — Build 목록 제외 방식이라 잔존 파일에 check가 침묵하므로, `off` 전환 시 기존 파일은 직접 삭제하라.

## 예제

```yaml
site:
  baseURL: https://parkjunwoo.com
  title: Junwoo Park
  author: Junwoo Park

languages: [ko, en, ja]          # ko = 기본 언어
sections: [opinion, tech]

structure:
  order: [image, attribution, body, related, further, sources, changelog]
  headings:
    related: { ko: "관련 글", en: "Related", ja: "関連記事" }
    sources: { ko: "출처", en: "Sources", ja: "出典" }

geo:
  crawlers: { training: allow, search: allow, fetch: allow, bytespider: block }
  llms_txt: auto                 # 단축형. 또는 객체 폼:
  # llms_txt:
  #   mode: auto                 # auto | manual | off
  #   languages: base            # base | all | [en, ko]
  #   header: |
  #     This site publishes ...
  #   pinned:
  #     - title: Master Index
  #       url: /reins.md
  #       desc: Index of all articles
  #       group: Core Content
  #   section_labels: { opinion: Concept, tech: Pattern }
  #   max_summary: 200
  jsonld: [Article, Person]
  freshness_days: 90
  min_sources: 1
  min_internal_links: 2
  min_meaningful_diff: 10

deploy:
  provider: s3-cloudfront
  terraform: true
  indexnow: true
```

전체 12언어 골든 예제: `pkg/blogyaml/testdata/valid/blog.yaml`

## Go API (`pkg/blogyaml`)

```go
b, diags, err := blogyaml.Load("blog.yaml")   // err = IO 에러만, 스키마 문제는 diags
b, idx, diags := blogyaml.Parse(name, data)   // strict 디코드 + 기본값 주입
diags := blogyaml.Validate(name, b, idx)      // 검증 룰 실행
```
