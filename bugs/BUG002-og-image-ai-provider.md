# BUG002 — OG 이미지 생성을 provider 선택형으로 (Gemini 등 AI 이미지 API 옵션)

> 상태: OPEN · 유형: Enhancement · 심각도: Medium · 발견: parkjunwoo.com OG 처리 검토(BUG001 후속)

## 한 줄 요약

`abloq image og`는 현재 **순수 Go 결정론 텍스트 카드**(`RenderOG`, 1200×630)만 만든다. 여기에 **Gemini 등 외부 AI 이미지 API**를 *선택 가능한 provider*로 추가하되, AI 생성이 비결정·네트워크라는 점을 abloq의 결정론 모델과 충돌 없이 다루는 절차를 정의한다.

## 현황

- `abloq image og <slug> <title> [--brand --font --out]` → `cmd/abloq/run_image_og.go` → `pkg/img/og.go OG()` → `pkg/img/render_og.go RenderOG()`.
- `RenderOG`: 1200×630, 흰 배경, 제목 중앙 줄바꿈, 하단 액센트 브랜드 라인. 임베드 Go Bold(라틴), CJK/RTL은 `--font`로 TTF.
- `OGSpec{Title, Brand, FontPath, Out}` (`pkg/img/og_spec.go`) — 텍스트 카드 전용.
- `pkg/img`는 **순수 Go·외부 의존 0**(ff 주석 명시). 변환 파이프라인 재사용 가능: `DecodeAny`/`ResizeMax(src,maxW)`/`FlattenWhite`/`SaveWebP`/`Convert(src,dst,maxW)`.

→ 결과물이 텍스트 카드 한 종류뿐이다. 일러스트/배경 아트형 OG가 필요하면 외부 도구로 만들어 `--font` 우회로도 못 메운다.

## 동기

OG 카드는 SNS·검색 썸네일의 첫인상이다. 텍스트 카드는 일관·저비용이지만, 주제별 **배경 아트**(추상 기술 비주얼 등)를 원하는 경우가 있다. 이미 사람이 Gemini 등으로 이미지를 만들어 손으로 `convert`하는 흐름이 있다면(예: parkjunwoo `files/images/Gemini_Generated_*`), 그 절차를 abloq 안으로 들여 **provider만 바꿔 동일 명령**으로 생성·후처리·참조 안내까지 일관되게 한다.

## 설계 제약 — 결정론 텐션 (핵심)

abloq의 중추 명제는 "**생성은 확률적, 검증은 결정론적**"이다. 현 `RenderOG`는 결정론적(같은 입력 → 같은 바이트, 멱등, `generate`/`check` 파생물과 동급으로 안전)이다. 반면 **AI 이미지 생성은 비결정(같은 프롬프트 → 매번 다른 이미지) + 네트워크 + 유료**다.

따라서 AI OG는 **파생물(derived)이 될 수 없다.** 절차상 위치는 인용 샘플링·퀘스트 산문과 같은 **1회성 부수효과 자산 생성**이다:

- AI provider 출력은 **명시 호출 시 1회 생성 → 사람/에이전트 검토 → 커밋**. 그 시점 이후로는 평범한 정적 `image:` 자산으로 취급한다.
- **절대 `generate`/`check` 루프에 넣지 않는다**(바이트 멱등 파생 가정을 깬다). `local`(결정론)만 빌드 루프에 들어갈 자격이 있다 — 그러나 OG는 애초에 빌드 파생물이 아니라 자산이므로 둘 다 명시 호출이다.
- 자격증명은 **env로만**(abloq 관례: archive/ingest/sample과 동일) — `GEMINI_API_KEY`/`GOOGLE_API_KEY`. blog.yaml·인자·`OGSpec`에 키를 넣지 않는다.
- 순수 Go 경계 보존: AI provider(HTTP)는 `pkg/img`에서 분리된 하위 패키지(`pkg/img/ogprovider`)에 둔다. `local`은 의존 0 유지. func는 `net/http` 직접 import 금지·pkg 위임(evidence/linkcheck 선례).

## 개선안

### CLI

```bash
abloq image og <slug> <title> \
  --provider gemini \        # local(기본) | gemini | ...
  --model <id> \             # provider별 모델(선택, 기본은 provider 기본값)
  --overlay \                # AI 배경 위에 제목/브랜드를 결정론적으로 합성(선택)
  [--brand ... --font ...]   # overlay·local 공통 텍스트 옵션
```

- provider 해석 순서: `--provider` 플래그 > blog.yaml `image.og.provider` > `local`.
- AI 경로 흐름: provider가 이미지 생성 → **기존 파이프라인으로 후처리**(임의 크기/정방형 → 1200×630 센터 크롭+리사이즈, `FlattenWhite`, `SaveWebP`) → `static/images/{slug}.webp` → 참조 안내 출력(현행 `runImageOG`와 동일 UX) + **사용 모델·예상 비용 echo**.
- 실패(키 없음·API 오류·쿼터)는 명확한 진단 + exit 1.

### Hybrid 옵션 (`--overlay`)

