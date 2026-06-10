//ff:func feature=scan type=client control=sequence topic=evidence
//ff:what 점검 URL의 scheme+host를 HostOverride로 재작성 — 경로·쿼리는 유지, 미설정·파싱 불가면 원본 그대로
package evidence

import "net/url"

// overrideURL rewrites the probe target onto the stub host while keeping the
// original path and query, so the stub can branch on the path (200/404).
// Without an override (production) the URL passes through untouched.
func (c *Checker) overrideURL(raw string) string {
	if c.HostOverride == "" {
		return raw
	}
	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	o, err := url.Parse(c.HostOverride)
	if err != nil {
		return raw
	}
	u.Scheme = o.Scheme
	u.Host = o.Host
	return u.String()
}
