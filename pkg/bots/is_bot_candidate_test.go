//ff:func feature=bots type=dict control=iteration dimension=1 topic=crawl
//ff:what IsBotCandidate가 봇 토큰·비브라우저 클라이언트는 후보로, 일반 브라우저·빈 UA는 비후보로 가르는지 검증
package bots

import "testing"

func TestIsBotCandidate(t *testing.T) {
	cases := []struct {
		ua   string
		want bool
	}{
		{"Mozilla/5.0 (compatible;PetalBot;+https://webmaster.petalsearch.com/site/petalbot)", true},
		{"curl/8.5.0", true},
		{"Googlebot/2.1 (+http://www.google.com/bot.html)", true},
		{"python-requests/2.32.0", true},
		{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36", false},
		{"Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1", false},
		{"", false},
		{"-", false},
	}
	for _, c := range cases {
		if got := IsBotCandidate(c.ua); got != c.want {
			t.Errorf("IsBotCandidate(%q) = %v, want %v", c.ua, got, c.want)
		}
	}
}
