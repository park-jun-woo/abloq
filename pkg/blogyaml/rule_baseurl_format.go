//ff:func feature=blogyaml type=rule control=sequence
//ff:what [baseurl-format] site.baseURL이 절대 http(s) URL이고 query/fragment가 없는지 검증
package blogyaml

import (
	"fmt"
	"net/url"
)

// ruleBaseURLFormat validates site.baseURL shape.
func ruleBaseURLFormat(filename string, b *Blog, idx lineIndex) []Diagnostic {
	diag := func(msg string) []Diagnostic {
		return []Diagnostic{{File: filename, Line: lineOf(idx, "site.baseURL"), Rule: "baseurl-format", Message: msg}}
	}
	raw := b.Site.BaseURL
	if raw == "" {
		return diag("site.baseURL is required")
	}
	u, err := url.Parse(raw)
	if err != nil {
		return diag(fmt.Sprintf("site.baseURL %q is not a valid URL: %v", raw, err))
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return diag(fmt.Sprintf("site.baseURL %q must use http or https scheme", raw))
	}
	if u.Host == "" {
		return diag(fmt.Sprintf("site.baseURL %q must have a host", raw))
	}
	if u.RawQuery != "" || u.Fragment != "" {
		return diag(fmt.Sprintf("site.baseURL %q must not have a query or fragment", raw))
	}
	return nil
}
