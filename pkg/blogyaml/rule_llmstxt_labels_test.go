//ff:func feature=blogyaml type=rule control=sequence
//ff:what ruleLlmsTxtLabels가 미선언 섹션 키를 룰ID·라인과 함께 거부하고 선언된 섹션 키를 통과시키는지 검증
package blogyaml

import (
	"strings"
	"testing"
)

func TestRuleLlmsTxtLabels(t *testing.T) {
	bad := &Blog{
		Sections: []string{"opinion", "tech"},
		Geo:      Geo{LlmsTxt: LlmsTxtSpec{SectionLabels: map[string]string{"tech": "Pattern", "blog": "X"}}},
	}
	diags := ruleLlmsTxtLabels("blog.yaml", bad, lineIndex{"geo.llms_txt.section_labels.blog": 9})
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %v", diags)
	}
	if diags[0].Rule != "llmstxt-labels" || diags[0].Line != 9 {
		t.Errorf("want rule llmstxt-labels line 9, got %+v", diags[0])
	}
	if !strings.Contains(diags[0].Message, "section_labels.blog") {
		t.Errorf("want offending key in message, got %q", diags[0].Message)
	}
	ok := &Blog{
		Sections: []string{"opinion", "tech"},
		Geo:      Geo{LlmsTxt: LlmsTxtSpec{SectionLabels: map[string]string{"opinion": "Concept"}}},
	}
	if diags := ruleLlmsTxtLabels("blog.yaml", ok, lineIndex{}); len(diags) != 0 {
		t.Errorf("want 0 diagnostics for declared section keys, got %v", diags)
	}
}
