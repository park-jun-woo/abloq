//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleBaseURLFormat이 빈 값/파싱 불가/비 http(s)/호스트 없음/query·fragment를 거부하고 정상 URL을 통과시키는지 검증
package blogyaml

import "testing"

func TestRuleBaseURLFormat(t *testing.T) {
	cases := []struct {
		name, baseURL, wantMsgPart string
		wantDiag                   bool
	}{
		{"empty", "", "is required", true},
		{"unparsable", "https://exa\nmple.com", "not a valid URL", true},
		{"bad scheme", "ftp://example.com", "must use http or https", true},
		{"no host", "https://", "must have a host", true},
		{"query", "https://example.com?utm=1", "must not have a query or fragment", true},
		{"fragment", "https://example.com#top", "must not have a query or fragment", true},
		{"valid https", "https://example.com", "", false},
		{"valid http with path", "http://example.com/blog", "", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkRuleBaseURLFormat(t, tc.baseURL, tc.wantMsgPart, tc.wantDiag) })
	}
}
