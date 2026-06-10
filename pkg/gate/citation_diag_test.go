//ff:func feature=gate type=rule control=sequence topic=evidence
//ff:what citationDiag 케이스 — 24h 내 영수증으로 네트워크 없이 판정(ok 통과, broken 진단), 메시지 분기 검증
package gate

import (
	"strings"
	"testing"
	"time"
)

func TestCitationDiag(t *testing.T) {
	// the URL is unreachable on purpose: a fresh receipt must prevent any fetch
	c := Citation{Label: "x", URL: "http://127.0.0.1:1/dead", Line: 7}
	now := time.Now()
	rcpts := map[string]receipt{c.URL: {CheckedAt: now, Verdict: "ok"}}
	if d, bad := citationDiag("f.md", c, rcpts, citationHTTP, now); bad {
		t.Errorf("fresh ok receipt must pass without network, got %+v", d)
	}
	rcpts[c.URL] = receipt{CheckedAt: now, Verdict: "broken", Detail: "HTTP 404"}
	d, bad := citationDiag("f.md", c, rcpts, citationHTTP, now)
	if !bad || d.Line != 7 || !strings.Contains(d.Message, "HTTP 404") {
		t.Errorf("fresh broken receipt: got (%+v, %v), want a citation-exists diagnostic at line 7", d, bad)
	}
	rcpts[c.URL] = receipt{CheckedAt: now, Verdict: "meta-mismatch", Detail: "title/og:title does not overlap citation label x"}
	if d, bad := citationDiag("f.md", c, rcpts, citationHTTP, now); !bad || !strings.Contains(d.Message, "og:title") {
		t.Errorf("meta-mismatch receipt: got (%+v, %v)", d, bad)
	}
}
