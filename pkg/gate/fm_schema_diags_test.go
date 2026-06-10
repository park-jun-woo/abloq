//ff:func feature=gate type=rule control=sequence
//ff:what fmSchemaDiags가 title 공백/빈 tags를 각각 키 라인 번호와 함께 진단하는지 검증
package gate

import "testing"

func TestFMSchemaDiags(t *testing.T) {
	b := loadGateBlog(t)
	doc := ParseArticle(b, "en", "---\ntitle: \"\"\ndate: 2026-01-01\nlastmod: 2026-01-01\ntags: []\n---\nbody\n")
	a := &Article{Lang: "en", Section: "tech", Slug: "x", Path: "p.md", Doc: doc}
	diags := fmSchemaDiags(a)
	if len(diags) != 2 {
		t.Fatalf("want 2 diagnostics (title, tags), got %v", diags)
	}
	if diags[0].Line != 2 {
		t.Errorf("title diag line = %d, want 2", diags[0].Line)
	}
	if diags[1].Line != 5 {
		t.Errorf("tags diag line = %d, want 5", diags[1].Line)
	}
}
