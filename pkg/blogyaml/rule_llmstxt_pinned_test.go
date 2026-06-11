//ff:func feature=blogyaml type=rule control=sequence
//ff:what ruleLlmsTxtPinned가 title·url 누락과 형식 위반 url을 거부하고 절대 URL·"/" 경로를 통과시키는지 검증
package blogyaml

import "testing"

func TestRuleLlmsTxtPinned(t *testing.T) {
	bad := &Blog{Geo: Geo{LlmsTxt: LlmsTxtSpec{Pinned: []LlmsPinned{{Title: "", URL: "reins.md"}}}}}
	idx := lineIndex{"geo.llms_txt.pinned[0]": 6, "geo.llms_txt.pinned[0].url": 7}
	diags := ruleLlmsTxtPinned("blog.yaml", bad, idx)
	if len(diags) != 2 {
		t.Fatalf("want 2 diagnostics (title required, url format), got %v", diags)
	}
	if diags[0].Rule != "llmstxt-pinned" || diags[0].Line != 6 {
		t.Errorf("title diagnostic = %+v, want rule llmstxt-pinned line 6", diags[0])
	}
	if diags[1].Rule != "llmstxt-pinned" || diags[1].Line != 7 {
		t.Errorf("url diagnostic = %+v, want rule llmstxt-pinned line 7", diags[1])
	}
	missing := &Blog{Geo: Geo{LlmsTxt: LlmsTxtSpec{Pinned: []LlmsPinned{{Title: "T"}}}}}
	if diags := ruleLlmsTxtPinned("blog.yaml", missing, lineIndex{}); len(diags) != 1 {
		t.Errorf("want 1 diagnostic for missing url, got %v", diags)
	}
	ok := &Blog{Geo: Geo{LlmsTxt: LlmsTxtSpec{Pinned: []LlmsPinned{
		{Title: "Root", URL: "/reins.md"},
		{Title: "Abs", URL: "https://example.com/x"},
	}}}}
	if diags := ruleLlmsTxtPinned("blog.yaml", ok, lineIndex{}); len(diags) != 0 {
		t.Errorf("want 0 diagnostics for valid pinned, got %v", diags)
	}
}
