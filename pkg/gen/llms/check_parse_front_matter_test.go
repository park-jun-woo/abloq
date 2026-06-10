//ff:func feature=gen type=parser control=sequence
//ff:what parseFrontMatter 케이스 하나를 실행해 ok/title/draft 디코드 결과를 검증
package llms

import "testing"

func checkParseFrontMatter(t *testing.T, data string, wantOK bool, wantTitle string, wantDraft bool) {
	t.Helper()
	fm, ok := parseFrontMatter([]byte(data))
	if ok != wantOK {
		t.Fatalf("parseFrontMatter ok = %v, want %v (input %q)", ok, wantOK, data)
	}
	if fm.Title != wantTitle {
		t.Errorf("title = %q, want %q", fm.Title, wantTitle)
	}
	if fm.Draft != wantDraft {
		t.Errorf("draft = %v, want %v", fm.Draft, wantDraft)
	}
}
