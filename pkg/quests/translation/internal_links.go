//ff:func feature=quest type=parser control=iteration dimension=1 topic=lossless
//ff:what 링크 목적지에서 내부 글 링크만 선별 — 루트 절대(/) + 확장자 없는 경로, 이미지·정적 자산(확장자 보유) 제외 (패리티 ⑥)
//ff:why 정적 자산 경로는 언어 불변(③과 동일 취급)이지만 글 URL은 언어 프리픽스가 필수 — 확장자 유무로 갈라 자산은 ⑥의 프리픽스 강제에서 제외한다 (Phase017 계획 ⑥)
package translation

import (
	"path"
	"strings"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// internalLinks returns the root-absolute article link destinations in the
// body prose: a leading "/" and no file extension (Hugo article URLs are
// extensionless directories; assets like /images/x.png carry an extension and
// are language-invariant, so they are excluded here). Fragments and query
// strings are stripped before the extension test.
func internalLinks(d *agate.Doc) []string {
	var links []string
	for _, dest := range linkDests(d) {
		if !strings.HasPrefix(dest, "/") {
			continue
		}
		p := strings.SplitN(strings.SplitN(dest, "#", 2)[0], "?", 2)[0]
		if path.Ext(p) != "" {
			continue
		}
		links = append(links, dest)
	}
	return links
}
