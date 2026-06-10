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
| `languages` | [string] | ✅ | — | BCP-47 언어 코드. **첫 항목 = 기본 언어** |
| `sections` | [string] | ✅ | — | 디렉토리 기반 섹션 (1개 이상) |
| `structure.order` | [string] | | — | 글의 정규 섹션 순서 = 구조 게이트 룰의 입력 |
| `structure.headings` | map | | — | 헤딩 키 → 언어 코드 → 현지화 헤딩. 기본 언어 항목 필수 |
| `geo.crawlers` | map | | `{}` | 크롤러 분류/봇 이름 → `allow` \| `block` |
| `geo.llms_txt` | string | | `auto` | llms.txt 생성 방식 |
| `geo.jsonld` | [string] | | `[Article, Person]` | 생성할 JSON-LD 타입 |
| `geo.freshness_days` | int | | `90` | 신선도 퀘스트 임계 (≥ 1) |
| `geo.min_sources` | int | | `1` | 근거 게이트 임계 (≥ 0) |
| `geo.min_internal_links` | int | | `2` | 클러스터 게이트 임계 (≥ 0) |
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
| `threshold-range` | `freshness_days ≥ 1`, `min_sources ≥ 0`, `min_internal_links ≥ 0` |
| `baseurl-format` | 절대 http(s) URL, host 존재, query/fragment 없음 |
| `crawlers-policy` | `geo.crawlers` 값이 `allow` \| `block` |

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
  llms_txt: auto
  jsonld: [Article, Person]
  freshness_days: 90
  min_sources: 1
  min_internal_links: 2

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
