//ff:func feature=sitesyaml type=rule control=iteration dimension=1 topic=gsc
//ff:what [gsc-site-url] gsc.site_url이 있으면 http(s) URL 또는 sc-domain: 속성 형식인지 검증 — 빈 값은 GSC 미사용으로 적법
package sitesyaml

import "fmt"

// ruleGSCSiteURL validates the optional Search Console property identifier.
// An empty site_url is legal — the site simply does not poll GSC.
func ruleGSCSiteURL(filename string, s *Sites, idx lineIndex) []Diagnostic {
	var diags []Diagnostic
	for i, site := range s.Sites {
		msg := gscSiteURLProblem(site.GSC.SiteURL)
		if msg == "" {
			continue
		}
		diags = append(diags, Diagnostic{
			File: filename, Line: lineOfSite(idx, i, "gsc.site_url"), Rule: "gsc-site-url",
			Message: fmt.Sprintf("sites[%d].gsc.site_url %s", i, msg),
		})
	}
	return diags
}
