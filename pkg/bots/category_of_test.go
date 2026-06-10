//ff:func feature=bots type=dict control=iteration dimension=1
//ff:what CategoryOf가 대소문자 무시 조회로 분류를 반환하고 미등록 UA는 false를 내는지 검증
package bots

import "testing"

func TestCategoryOf(t *testing.T) {
	cases := []struct {
		name         string
		userAgent    string
		wantCategory string
		wantOK       bool
	}{
		{"training bot", "GPTBot", "training", true},
		{"search bot", "OAI-SearchBot", "search", true},
		{"fetcher", "Claude-User", "fetch", true},
		{"case insensitive", "claudebot", "training", true},
		{"unknown", "Googlebot", "", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkCategoryOf(t, tc.userAgent, tc.wantCategory, tc.wantOK) })
	}
}
