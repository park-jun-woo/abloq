//ff:func feature=blogyaml type=rule control=sequence
//ff:what ruleLlmsTxtMaxSummary가 음수를 룰ID·라인과 함께 거부하고 0(무제한)·양수를 통과시키는지 검증
package blogyaml

import "testing"

func TestRuleLlmsTxtMaxSummary(t *testing.T) {
	bad := &Blog{Geo: Geo{LlmsTxt: LlmsTxtSpec{MaxSummary: -1}}}
	diags := ruleLlmsTxtMaxSummary("blog.yaml", bad, lineIndex{"geo.llms_txt.max_summary": 8})
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %v", diags)
	}
	if diags[0].Rule != "llmstxt-max-summary" || diags[0].Line != 8 {
		t.Errorf("want rule llmstxt-max-summary line 8, got %+v", diags[0])
	}
	if diags := ruleLlmsTxtMaxSummary("blog.yaml", &Blog{}, lineIndex{}); len(diags) != 0 {
		t.Errorf("want 0 diagnostics for 0 (unlimited), got %v", diags)
	}
	pos := &Blog{Geo: Geo{LlmsTxt: LlmsTxtSpec{MaxSummary: 200}}}
	if diags := ruleLlmsTxtMaxSummary("blog.yaml", pos, lineIndex{}); len(diags) != 0 {
		t.Errorf("want 0 diagnostics for positive cap, got %v", diags)
	}
}
