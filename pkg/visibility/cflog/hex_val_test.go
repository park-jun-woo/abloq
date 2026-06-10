//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what hexVal이 0-9/a-f/A-F만 값으로 풀고 그 외는 false인지 검증
package cflog

import "testing"

func TestHexVal(t *testing.T) {
	cases := []struct {
		in   byte
		want byte
		ok   bool
	}{
		{'0', 0, true}, {'9', 9, true}, {'a', 10, true}, {'f', 15, true},
		{'A', 10, true}, {'F', 15, true}, {'g', 0, false}, {'%', 0, false},
	}
	for _, c := range cases {
		got, ok := hexVal(c.in)
		if got != c.want || ok != c.ok {
			t.Errorf("hexVal(%q) = (%d, %v), want (%d, %v)", c.in, got, ok, c.want, c.ok)
		}
	}
}
