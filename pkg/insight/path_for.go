//ff:func feature=insight type=parser control=sequence
//ff:what 글 경로의 규약상 insight.yaml 경로 — 번들(index.md)은 디렉토리의 insight.yaml, 플랫은 {어간}.insight.yaml 사이드카
package insight

import (
	"path/filepath"
	"strings"
)

// PathFor returns the conventional insight.yaml path for an article: bundles
// (<dir>/index.md) keep insight.yaml next to index.md, flat articles use a
// <stem>.insight.yaml sidecar. Pairing is by file system name — front matter
// slug overrides are ignored (deterministic, no parsing).
func PathFor(articlePath string) string {
	dir, base := filepath.Split(articlePath)
	if base == "index.md" {
		return filepath.Join(dir, "insight.yaml")
	}
	stem := strings.TrimSuffix(base, filepath.Ext(base))
	return filepath.Join(dir, stem+".insight.yaml")
}
