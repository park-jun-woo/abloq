# 기존 Hugo 블로그 → abloq 인스턴스 마이그레이션 가이드

이미 운영 중인 Hugo 사이트를 abloq 인스턴스로 전환하는 절차다. 핵심은 **블로그를
다시 만들지 않는다는 것** — 기존 `content/`·`layouts/`·`assets/`·`static/`은
그대로 두고, `blog.yaml`(SSOT)을 추가한 뒤 파생물(hugo.toml·robots.txt·llms.txt·
jsonld.json)을 abloq가 생성하도록 소유권을 옮긴다. 그 결과 빌드 산출물이 기존과
**바이트 동일**하게 유지되는지를 먼저 증명하고, 그다음 본번에 적용한다.

용어: **인스턴스 루트** = `blog.yaml`이 사는 디렉토리. Hugo 프로젝트 루트와 같은
곳에 두는 것이 보통이다(저장소 루트일 수도, `hugo/` 같은 하위 디렉토리일 수도
있다). 아래에서 `<instance>`로 표기한다.

## 0. 전제

- `abloq` 설치: `go install github.com/park-jun-woo/abloq/cmd/abloq@latest`
- `hugo` 설치(번역 게이트 `hugo-build`와 빌드 검증에 필요). 마이그레이션 검증은
  **기존 사이트를 빌드하던 바로 그 hugo 버전·플래그**로 해야 비교가 의미 있다.
- 작업 트리가 깨끗한 git 저장소.

## 1. 동등성 우선 — 사본에서 먼저 증명한다

본번 저장소를 건드리기 전에 **사본을 떠서** abloq 산출물을 적용하고, 같은 hugo
바이너리·같은 플래그(보통 no-minify)로 빌드해 **기존 설정 빌드와 diff**한다.

```bash
cp -r <repo> /tmp/abloq-migrate && cd /tmp/abloq-migrate
# ↓ §2의 파일 추가·교체를 사본에 적용한 뒤
diff -rq <기존_빌드_public> <abloq_빌드_public>
```

목표는 **HTML/XML 페이지 바이트 동일**과 **`sitemap.xml`(hreflang 포함) 동일**,
**URL 구조 동일**이다. 차이가 나오면 그것이 의도된 차이(§3)인지, 콘텐츠
비결정성(§4)인지 분류한 뒤에만 본번으로 넘어간다.

## 2. 본번에 바꿀 파일 (순서대로)

작업 디렉토리는 저장소 루트, abloq 인스턴스 루트는 `<instance>`.

1. **`<instance>/blog.yaml` 추가** — SSOT 선언. 사이트·언어·섹션·구조·GEO 임계·
   배포를 한 파일에 담는다. 스키마: `docs/blog-yaml.md`. (새 사이트라면
   `abloq init`이 스캐폴드하지만, 기존 사이트는 현 상태를 *그대로 선언*하는
   blog.yaml을 손으로 맞춘다 — §6 참조.)
2. **`<instance>/config/_default/hugo.toml` 추가 (사람 소유 오버레이)** —
   blog.yaml이 다루지 **않는** 인스턴스 전용 설정만 분리한다: 언어별 title/params/
   menus, pagination, taxonomies, markup, params.* 등. **자유 수정 가능.**
   단 blog.yaml이 내는 키는 여기 **재선언 금지**(아래 §3-1 목록).
3. **`<instance>/hugo.toml` 교체** — `abloq generate <instance>`가 생성·소유한다.
   이후 직접 수정 금지(`abloq check`가 드리프트를 검출). 기존 파일은 git 이력에 남는다.
4. **`<instance>/static/robots.txt` 교체** — `abloq generate` 산출물(§3-2).
5. **`<instance>/static/llms.txt` 교체** — `abloq generate` 산출물.
   **⚠ 결정 필요(§3-3)**: 손큐레이션본을 유지하려면 `geo.llms_txt: manual`(또는
   `off`)로 옵트아웃한다.
6. **`<instance>/data/jsonld.json` 추가** — `abloq generate` 산출물. 레이아웃이
   아직 소비하지 않아도 존재만으로 무해(추후 `partials/schema.html` 전환 옵션).
7. **`<instance>/CLAUDE.md` 추가** — `abloq claudemd <instance>` 생성(블로그 운영
   매뉴얼). 저장소 루트에 별도 `CLAUDE.md`가 있다면 게시 절차 섹션을 이 파일
   참조로 교체하고, 루트는 프로필·자격증명·인프라 등 저장소 전반만 남긴다.
