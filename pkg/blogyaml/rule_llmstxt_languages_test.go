//ff:func feature=blogyaml type=rule control=sequence
//ff:what ruleLlmsTxtLanguages가 base/all/선언 부분집합을 통과시키고 미선언 언어를 룰ID·라인과 함께 거부하는지 검증
package blogyaml

import (
	"strings"
	"testing"
)

func TestRuleLlmsTxtLanguages(t *testing.T) {
	blogWith := func(langs []string) *Blog {
		return &Blog{Languages: []string{"ko", "en"}, Geo: Geo{LlmsTxt: LlmsTxtSpec{Languages: langs}}}
	}
	bad := blogWith([]string{"ko", "fr"})
	diags := ruleLlmsTxtLanguages("blog.yaml", bad, lineIndex{"geo.llms_txt.languages[1]": 7})
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %v", diags)
	}
	if diags[0].Rule != "llmstxt-languages" || diags[0].Line != 7 {
		t.Errorf("want rule llmstxt-languages line 7, got %+v", diags[0])
	}
	if !strings.Contains(diags[0].Message, `"fr"`) {
		t.Errorf("want offending language in message, got %q", diags[0].Message)
	}
	if diags := ruleLlmsTxtLanguages("blog.yaml", blogWith([]string{"base"}), lineIndex{}); len(diags) != 0 {
		t.Errorf("want 0 diagnostics for base, got %v", diags)
	}
	if diags := ruleLlmsTxtLanguages("blog.yaml", blogWith([]string{"all"}), lineIndex{}); len(diags) != 0 {
		t.Errorf("want 0 diagnostics for all, got %v", diags)
	}
	if diags := ruleLlmsTxtLanguages("blog.yaml", blogWith([]string{"en"}), lineIndex{}); len(diags) != 0 {
		t.Errorf("want 0 diagnostics for declared subset, got %v", diags)
	}
}
