//ff:func feature=blogyaml type=rule control=sequence
//ff:what ruleLlmsTxtMode가 auto/manual/off(와 미주입 영값)를 통과시키고 그 외 값을 룰ID·라인과 함께 거부하는지 검증
package blogyaml

import (
	"strings"
	"testing"
)

func TestRuleLlmsTxtMode(t *testing.T) {
	bad := &Blog{Geo: Geo{LlmsTxt: LlmsTxtSpec{Mode: "bogus"}}}
	diags := ruleLlmsTxtMode("blog.yaml", bad, lineIndex{"geo.llms_txt": 4})
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %v", diags)
	}
	if diags[0].Rule != "llmstxt-mode" || diags[0].Line != 4 {
		t.Errorf("want rule llmstxt-mode line 4, got %+v", diags[0])
	}
	if !strings.Contains(diags[0].Message, "auto, manual, off") {
		t.Errorf("want enum listed in message, got %q", diags[0].Message)
	}
	objLine := lineIndex{"geo.llms_txt": 4, "geo.llms_txt.mode": 5}
	if diags := ruleLlmsTxtMode("blog.yaml", bad, objLine); diags[0].Line != 5 {
		t.Errorf("object form must point at the mode key, got line %d", diags[0].Line)
	}
	ok := &Blog{Geo: Geo{LlmsTxt: LlmsTxtSpec{Mode: "manual"}}}
	if diags := ruleLlmsTxtMode("blog.yaml", ok, lineIndex{}); len(diags) != 0 {
		t.Errorf("want 0 diagnostics for manual, got %v", diags)
	}
	zero := &Blog{}
	if diags := ruleLlmsTxtMode("blog.yaml", zero, lineIndex{}); len(diags) != 0 {
		t.Errorf("zero-value mode must normalize to auto, got %v", diags)
	}
}
