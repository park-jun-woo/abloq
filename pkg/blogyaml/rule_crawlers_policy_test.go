//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleCrawlersPolicy가 allow|block 외의 정책 값만 거부하고 진단을 키 정렬 순서로 내는지 검증
package blogyaml

import "testing"

func TestRuleCrawlersPolicy(t *testing.T) {
	cases := []struct {
		name      string
		crawlers  map[string]string
		wantDiags int
		wantFirst string
	}{
		{"nil map", nil, 0, ""},
		{"all valid", map[string]string{"gptbot": "allow", "bytespider": "block"}, 0, ""},
		{"one invalid", map[string]string{"gptbot": "deny"}, 1, "geo.crawlers.gptbot"},
		{"two invalid sorted", map[string]string{"zbot": "maybe", "abot": "deny", "gptbot": "allow"}, 2, "geo.crawlers.abot"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkRuleCrawlersPolicy(t, tc.crawlers, tc.wantDiags, tc.wantFirst) })
	}
}
