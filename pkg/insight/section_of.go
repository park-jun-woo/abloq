//ff:func feature=insight type=parser control=iteration dimension=1
//ff:what 글 경로에서 실위치 섹션을 도출 — 마지막 content/{lang}/{section}/... 세그먼트 기준, 없으면 빈 문자열
package insight

import (
	"path/filepath"
	"strings"
)

// sectionOf derives the section from the article's file system path
// (content/<lang>/<section>/...). Front matter is never consulted —
// pairing and location follow file system names only.
func sectionOf(articlePath string) string {
	segs := strings.Split(filepath.ToSlash(articlePath), "/")
	section := ""
	for i, seg := range segs {
		if seg == "content" && i+3 < len(segs) {
			section = segs[i+2]
		}
	}
	return section
}
