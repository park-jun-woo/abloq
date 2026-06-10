//ff:func feature=blogyaml type=rule control=sequence
//ff:what ruleSectionsEmpty 케이스 하나를 실행해 진단 수·룰ID·라인을 검증
package blogyaml

import "testing"

func checkRuleSectionsEmpty(t *testing.T, sections []string, wantDiags int) {
	t.Helper()
	b := &Blog{Sections: sections}
	diags := ruleSectionsEmpty("blog.yaml", b, lineIndex{"sections": 4})
	if len(diags) != wantDiags {
		t.Fatalf("want %d diagnostics, got %d: %v", wantDiags, len(diags), diags)
	}
	if wantDiags == 1 {
		if diags[0].Rule != "sections-empty" {
			t.Errorf("want rule sections-empty, got %q", diags[0].Rule)
		}
		if diags[0].Line != 4 {
			t.Errorf("want line 4, got %d", diags[0].Line)
		}
	}
}
