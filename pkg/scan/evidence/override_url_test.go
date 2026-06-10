//ff:func feature=scan type=client control=sequence topic=evidence
//ff:what overrideURL 케이스 — scheme+host만 교체·경로 유지, 미설정·파싱 불가 URL은 원본 그대로
package evidence

import "testing"

func TestOverrideURL(t *testing.T) {
	c := &Checker{HostOverride: "http://127.0.0.1:8099"}
	got := c.overrideURL("https://example.org/source-1?v=2")
	if got != "http://127.0.0.1:8099/source-1?v=2" {
		t.Errorf("override = %q", got)
	}
	if got := c.overrideURL("://bad"); got != "://bad" {
		t.Errorf("unparseable URL must pass through: %q", got)
	}
	none := &Checker{}
	if got := none.overrideURL("https://example.org/x"); got != "https://example.org/x" {
		t.Errorf("no override must pass through: %q", got)
	}
	bad := &Checker{HostOverride: "://bad"}
	if got := bad.overrideURL("https://example.org/x"); got != "https://example.org/x" {
		t.Errorf("unparseable override must pass through: %q", got)
	}
}
