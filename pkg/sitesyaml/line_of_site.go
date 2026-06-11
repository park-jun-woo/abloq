//ff:func feature=sitesyaml type=parser control=sequence topic=diagnostics
//ff:what 사이트 항목 i의 키 라인을 조회 — 키가 없으면 항목 자체 라인으로, 그것도 없으면 1로 폴백
package sitesyaml

import "fmt"

// lineOfSite returns the source line of sites[i].key, falling back to the
// list item's own line when the key is absent (required-key diagnostics
// should point at the offending site entry, not at line 1).
func lineOfSite(idx lineIndex, i int, key string) int {
	item := fmt.Sprintf("sites[%d]", i)
	if line, ok := idx[item+"."+key]; ok {
		return line
	}
	return lineOf(idx, item)
}
