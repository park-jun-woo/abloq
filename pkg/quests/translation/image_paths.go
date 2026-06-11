//ff:func feature=quest type=parser control=iteration dimension=1 topic=lossless
//ff:what 본문 prose의 마크다운 이미지 경로 전수 추출 — `![..](path)`의 path, 코드 펜스 안 제외 (패리티 ③)
package translation

import (
	"regexp"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

var reImagePath = regexp.MustCompile(`!\[[^\]]*\]\(([^)\s]+)[^)]*\)`)

// imagePaths extracts every markdown image destination in the body prose.
// Image paths are translation-invariant; the parity rule compares them as a
// bidirectional multiset.
func imagePaths(d *agate.Doc) []string {
	var paths []string
	for _, raw := range proseLines(d) {
		for _, m := range reImagePath.FindAllStringSubmatch(raw, -1) {
			paths = append(paths, m[1])
		}
	}
	return paths
}
