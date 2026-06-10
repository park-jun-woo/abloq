//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what unquote가 %XX만 풀고 비정상 시퀀스·'+'는 원문 유지하는지 검증 (python unquote 동등성)
package cflog

import "testing"

func TestUnquote(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Mozilla/5.0%20(compatible;%20Bot/1.0)", "Mozilla/5.0 (compatible; Bot/1.0)"},
		{"/tech/post%2Da/", "/tech/post-a/"},
		{"a+b", "a+b"},
		{"100%", "100%"},
		{"%G1", "%G1"},
		{"%2", "%2"},
		{"", ""},
		{"%ed%95%9c", "한"},
	}
	for _, c := range cases {
		if got := unquote(c.in); got != c.want {
			t.Errorf("unquote(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
