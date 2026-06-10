//ff:func feature=gen type=rule control=iteration dimension=1 topic=drift
//ff:what ruleFor가 파생물 경로 4종을 게이트 룰ID로, 그 외를 derived-sync로 매핑하는지 검증
package gen

import "testing"

func TestRuleFor(t *testing.T) {
	cases := []struct {
		path string
		want string
	}{
		{"hugo.toml", "hugo-config-sync"},
		{"static/robots.txt", "robots-policy-match"},
		{"static/llms.txt", "llmstxt-sync"},
		{"data/jsonld.json", "jsonld-sync"},
		{"something/else.txt", "derived-sync"},
	}
	for _, tc := range cases {
		t.Run(tc.path, func(t *testing.T) { checkRuleFor(t, tc.path, tc.want) })
	}
}
