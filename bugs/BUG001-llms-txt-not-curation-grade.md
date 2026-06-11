# BUG001 — llms.txt 자동본이 큐레이션 등급이 아니다 (+ 옵트아웃 불가)

> 상태: OPEN · 심각도: High(채택 차단) · 발견: parkjunwoo.com 본번 적용 검토(도그푸드 #0 후속)

## 한 줄 요약

`abloq generate`의 llms.txt 자동본은 **사이트맵 등급 덤프**라 손큐레이션 llms.txt를 대체하면 GEO가 후퇴한다. 게다가 `geo.llms_txt` 옵션이 **죽은 설정**이라 자동 생성을 끄고 큐레이션본을 유지할 방법조차 없다 — abloq 채택이 막힌다.

## 증상

실코퍼스(parkjunwoo.com, 12언어·732편)에서 현행 **손큐레이션 llms.txt**와 **abloq 자동본**을 비교:

| | 현행 큐레이션 | abloq 자동본 |
|---|---|---|
| 길이 | **65줄** | **879줄** (13.5×) |
| 언어 | 영어(정준) 단일 | 12개 언어 전부 |
| 구조 | `Core Content → Concept → Pattern → Toolchain` 편집 위계 | `## {lang}/{section}` 평면 그룹(언어×섹션 60개) |
| 선두 | 마스터 인덱스 `reins.md` 링크 | 알파벳순 첫 그룹 `en/opinion`의 최신글 |
| 헤더 | 사이트 포지셔닝 문단(무엇을 하는 사이트인지) | `> PARK JUN WOO — https://...` (저자—baseURL 한 줄) |
| 설명문 | 항목별 손요약(짧고 선별) | front matter `summary` 전문(긴 문장) |

llms.txt의 존재 이유는 LLM에게 **간결·우선순위화된** 안내를 주는 것인데, 자동본은 *전 언어 × 전 글*을 알파벳순 평면 나열해 사실상 sitemap에 가깝다. 가장 중요한 진입점(마스터 인덱스)이 879줄 어딘가에 묻힌다.

## 재현

```bash
# parkjunwoo 인스턴스(hugo/) 기준
abloq generate hugo
wc -l hugo/static/llms.txt          # → 879줄 (큐레이션본은 65줄)
head -5 hugo/static/llms.txt        # 헤더가 "> author — baseURL" 한 줄뿐, 첫 그룹이 en/opinion
```

## 근본 원인 (두 겹)

### 원인 1 — `geo.llms_txt`가 죽은 설정 (옵트아웃 불가)

- `pkg/blogyaml/geo.go:16` — `LlmsTxt string` 필드 파싱.
- `pkg/blogyaml/default_blog.go:14` — 기본값 `"auto"`.
- **소비처가 없다.** `grep -rn "LlmsTxt\|Geo.Llms" pkg cmd --include=*.go` 결과 위 두 줄(선언·기본값)뿐, 값으로 분기하는 코드 0.
- `pkg/gen/build.go:19` — `Build()`가 값과 무관하게 `static/llms.txt`를 **항상** 출력 목록에 넣는다.
- `cmd/abloq/run_check.go:18` — `gen.Check(dir, gen.Build(dir, b))`. check가 기대 파일을 `gen.Build`로 열거하므로, llms.txt가 항상 `llmstxt-sync` 룰로 강제된다.

→ 결과: 빌드 파이프라인에 `generate`/`check`를 물리면 큐레이션본이 **매번 자동본으로 덮어쓰이거나**, 안 물리면 abloq의 드리프트 게이트(핵심 가치)가 깨진다. 중간이 없다.

### 원인 2 — `auto` 렌더러가 llms.txt 등급이 아님

`pkg/gen/llms/render.go` `Render()`:
- 헤더 = `# {Title}` + `> {headerNote}` 뿐. `headerNote`(`header_note.go`)는 `author — baseURL`만 — 사이트가 **무엇을 하는지**(포지셔닝)가 없다.
- `Collect()`(`collect.go`)가 `b.Languages` 전부를 순회 → 전 언어가 한 파일에. 언어 스코프 개념 없음.
- 그룹 헤딩이 `## {lang}/{section}` 기계 라벨, 정렬은 `언어(선언순)→섹션(선언순)→날짜 desc`(`sort_posts.go`). **중요도 순서·핀 엔트리·사람 라벨 없음.**
- `postLine()`이 `summary` 전문을 그대로 → 길다.

## 기대 동작 / 개선안

`geo.llms_txt`를 **문자열 단축형 또는 객체**로 확장한다(하위호환: 문자열 `auto` 유지).

```yaml
geo:
  llms_txt: auto            # 단축형: auto | manual | off  (하위호환)
  # 또는 객체:
  llms_txt:
    mode: auto              # auto | manual | off
    languages: base         # base(기본 언어 1개) | all | [en, ko]   (기본 base)
    header: |               # 선택: 사이트 포지셔닝 블록(자유 마크다운, 목록 위에 삽입)
      This site publishes ... Reins Engineering ...
    pinned:                 # 선택: 선두 고정 엔트리(마스터 인덱스 등)
      - title: Reins Engineering
        url: /reins.md
        desc: Index of all Reins Engineering articles and tools
        group: Core Content
    section_labels:         # 선택: 섹션 → 사람이 읽는 그룹 라벨
      opinion: Concept
      tech: Pattern
    max_summary: 200        # 선택: 설명문 길이 상한(0=무제한)
```

### Tier 0 — 옵트아웃 (최소, 채택 차단 해소)

`mode: manual|off`면 `Build()` 출력 목록에서 llms.txt를 제외 → `generate`가 건드리지 않고 `check`도 강제하지 않는다. 손큐레이션본이 깨끗하게 공존.

```go
// pkg/gen/build.go
outs := []Output{ {hugo.toml…}, {robots.txt…}, {jsonld.json…} }
if mode := b.Geo.LlmsTxtMode(); mode != "manual" && mode != "off" {
    outs = append(outs, Output{Path: "static/llms.txt", Data: llms.Render(b, llms.Collect(dir, b))})
}
```
`Check`가 `gen.Build`로 기대 목록을 뽑으므로 **이 한 곳만 고치면 generate·check 양쪽이 동시에 llms.txt에서 손을 뗀다.**

### Tier 1 — `auto`를 실제 쓸 만하게

1. **언어 스코프**: 기본 `base`(= `languages` 첫 항목 1개)만 렌더. llms.txt는 정준 언어 가이드다. 단일 언어일 때 그룹 헤딩에서 `{lang}/` 접두 제거(`## opinion`).
2. **포지셔닝 헤더**: `header`(또는 `site.description`)가 있으면 `>` 블록 다음에 삽입.
3. **핀 엔트리**: `pinned`를 목록 최상단에 먼저 렌더(`group` 지정 시 그 그룹 헤딩 아래). 마스터 인덱스를 선두에 둘 수 있다.
4. **사람 라벨·중요도**: `section_labels`로 기계 라벨 대체, 섹션은 선언 순서 유지.
5. **간결화**: `max_summary`로 설명문 절단.

### Tier 2 — (선택) 큐레이션 free-text 헤더 전체 주입

`header`를 임의 마크다운 블록으로 허용(MIGRATION §3-3 옵션 b). pinned/labels로 부족할 때 통째 손큐레이션.

## 수용 기준

- [ ] `geo.llms_txt: manual` 선언 시, 손작성 `static/llms.txt`가 `abloq generate`·`check`를 **무변경 통과**한다.
- [ ] 기본(`auto`) 출력이 **기본 언어 단일** 스코프다(전 언어 덤프 아님). `languages: all`로 옵트인 가능.
- [ ] `pinned`로 마스터 인덱스를 **선두**에 둘 수 있다.
- [ ] `header`/`site.description`가 출력 헤더에 반영된다.
- [ ] parkjunwoo의 65줄 큐레이션본을 **선언적 설정에서 근사 재현**하거나(이상), 최소한 `manual`로 깨끗이 옵트아웃(필수)할 수 있다.
- [ ] 기존 `geo.llms_txt: auto` 문자열은 그대로 동작(하위호환). 출력 멱등성·`check` 정합 유지.
- [ ] 골든 테스트 갱신(`pkg/gen/llms/render_test.go`, `pkg/gen/testdata/golden/`).

## 변경 파일 (예상)

| 유형 | 경로 | 설명 |
|---|---|---|
| 스키마 | `pkg/blogyaml/geo.go` | `LlmsTxt` 문자열→union(문자열 단축형 \| 객체). `LlmsTxtMode()` 헬퍼. `site.description`(선택) 추가 |
| 기본값 | `pkg/blogyaml/default_blog.go` | mode 기본 `auto`, languages 기본 `base` |
| 검증 | `pkg/blogyaml/validate*` | mode∈{auto,manual,off}, languages 유효성, pinned URL 형식 |
| 생성 게이트 | `pkg/gen/build.go` | mode manual/off 시 llms.txt 제외 |
| 렌더 | `pkg/gen/llms/{render,collect,sort_posts,post_line,header_note}.go` | 언어 스코프, 핀 엔트리, 라벨, 단일언어 그룹, summary 상한 |
| 룰 매핑 | `pkg/gen/rule_for.go` | (필요 시) llms.txt 부재 허용 |
| 문서 | `docs/blog-yaml.md` | 스키마 레퍼런스 갱신 |
| 테스트 | 위 각 `*_test.go` + 골든 | 케이스 추가 |

## 영향 / 우선순위

- **Tier 0(옵트아웃)**: High — parkjunwoo 본번 적용의 **유일한 차단 요소**. 이것만 들어가면 큐레이션 유지로 즉시 적용 가능.
- **Tier 1(auto 개선)**: Medium — abloq를 손큐레이션 없이 쓰려는 일반 블로그의 GEO 품질. 다른 도그푸드/외부 채택에 필요.
- **Tier 2**: Low — 선택적 편의.

## 참고

- 본번 적용 가이드: `dogfood/parkjunwoo/MIGRATION.md` §3-3(llms.txt 분석), §8(승인 대기).
- 인스턴스 SSOT: `dogfood/parkjunwoo/blog.yaml`(`geo.llms_txt: auto` 선언 중).
- 현행 큐레이션본 실물: parkjunwoo `hugo/static/llms.txt`(65줄).
