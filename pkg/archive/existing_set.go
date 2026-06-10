//ff:func feature=archive type=client control=iteration dimension=1
//ff:what 기존 영수증 (kind, target) 페어를 멱등 필터용 집합으로 변환
package archive

// existingSet keys every already-receipted pair as kind + "\n" + target.
func existingSet(existing []Existing) map[string]bool {
	seen := make(map[string]bool, len(existing))
	for _, e := range existing {
		seen[e.Kind+"\n"+e.Target] = true
	}
	return seen
}
