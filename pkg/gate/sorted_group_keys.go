//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what 그룹 맵 키를 정렬해 반환 — slug-consistency 진단 순서를 결정적으로 유지
package gate

import "sort"

// sortedGroupKeys returns the group keys in ascending order.
func sortedGroupKeys(groups map[string][]*Article) []string {
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
