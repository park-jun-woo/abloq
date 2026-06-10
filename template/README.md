# template/ — abloq blog template (layout pack + deploy IaC)

`abloq init`이 degit식으로 복제하는 템플릿 전체. 페이로드는 `files/` 아래에
있고 `embed.go`가 바이너리에 임베드한다 — abloq 바이너리 하나만 있으면
어디서든 블로그를 스캐폴드할 수 있다.

## 내용

- `files/layouts|assets` — parkjunwoo.com에서 추출한 기본 레이아웃 팩.
  개인 하드코딩(저자명·메뉴·색·도메인) 제거: 정체성은 blog.yaml → 생성된
  hugo.toml params와 `data/jsonld.json`에서 오고, 소셜 링크·홈 배너는
  `layouts/partials/hooks/` 파셜 훅이다. RTL(ar/he/fa/ur) 처리 유지.
- `files/layouts/partials/jsonld.html` — Phase002가 생성하는
  `data/jsonld.json`을 소비하는 JSON-LD 파셜.
- `files/deploy/terraform` — S3+CloudFront(+옵션 WAF·CF 로그 버킷·IndexNow
  키 배치) IaC. 모듈 단위 토글.

## 레이아웃 정책 (설계서 §9 일반화의 늪 방어)

**레이아웃은 교체 가능하나 지원하지 않는다.** 게이트와 생성기는 blog.yaml,
content/, 생성 파일만 만지므로 어떤 Hugo 레이아웃으로 바꿔도 동작한다.
단, 교체 레이아웃은 두 가지 계약을 지켜야 한다: `data/jsonld.json` 소비,
글 페이지의 자기 포함 hreflang alternate 출력(hreflang-complete 게이트가
빌드 HTML을 읽는다).

이 디렉토리는 `.ffignore` 대상 — filefunc 규칙은 Go 코드에만 적용된다.
