//ff:func feature=blogyaml type=rule control=sequence
//ff:what ruleBaseURLFormat 케이스 하나를 실행해 진단 유무·룰ID·라인·메시지를 검증
package blogyaml

import (
	"strings"
	"testing"
)

func checkRuleBaseURLFormat(t *testing.T, baseURL, wantMsgPart string, wantDiag bool) {
	t.Helper()
	b := &Blog{Site: Site{BaseURL: baseURL}}
	diags := ruleBaseURLFormat("blog.yaml", b, lineIndex{"site.baseURL": 2})
	if !wantDiag {
		if len(diags) != 0 {
			t.Fatalf("want 0 diagnostics, got %v", diags)
		}
		return
	}
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %d: %v", len(diags), diags)
	}
	if diags[0].Rule != "baseurl-format" {
		t.Errorf("want rule baseurl-format, got %q", diags[0].Rule)
	}
	if diags[0].Line != 2 {
		t.Errorf("want line 2, got %d", diags[0].Line)
	}
	if !strings.Contains(diags[0].Message, wantMsgPart) {
		t.Errorf("want message containing %q, got %q", wantMsgPart, diags[0].Message)
	}
}
