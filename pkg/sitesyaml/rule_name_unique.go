//ff:func feature=sitesyaml type=rule control=iteration dimension=1
//ff:what [name-unique] 사이트 name 중복 금지 검증 — name은 sites 테이블 유니크 키이자 {site} path param이다. 빈 name은 name-format 소관이라 스킵
package sitesyaml

import "fmt"

// ruleNameUnique rejects duplicate site names: the name is the unique upsert
// key of the sites table. Empty names are skipped — name-format reports them.
func ruleNameUnique(filename string, s *Sites, idx lineIndex) []Diagnostic {
	seen := map[string]bool{}
	var diags []Diagnostic
	for i, site := range s.Sites {
		if site.Name == "" {
			continue
		}
		if seen[site.Name] {
			diags = append(diags, Diagnostic{
				File: filename, Line: lineOfSite(idx, i, "name"), Rule: "name-unique",
				Message: fmt.Sprintf("sites[%d].name %q is declared more than once", i, site.Name),
			})
			continue
		}
		seen[site.Name] = true
	}
	return diags
}
