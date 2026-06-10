//ff:func feature=gate type=parser control=iteration dimension=1
//ff:what fmLineValue가 키의 스칼라 값을 따옴표 제거 후 추출하고 부재 시 빈 문자열을 반환하는지 검증
package gate

import "testing"

func TestFMLineValue(t *testing.T) {
	fm := "title: \"Hello\"\nlastmod: 2026-01-02T00:00:00+09:00\nslug: 'my-slug'"
	cases := []struct{ name, key, want string }{
		{"quoted", "title", "Hello"},
		{"plain", "lastmod", "2026-01-02T00:00:00+09:00"},
		{"single quoted", "slug", "my-slug"},
		{"absent", "image", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := fmLineValue(fm, tc.key); got != tc.want {
				t.Errorf("fmLineValue(%s) = %q, want %q", tc.key, got, tc.want)
			}
		})
	}
}
