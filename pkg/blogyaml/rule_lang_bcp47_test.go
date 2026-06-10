//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleLangBCP47이 빈 languages와 잘못된 BCP-47 코드를 거부하고 유효 코드를 통과시키는지 검증
package blogyaml

import "testing"

func TestRuleLangBCP47(t *testing.T) {
	cases := []struct {
		name      string
		languages []string
		wantDiags int
	}{
		{"empty list", nil, 1},
		{"valid codes", []string{"ko", "en", "zh-Hans", "pt-BR"}, 0},
		{"one invalid", []string{"ko", "not_a_lang!"}, 1},
		{"two invalid", []string{"!!", "ko", "??"}, 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkRuleLangBCP47(t, tc.languages, tc.wantDiags) })
	}
}
