//ff:func feature=gate type=rule control=sequence topic=evidence
//ff:what citationDiags가 인용 목록에서 실패 건만 진단으로 모으는지 검증 — ok 영수증 1건 + broken 영수증 1건
package gate

import (
	"testing"
	"time"
)

func TestCitationDiags(t *testing.T) {
	cs := []Citation{
		{Label: "a", URL: "http://127.0.0.1:1/a", Line: 3},
		{Label: "b", URL: "http://127.0.0.1:1/b", Line: 9},
	}
	rcpts := map[string]receipt{
		cs[0].URL: {CheckedAt: time.Now(), Verdict: "ok"},
		cs[1].URL: {CheckedAt: time.Now(), Verdict: "broken", Detail: "HTTP 404"},
	}
	diags := citationDiags("f.md", cs, rcpts, citationHTTP)
	checkDiags(t, diags, 1, "citation-exists", "HTTP 404")
	if diags[0].Line != 9 {
		t.Errorf("want line 9, got %d", diags[0].Line)
	}
}