AI는 **배경 아트**만 생성하고, 제목/브랜드 라인은 `RenderOG`의 텍스트 합성을 그 배경 위에 결정론적으로 올린다. 가독성 + 브랜드 일관성을 지키면서 비주얼만 AI로. 이를 위해 `render_og.go`에서 **텍스트 오버레이를 임의 배경에 합성 가능한 함수로 추출**한다(현재는 흰 캔버스에 직접 그림).

### blog.yaml

```yaml
image:
  og:
    provider: local          # local | gemini | ...   (기본 local)
    model: ""                # provider 기본 모델 사용 시 빈 값
    overlay: true            # 기본 합성 여부
    prompt: |                # 프롬프트 템플릿 ({title}/{summary}/{brand} 치환)
      Minimal abstract tech background for "{title}". Brand-color accent.
      No text, no words, 1200x630 composition with safe central margin.
```

- **키·시크릿은 blog.yaml에 절대 두지 않는다**(env 전용).
- `image.og` 미선언 시 전부 기존 `local` 동작(하위호환).

### Provider 추상화

```go
// pkg/img/ogprovider/provider.go
type Provider interface {
    // Generate returns a raw image (any size) for the prompt; caller post-processes.
    Generate(ctx context.Context, prompt string) (image.Image, error)
}
// local.go(RenderOG 래핑) · gemini.go(HTTP, env 키) · 테스트용 stub(고정 픽스처 이미지)
```

테스트는 **stub provider**(캔버스 픽스처 반환)로 후처리·오버레이·경로 안내를 검증 — **실 네트워크 금지**(archive-stub/GSC-fixture 선례).

## 수용 기준

- [ ] 기본 `--provider local`(또는 미지정)은 현행과 **바이트 동일** — 결정론·의존 0 유지.
- [ ] `--provider gemini`가 env 키로 API 호출 → 1200×630 WebP를 `static/images/{slug}.webp`에 생성, 참조 안내 + 모델/비용 출력.
- [ ] `--overlay`가 AI 배경 위에 제목/브랜드를 결정론적으로 합성한다.
- [ ] AI 경로는 `generate`/`check` 어디에도 들어가지 않는다(비결정 자산은 파생물 금지).
- [ ] 키는 env로만; blog.yaml/인자/OGSpec에 시크릿 없음. 키 부재 시 명확한 exit 1.
- [ ] Provider 인터페이스가 stub 가능 — 전 테스트가 실 네트워크 없이 통과.
- [ ] `pkg/img` 순수 Go 로컬 경로의 외부 의존 0 유지(AI는 `pkg/img/ogprovider`로 격리, func의 net/http 직접 import 금지).
- [ ] `image.og` 미선언 시 완전한 하위호환.

## 변경 파일 (예상)

| 유형 | 경로 | 설명 |
|---|---|---|
| 신규 pkg | `pkg/img/ogprovider/{provider.go, local.go, gemini.go}` | Provider 인터페이스 + 로컬 래퍼 + Gemini(HTTP·env 키) |
| 후처리 | `pkg/img/{resize_max,flatten_white,save_webp}.go` 재사용 + (필요 시) `crop_center.go` | 임의 크기 → 1200×630 센터 크롭 |
| 오버레이 | `pkg/img/render_og.go` | 텍스트 합성을 임의 배경에 적용 가능하도록 추출 |
| 스펙 | `pkg/img/og_spec.go` | Provider/Model/Overlay/Prompt 필드(또는 별도 OGRequest) |
| 본체 | `pkg/img/og.go` | provider 분기: local→RenderOG, AI→Generate→후처리(+overlay) |
| CLI | `cmd/abloq/new_image_og_cmd.go`, `run_image_og.go` | `--provider/--model/--overlay` 플래그, provider 해석·env 키·모델/비용 echo |
| 스키마 | `pkg/blogyaml/` (image 블록) + `validate*` | `image.og`(provider/model/overlay/prompt), 시크릿 금지 검증 |
| 문서 | `docs/blog-yaml.md`, `MANUAL.md`(§3 이미지) | provider 옵션·env 키·절차 |
| 테스트 | 위 각 `*_test.go` + stub provider | 후처리·오버레이·키 부재·하위호환 |

## 영향 / 우선순위

- **Medium(Enhancement).** 기본 `local`은 무변동이라 회귀 위험 낮음. AI provider는 순수 opt-in.
- 확장 여지: `openai`(DALL·E/gpt-image), 로컬 SD 등 provider 추가는 같은 인터페이스로. provider별 비용·쿼터·세이프티 정책은 각 구현 책임.
- BUG001과 같은 결의 사안 — abloq의 생성 표면에 **결정론 코어 + opt-in 비결정 provider**를 깨끗이 분리해 얹는 패턴.

## 참고

- 현 OG 처리(템플릿 head.html 우선순위 + image og/convert): BUG001 §"abloq의 OG 처리" 논의 및 `dogfood/parkjunwoo/MIGRATION.md`.
- 자격증명 env 전용 관례: `MANUAL.md` §3(archive/ingest/sample).
- 후처리 파이프라인 선례: `pkg/img/convert.go`(`Convert(src,dst,maxW)`), `ResizeMax`, `FlattenWhite`, `SaveWebP`.
