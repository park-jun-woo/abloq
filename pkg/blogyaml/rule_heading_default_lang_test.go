//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleHeadingDefaultLang이 기본 언어 누락 헤딩만 거부하고 languages가 비면 침묵(lang-bcp47 소관)하는지 검증
package blogyaml

import "testing"

func TestRuleHeadingDefaultLang(t *testing.T) {
	cases := []struct {
		name      string
		languages []string
		headings  map[string]map[string]string
		wantDiags int
	}{
		{"empty languages defers to lang-bcp47", nil, map[string]map[string]string{"tldr": {"en": "TL;DR"}}, 0},
		{"nil headings", []string{"ko"}, nil, 0},
		{"all cover default", []string{"ko", "en"}, map[string]map[string]string{"tldr": {"ko": "요약", "en": "TL;DR"}}, 0},
		{"missing default", []string{"ko", "en"}, map[string]map[string]string{"tldr": {"en": "TL;DR"}}, 1},
		{"two missing", []string{"ko"}, map[string]map[string]string{"tldr": {"en": "TL;DR"}, "faq": {"en": "FAQ"}}, 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkRuleHeadingDefaultLang(t, tc.languages, tc.headings, tc.wantDiags) })
	}
}
