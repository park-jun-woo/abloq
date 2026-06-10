//ff:func feature=visibility type=client control=iteration dimension=1 topic=crawl
//ff:what canonicalQuery가 키·값 인코딩과 사전순 정렬로 정준 쿼리를 만드는지 검증 — 공백은 %20, 빈 값 유지
package cflog

import "testing"

func TestCanonicalQuery(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Action=ListUsers&Version=2010-05-08", "Action=ListUsers&Version=2010-05-08"},
		{"b=2&a=1", "a=1&b=2"},
		{"prefix=logs%2F2026&list-type=2", "list-type=2&prefix=logs%2F2026"},
		{"k=a b", "k=a%20b"},
		{"empty=", "empty="},
		{"", ""},
	}
	for _, c := range cases {
		if got := canonicalQuery(c.in); got != c.want {
			t.Errorf("canonicalQuery(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
