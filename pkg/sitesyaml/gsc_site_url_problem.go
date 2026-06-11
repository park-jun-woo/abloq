//ff:func feature=sitesyaml type=rule control=sequence topic=gsc
//ff:what GSC 속성 식별자 1건의 형식 문제를 메시지로 반환 — 빈 값·sc-domain:도메인·절대 http(s) URL이면 "" (적법)
package sitesyaml

import (
	"fmt"
	"net/url"
	"strings"
)

// gscSiteURLProblem returns "" when the property identifier is legal: empty
// (GSC unused), "sc-domain:<domain>" (domain property) or an absolute
// http(s) URL (URL-prefix property). Otherwise it describes the problem.
func gscSiteURLProblem(raw string) string {
	if raw == "" {
		return ""
	}
	if rest, ok := strings.CutPrefix(raw, "sc-domain:"); ok {
		if rest == "" {
			return fmt.Sprintf("%q must name a domain after the sc-domain: prefix", raw)
		}
		return ""
	}
	u, err := url.Parse(raw)
	if err != nil {
		return fmt.Sprintf("%q is not a valid URL: %v", raw, err)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Sprintf("%q must be an http(s) URL or use the sc-domain: prefix", raw)
	}
	if u.Host == "" {
		return fmt.Sprintf("%q must have a host", raw)
	}
	return ""
}
