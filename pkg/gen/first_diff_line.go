//ff:func feature=gen type=rule control=iteration dimension=1 topic=drift
//ff:what 기대/실제 바이트를 줄 단위 비교해 첫 불일치 라인 번호와 양쪽 줄 내용을 반환
package gen

import "strings"

// firstDiffLine returns the 1-based line number of the first differing line,
// with the expected and actual line text. Equal inputs return (0, "", "").
func firstDiffLine(want, got []byte) (int, string, string) {
	w := strings.Split(string(want), "\n")
	g := strings.Split(string(got), "\n")
	for i := 0; i < len(w) || i < len(g); i++ {
		var wl, gl string
		if i < len(w) {
			wl = w[i]
		}
		if i < len(g) {
			gl = g[i]
		}
		if wl != gl {
			return i + 1, wl, gl
		}
	}
	return 0, "", ""
}
