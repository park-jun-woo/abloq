//ff:func feature=gate type=rule control=iteration dimension=1 topic=baseline
//ff:what 두 front matter를 lastmod 라인 제외 후 라인 단위 비교 — 첫 차이를 반환
package gate

import "fmt"

// fmIntactDiff reports whether two front matters are identical up to an
// allowed change of the lastmod line. Returns the first difference when not.
func fmIntactDiff(orig, neu string) (string, bool) {
	o, n := fmLinesClean(orig), fmLinesClean(neu)
	if len(o) != len(n) {
		return fmt.Sprintf("line count %d -> %d", len(o), len(n)), false
	}
	for i := range o {
		if o[i] != n[i] {
			return o[i], false
		}
	}
	return "", true
}
