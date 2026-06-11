//ff:func feature=insight type=parser control=sequence
//ff:what front matter 분리 검증 — 블록 제거, 블록 없음·미종결은 전문 유지, 본문 없는 글은 빈 문자열
package insight

import "testing"

func TestStripFrontMatter(t *testing.T) {
	got := stripFrontMatter([]byte("---\ntitle: \"T\"\n---\n\nbody anchor\n"))
	if got != "\nbody anchor\n" {
		t.Errorf("want body without front matter, got %q", got)
	}
	if got := stripFrontMatter([]byte("no front matter\n")); got != "no front matter\n" {
		t.Errorf("want whole file as body when no block, got %q", got)
	}
	if got := stripFrontMatter([]byte("---\ntitle: T\nno end")); got != "---\ntitle: T\nno end" {
		t.Errorf("want whole file as body for unterminated block, got %q", got)
	}
	if got := stripFrontMatter([]byte("---\ntitle: T\n---")); got != "" {
		t.Errorf("want empty body when nothing follows the block, got %q", got)
	}
}
