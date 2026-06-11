//ff:func feature=quest type=parser control=sequence
//ff:what insight.yaml 경로의 규약상 대상 글 경로(insight.PathFor의 역산) — 번들 insight.yaml→index.md, *.insight.yaml→*.md
package writing

import (
	"path"
	"strings"
)

// articleFor inverts the Phase015 sidecar convention (insight.PathFor): a
// bundle's insight.yaml maps to its index.md, a <stem>.insight.yaml sidecar
// maps to <stem>.md. Paths are slash-separated, instance-root relative.
func articleFor(insightPath string) (string, bool) {
	dir, base := path.Split(insightPath)
	if base == "insight.yaml" {
		return dir + "index.md", true
	}
	if strings.HasSuffix(base, ".insight.yaml") {
		return dir + strings.TrimSuffix(base, ".insight.yaml") + ".md", true
	}
	return "", false
}
