//ff:func feature=bots type=dict control=sequence
//ff:what CategoryOf 케이스 하나를 실행해 분류와 ok 반환값을 검증
package bots

import "testing"

func checkCategoryOf(t *testing.T, userAgent, wantCategory string, wantOK bool) {
	t.Helper()
	category, ok := CategoryOf(userAgent)
	if ok != wantOK {
		t.Fatalf("CategoryOf(%q) ok = %v, want %v", userAgent, ok, wantOK)
	}
	if category != wantCategory {
		t.Errorf("CategoryOf(%q) = %q, want %q", userAgent, category, wantCategory)
	}
}
