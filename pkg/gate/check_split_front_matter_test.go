//ff:func feature=gate type=parser control=sequence
//ff:what splitFrontMatter 케이스 하나를 실행해 fm/body/ok를 검증
package gate

import "testing"

func checkSplitFrontMatter(t *testing.T, in, wantFM, wantBody string, wantOK bool) {
	t.Helper()
	fm, body, ok := splitFrontMatter(in)
	if ok != wantOK {
		t.Fatalf("ok = %v, want %v", ok, wantOK)
	}
	if fm != wantFM {
		t.Errorf("fm = %q, want %q", fm, wantFM)
	}
	if body != wantBody {
		t.Errorf("body = %q, want %q", body, wantBody)
	}
}
