//ff:func feature=quest type=parser control=sequence
//ff:what 대상 언어의 번역 글 경로 파생 — 원문과 같은 모양(플랫 .md 또는 번들 index.md)으로 언어 세그먼트만 치환
package translation

import "strings"

// transPath mirrors the origin's path shape into the target language:
// content/<origin>/<section>/<slug>.md -> content/<lang>/<section>/<slug>.md,
// bundles keep their <slug>/index.md form.
func transPath(src seedSrc, lang string) string {
	if strings.HasSuffix(src.origin, "/index.md") {
		return "content/" + lang + "/" + src.section + "/" + src.slug + "/index.md"
	}
	return "content/" + lang + "/" + src.section + "/" + src.slug + ".md"
}
