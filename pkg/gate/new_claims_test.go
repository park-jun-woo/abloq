//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what newClaims가 미변경 글(Base==Doc)·기존 주장 보존·신규 주장 검출·신규 글(Base nil) 전수 검출을 가르는지 검증
package gate

import "testing"

func TestNewClaims(t *testing.T) {
	b := loadGateBlog(t)
	hi := buildHeadingIndex(b)
	oldBody := "---\ntitle: t\n---\n\nThroughput improved by 42% after the change.\n"
	newBody := "---\ntitle: t\n---\n\nThroughput improved by 42% after the change.\n\nLatency dropped by 12 ms in the same run.\n"
	doc := parseDoc(hi, "en", newBody)
	base := parseDoc(hi, "en", oldBody)

	unchanged := &Article{Lang: "en", Doc: doc, Base: doc}
	if got := newClaims(unchanged); len(got) != 0 {
		t.Errorf("unchanged article: want 0 new claims, got %d", len(got))
	}
	changed := &Article{Lang: "en", Doc: doc, Base: base}
	got := newClaims(changed)
	if len(got) != 1 {
		t.Fatalf("changed article: want 1 new claim, got %d: %v", len(got), got)
	}
	if got[0].Sourced {
		t.Errorf("new claim should be unsourced, got %+v", got[0])
	}
	fresh := &Article{Lang: "en", Doc: doc}
	if got := newClaims(fresh); len(got) != len(DetectClaims(doc)) {
		t.Errorf("new article: want every claim new (%d), got %d", len(DetectClaims(doc)), len(got))
	}
}