8. **빌드 파이프라인(Makefile 등) 수정** —
   - **게이트용 빌드는 minify 없이** `<instance>/public`에 떠야 한다. hreflang
     게이트(`hreflang-complete`)가 인스턴스 루트의 `public/` 원본 HTML을 읽기
     때문. 배포 빌드를 `<instance>/public`으로 일원화하고 배포 sync 경로를 거기에
     맞추거나, 게이트용 빌드를 별도로 한 번 더 뜬다.
   - 기존에 `.md` 병행 서빙 스크립트가 있다면 `abloq postbuild md <instance>`로
     치환(전 언어 .md를 노이즈 제로로 산출 — 직접 만든 스크립트보다 커버리지·
     포맷이 낫다).
9. **`.gitignore`** — `<instance>/public/` 추가(빌드 출력 커밋 금지).
10. **검증 시퀀스 실행** —
    ```bash
    abloq validate <instance>      # blog.yaml 자체 검증
    abloq generate <instance>      # 파생물 생성
    abloq check    <instance>      # 파생물 드리프트 0 확인
    (cd <instance> && hugo -d public)   # minify 없이
    abloq postbuild md <instance>
    abloq gate --offline <instance>     # 전 글 PASS 확인 (네트워크 룰 스킵)
    abloq gate         <instance>       # 네트워크 가능 환경이면 citation-exists 포함
    ```

이미지 변환·OG 생성 스크립트가 따로 있다면 `abloq image convert`/`abloq image og`로
치환 가능하나 강제는 아니다. 기존 IaC(terraform 등)는 이번 범위에서 검증하지 않으니
**유지가 기본값** — 배포 모듈 정합 전환은 별도 승인 사항으로 분리한다.

## 3. 산출물 diff 분석 (의도된 차이)

### 3-1. hugo.toml — abloq가 소유하는 키 (오버레이에 재선언 금지)

abloq generate는 blog.yaml 파생 **코어만** 담는다:

- `baseURL`, `title`
- `defaultContentLanguage`, `defaultContentLanguageInSubdir`
- `[sitemap]`
- `[languages.*]`의 `languageCode`, `contentDir`, `weight`

`weight`는 blog.yaml의 언어 선언 순서를 따른다(첫 언어 = 기본). 그 외 모든
설정(언어별 title/params/menus, pagination 등)은 §2-2 오버레이로 옮긴다 —
Hugo의 config 병합으로 결과가 동등해진다(사본 빌드 diff로 증명).

> **흔한 개선 효과:** 기존 설정에 `languageCode`가 없었다면 abloq가 이를 추가하면서
> RSS 메타·언어 인지 타이틀 처리가 개선된다. 그 결과 slug 역산으로 만들던 일부
> 태그/택소노미 페이지 타이틀이 **태그 원문 보존**으로 바뀔 수 있다(의도된 개선).

### 3-2. robots.txt

- 정책 의미는 보존된다: AI 크롤러 allow/block 정책 + `User-agent: *` + sitemap.
- 표기가 정규화된다: `Allow: /` → 빈 `Disallow:`(동일 의미), 분류
  (training/search/fetch)별 사전순 그룹. robots UA 매칭은 대소문자 무관이라
  표기 차이는 효과 동일.
- 일반 웹 검색 크롤러(bingbot 등)는 abloq의 AI봇 사전 소관이 아니다 —
  `User-agent: * Allow: /`가 커버한다. 차단 정책으로 전환할 때만 재검토.
- blog.yaml `geo.crawlers`에 봇/분류 → `allow|block`을 선언하면 그대로 반영된다.

### 3-3. llms.txt — ⚠ 결정 필요

- abloq `auto`는 전 언어 × 섹션 자동 인덱스를 만들고 글별 front matter `summary`를
  설명으로 쓴다.
- 손으로 큐레이션한 기존 llms.txt(주제 그룹·수동 요약·특정 문서 핀)가 있다면
  자동본은 그 큐레이션을 잃는다. 선택지(`geo.llms_txt`, 상세 `docs/blog-yaml.md`):
  - **`auto`** — SSOT 일관성, 갱신 자동. 큐레이션 헤더/핀/섹션 라벨을 객체 폼으로
    부분 주입 가능(`header`/`pinned`/`section_labels`/`max_summary`).
  - **`manual`** — 손큐레이션본을 그대로 두고 `generate`/`check`가 llms.txt를
    불간섭. 기존 인덱스를 100% 보존하고 싶을 때.
  - **`off`** — llms.txt를 내지 않는다.
- 권고: 큐레이션이 가벼우면 `auto`(+객체 폼 헤더), 무거우면 `manual`로 보존 후
  여유 있을 때 `auto`로 이관.

### 3-4. .md 병행 서빙

`abloq postbuild md <instance>`는 전 글(기본 언어는 루트, 그 외 `/{lang}/...`)에
노이즈 제로 `.md`를 병행 산출한다 — AI 컨텍스트 포맷. 직접 만든 스크립트가 기본
언어만, 또는 제목 뒤 빈 줄 같은 노이즈를 넣었다면 이쪽이 더 깨끗하다.

