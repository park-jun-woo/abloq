//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what 글 그룹(섹션/어간)별로 존재하는 언어 목록 수집 — hreflang 상호 참조의 기대 집합
package gate

// siblingLangs maps each article group key to the languages it exists in,
// in discovery (blog.yaml language) order.
func siblingLangs(arts []*Article) map[string][]string {
	sibs := map[string][]string{}
	for _, a := range arts {
		key := a.Section + "/" + a.Slug
		sibs[key] = append(sibs[key], a.Lang)
	}
	return sibs
}
