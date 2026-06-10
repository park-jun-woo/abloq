//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what tripletVal이 유효한 %XX만 디코드하고 비-%·범위 밖·비16진수는 false인지 검증
package cflog

import "testing"

func TestTripletVal(t *testing.T) {
	cases := []struct {
		s    string
		i    int
		want byte
		ok   bool
	}{
		{"%20", 0, 0x20, true},
		{"a%2Fb", 1, 0x2F, true},
		{"abc", 0, 0, false},
		{"%2", 0, 0, false},
		{"%G1", 0, 0, false},
		{"%1G", 0, 0, false},
	}
	for _, c := range cases {
		got, ok := tripletVal(c.s, c.i)
		if got != c.want || ok != c.ok {
			t.Errorf("tripletVal(%q, %d) = (%#x, %v), want (%#x, %v)", c.s, c.i, got, ok, c.want, c.ok)
		}
	}
}
