//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ogNameSafe의 허용(소문자/숫자/하이픈/언더스코어)·거부(빈 문자열/대문자/공백/유니코드) 케이스 검증
package blogyaml

import "testing"

func TestOGNameSafe(t *testing.T) {
	for _, ok := range []string{"minimal", "photo-2", "a_b", "x0"} {
		if !ogNameSafe(ok) {
			t.Errorf("ogNameSafe(%q) = false, want true", ok)
		}
	}
	for _, bad := range []string{"", "Minimal", "a b", "사진", "a/b", "a.b"} {
		if ogNameSafe(bad) {
			t.Errorf("ogNameSafe(%q) = true, want false", bad)
		}
	}
}
