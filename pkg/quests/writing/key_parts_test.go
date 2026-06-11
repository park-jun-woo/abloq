//ff:func feature=quest type=parser control=sequence
//ff:what keyParts 검증 — 플랫 글/번들 경로의 lang·section·slug 파생, content 밖·깊이 불일치는 false
package writing

import "testing"

func TestKeyParts(t *testing.T) {
	lang, section, slug, ok := keyParts("content/en/posts/hello.md")
	if !ok || lang != "en" || section != "posts" || slug != "hello" {
		t.Errorf("flat = %s/%s/%s ok=%v", lang, section, slug, ok)
	}
	lang, section, slug, ok = keyParts("content/ko/tech/bundle/index.md")
	if !ok || lang != "ko" || section != "tech" || slug != "bundle" {
		t.Errorf("bundle = %s/%s/%s ok=%v", lang, section, slug, ok)
	}
	if _, _, _, ok := keyParts("static/en/posts/hello.md"); ok {
		t.Error("non-content: want ok=false")
	}
	if _, _, _, ok := keyParts("content/en/hello.md"); ok {
		t.Error("too shallow: want ok=false")
	}
}
