//ff:func feature=content type=parser control=sequence
//ff:what parseFrontMatter가 정상 블록의 (front matter, 본문) 분리와 블록 부재·미종결·YAML 오류·종결 뒤 개행 없음 경계를 처리하는지 검증
package content

import "testing"

func TestParseFrontMatter(t *testing.T) {
	fm, body, ok := parseFrontMatter([]byte("---\ntitle: T\ndate: 2026-06-01\n---\nbody line\n"))
	if !ok || fm.Title != "T" || fm.Date != "2026-06-01" {
		t.Errorf("normal block = %+v ok=%v", fm, ok)
	}
	if body != "body line\n" {
		t.Errorf("body = %q", body)
	}
	if _, _, ok := parseFrontMatter([]byte("no front matter")); ok {
		t.Error("missing block must be !ok")
	}
	if _, _, ok := parseFrontMatter([]byte("---\ntitle: T\n")); ok {
		t.Error("unterminated block must be !ok")
	}
	if _, _, ok := parseFrontMatter([]byte("---\n: bad: [\n---\n")); ok {
		t.Error("invalid YAML must be !ok")
	}
	if _, body, ok := parseFrontMatter([]byte("---\ntitle: T\n---")); !ok || body != "" {
		t.Errorf("closing fence at EOF: body=%q ok=%v", body, ok)
	}
}
