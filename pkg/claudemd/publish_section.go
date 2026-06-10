//ff:func feature=claudemd type=generator control=sequence
//ff:what CLAUDE.md 게시 절차 섹션 — 집필 퀘스트 호출 순서(작성→게이트→번역→generate→빌드→postbuild→배포)
package claudemd

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// publishSection renders the publish procedure: prose work is the agent's,
// every judgement is a deterministic abloq command.
func publishSection(b *blogyaml.Blog) string {
	def, rest := "", []string(nil)
	if len(b.Languages) > 0 {
		def, rest = b.Languages[0], b.Languages[1:]
	}
	var sb strings.Builder
	sb.WriteString("## 게시 절차 (집필 퀘스트)\n\n")
	sb.WriteString("산문만 에이전트가 쓴다. 판정은 전부 게이트다 — 게이트 실패는 작업 미완료다.\n\n")
	fmt.Fprintf(&sb, "1. 원고 작성: `content/%s/{section}/{slug}.md` — 구조 계약 준수, 모든 인용은 실재 URL.\n", def)
	sb.WriteString("   slug는 영문 소문자-하이픈, 불필요한 관사(a/an/the) 제거 — 전 언어가 같은 파일명으로 번역 매칭된다\n")
	sb.WriteString("2. `abloq gate --offline .` — 구조·근거 게이트 통과까지 수정\n")
	fmt.Fprintf(&sb, "3. 번역: 나머지 언어(%s) 전부에 **동일 slug**로 작성 후 다시 `abloq gate --offline .`\n", strings.Join(rest, ", "))
	sb.WriteString("   기술 용어는 원문(영문) 유지, 문화적 예시(인물·인사말·역사 사례)는 현지화한다\n")
	sb.WriteString("4. `abloq generate .` — 글 추가·삭제 시 llms.txt 등 파생물 재생성 (필수)\n")
	sb.WriteString("5. `hugo` — 빌드 0 에러 (게이트 전에는 --minify 금지: hreflang 검사가 원본 HTML을 읽는다)\n")
	sb.WriteString("6. `abloq gate .` — 빌드 후 전체 룰 (hreflang-complete, citation-exists 포함)\n")
	sb.WriteString("7. `abloq postbuild md .` — 글마다 노이즈 제로 .md 병행 산출 (AI 컨텍스트)\n")
	sb.WriteString("8. 배포는 사람이 트리거한다 — `deploy/terraform/README.md`의 S3 sync + invalidation 참조\n\n")
	sb.WriteString("이미지: `abloq image convert <원본> --slug {slug}` → `static/images/{slug}.webp`, 본문 첫 줄 `![..](/images/{slug}.webp)`.\n")
	sb.WriteString("OG 카드가 필요하면 `abloq image og {slug} \"제목\"` (한글 제목은 `--font <TTF>` 필요).\n")
	sb.WriteString("OG 이미지 우선순위: front matter `image:` > 본문 첫 `![](...)` > `/og-image.webp` (레이아웃 규약).\n\n")
	return sb.String()
}