## 4. 잔여 빌드 diff 다루기 (abloq 무관)

바이트 diff가 남는데 abloq 산출물 교체로 설명되지 않는다면 **콘텐츠 비결정성**을
의심한다. 대표 사례:

- **같은 태그의 표기 흔들림.** 한 택소노미를 두 가지 표기(공백·구분점·전각 차이
  등)로 쓴 글이 섞이면, 같은 설정으로 두 번 빌드해도 태그 페이지 타이틀이 바뀐다
  (hugo 자체 비결정성, abloq와 무관). 기존 설정 self-diff(같은 입력 2회 빌드)로
  확인되면 **콘텐츠 표기 통일**로 해소한다.

기존 설정을 두 번 빌드해 self-diff를 떠 보면, 어떤 차이가 abloq가 아니라 입력
자체에서 오는지 분리할 수 있다.

## 5. 게이트 회귀 — 기존 코퍼스에서 유의할 점

- **레거시 전수 감사가 아니다.** 네트워크/근거 룰(`citation-exists`,
  `numeric-claim-sourced`)은 **git HEAD 대비 신규 주장만** 판정한다. 즉 기존
  코퍼스의 오래된 무출처 주장이 마이그레이션 자체를 FAIL시키지 않는다. 레거시
  전수 점검은 스캐너(`abloq scan evidence` 등) 소관이다.
- **특수 페이지 정책.** front matter에 `layout` 키를 가진 페이지(전용 레이아웃
  소유)는 글 모양 룰(image-first·image-attribution·section-order·
  heading-canonical·front-matter-schema·min-sources·numeric-claim-sourced)을
  스킵하고, 무결성 룰(slug-consistency·기준선 비교·hreflang-complete·
  citation-exists)만 적용한다. about/소개 같은 페이지가 글 룰에 걸리면 `layout`
  키로 분류한다.
- **`geo.min_sources`는 현실을 선언한다.** 기존 사이트가 글마다 출처 섹션을
  강제하지 않았다면 `min_sources: 0`으로 *현 상태를 그대로* 선언한다(전수 FAIL
  방지). 임계 상향은 사람의 결정 — 올리는 순간부터 신규/수정 글에 적용된다.
- **치즈 방어는 마이그레이션 후에도 유효.** lastmod 단독 갱신 → `honest-lastmod`
  FAIL, 신규 무출처 수치 주장 → FAIL. 사본에서 한 번 확인해 두면 안심.

## 6. blog.yaml 선언 메모

- **언어 선언 순서 = weight 순서**(첫 항목 = 기본 언어). 기존 hugo.toml의 weight
  순서를 그대로 옮긴다.
- **`structure.headings`는 언어별 정준 표기 1개만.** 코퍼스에 한 섹션을 여러 표기
  (띄어쓰기·동의어 변형 등)로 쓴 글이 있다면, **실사용 빈도가 가장 높은 표기**를
  정준으로 선언한다. 변형 표기는 게이트가 "인식하지 않는 본문 헤딩"으로 취급 —
  FAIL을 만들지 않지만 순서·보존 검사 대상도 아니다. 변형 동의어 지원은 스키마
  확장 후보다.
- blog.yaml은 strict 파싱이다 — 스키마에 없는 키는 `unknown-key` 에러. 인스턴스
  전용 설정은 blog.yaml이 아니라 §2-2 hugo.toml 오버레이로 보낸다.

## 7. 적용과 롤백

마이그레이션은 **단일 커밋**으로 적용한다. 롤백 = `git revert <커밋>` 1회.

- 교체 파일(hugo.toml·robots.txt·llms.txt)의 원본은 직전 커밋에 그대로 있다.
- 추가 파일(blog.yaml·config/_default/·data/jsonld.json·CLAUDE.md)은 revert로
  제거된다.
- 배포 산출물은 배포 타깃 재실행으로 직전 상태 복원(예: S3 sync --delete).
- **abloq 의존은 빌드 파이프라인에만 있다.** 생성 파일이 전부 정적이므로 사이트
  자체는 hugo 단독으로도 기존과 동일하게 빌드된다 — abloq가 빠져도 사이트는 산다.

## 8. 본번 적용 전 체크리스트

1. §1 사본 빌드가 기존과 바이트 동일(의도된 §3 차이 제외).
2. §3-3 llms.txt 방침 결정(`auto`/`manual`/`off`).
3. §4 잔여 diff가 전부 콘텐츠 비결정성으로 설명되고, 필요한 표기 통일을 커밋.
4. §2-8 게이트용 빌드 경로(non-minify `public/`) 확정.
5. `abloq gate --offline <instance>` 전 글 PASS.
6. IaC/배포 모듈 정합 전환은 이번 범위에서 제외(기존 유지가 기본값).
