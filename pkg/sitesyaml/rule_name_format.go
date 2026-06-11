//ff:func feature=sitesyaml type=rule control=iteration dimension=1
//ff:what [name-format] name 필수 + URL-safe 슬러그(소문자·숫자·하이픈, 하이픈 시작/끝 금지) 검증 — {site} path param으로 쓰인다
package sitesyaml

import (
	"fmt"
	"regexp"
)

var siteNameRe = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

// ruleNameFormat requires every site name and keeps it a URL-safe slug —
// the name is the {site} path parameter of the backend API.
func ruleNameFormat(filename string, s *Sites, idx lineIndex) []Diagnostic {
	var diags []Diagnostic
	for i, site := range s.Sites {
		if site.Name == "" {
			diags = append(diags, Diagnostic{
				File: filename, Line: lineOfSite(idx, i, "name"), Rule: "name-format",
				Message: fmt.Sprintf("sites[%d].name is required", i),
			})
			continue
		}
		if !siteNameRe.MatchString(site.Name) {
			diags = append(diags, Diagnostic{
				File: filename, Line: lineOfSite(idx, i, "name"), Rule: "name-format",
				Message: fmt.Sprintf("sites[%d].name %q must be a URL-safe slug (lowercase letters, digits, single hyphens)", i, site.Name),
			})
		}
	}
	return diags
}
