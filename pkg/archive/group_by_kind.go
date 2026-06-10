//ff:func feature=archive type=client control=iteration dimension=1
//ff:what 선택된 pending 영수증을 kind별 그룹으로 분류 (입력 순서 보존)
package archive

// groupByKind splits the selected pending receipts per kind, preserving
// the input (receipts.id) order inside each group.
func groupByKind(pending []Pending) map[string][]Pending {
	groups := make(map[string][]Pending, len(Kinds))
	for _, p := range pending {
		groups[p.Kind] = append(groups[p.Kind], p)
	}
	return groups
}
