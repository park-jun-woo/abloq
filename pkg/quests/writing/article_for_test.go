//ff:func feature=quest type=parser control=sequence
//ff:what articleFor 검증 — 번들 insight.yaml→index.md, 사이드카 *.insight.yaml→*.md, 그 외 false
package writing

import "testing"

func TestArticleFor(t *testing.T) {
	got, ok := articleFor("content/en/posts/b/insight.yaml")
	if !ok || got != "content/en/posts/b/index.md" {
		t.Errorf("bundle = %q ok=%v", got, ok)
	}
	got, ok = articleFor("content/en/posts/a.insight.yaml")
	if !ok || got != "content/en/posts/a.md" {
		t.Errorf("sidecar = %q ok=%v", got, ok)
	}
	if _, ok := articleFor("content/en/posts/a.yaml"); ok {
		t.Error("non-sidecar: want ok=false")
	}
}
