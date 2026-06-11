//ff:func feature=quest type=parser control=sequence
//ff:what 루트 기준 글 경로에서 Item Key 부품(lang/section/slug) 파생 — content/<lang>/<section>/<slug>.md 또는 <slug>/index.md
package writing

import (
	"path"
	"strings"
)

// keyParts derives the item key parts from a root-relative article path.
// Flat articles are content/<lang>/<section>/<slug>.md; bundles are
// content/<lang>/<section>/<slug>/index.md. Anything else is not seedable.
func keyParts(article string) (lang, section, slug string, ok bool) {
	parts := strings.Split(path.Clean(article), "/")
	if len(parts) == 5 && parts[0] == "content" && parts[4] == "index.md" {
		return parts[1], parts[2], parts[3], true
	}
	if len(parts) == 4 && parts[0] == "content" && strings.HasSuffix(parts[3], ".md") {
		return parts[1], parts[2], strings.TrimSuffix(parts[3], ".md"), true
	}
	return "", "", "", false
}
