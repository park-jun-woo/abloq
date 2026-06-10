//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what 대상 글을 섹션/파일 어간 키로 그룹화 — 같은 글의 언어 버전 묶음
package gate

// groupArticles groups the language versions of each article by
// "<section>/<file stem>".
func groupArticles(arts []*Article) map[string][]*Article {
	groups := map[string][]*Article{}
	for _, a := range arts {
		key := a.Section + "/" + a.Slug
		groups[key] = append(groups[key], a)
	}
	return groups
}
