//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what 맵 키를 정렬해 반환 — 맵 기반 룰의 진단 순서를 결정적으로 유지
package blogyaml

import "sort"

// sortedKeys returns the map keys in ascending order for deterministic diagnostics.
func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
