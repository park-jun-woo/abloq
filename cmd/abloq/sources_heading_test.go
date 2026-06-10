//ff:func feature=init type=dict control=iteration dimension=1
//ff:what sourcesHeading이 등록 언어의 현지화 헤딩을 주고 미등록 언어에 "Sources"로 폴백하는지 검증
package main

import "testing"

func TestSourcesHeading(t *testing.T) {
	cases := []struct{ lang, want string }{
		{"ko", "출처"}, {"en", "Sources"}, {"ja", "出典"}, {"unknown", "Sources"},
	}
	for _, c := range cases {
		if got := sourcesHeading(c.lang); got != c.want {
			t.Errorf("sourcesHeading(%s) = %q, want %q", c.lang, got, c.want)
		}
	}
}
