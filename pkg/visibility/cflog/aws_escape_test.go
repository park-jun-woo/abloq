//ff:func feature=visibility type=client control=iteration dimension=1 topic=crawl
//ff:what awsEscape가 비예약 문자만 남기고 대문자 %XX로 인코딩하는지, keepSlash가 '/'를 보존하는지 검증
package cflog

import "testing"

func TestAwsEscape(t *testing.T) {
	cases := []struct {
		in        string
		keepSlash bool
		want      string
	}{
		{"AZaz09-._~", false, "AZaz09-._~"},
		{"a b", false, "a%20b"},
		{"a/b", false, "a%2Fb"},
		{"a/b", true, "a/b"},
		{"한", false, "%ED%95%9C"},
		{"a+b=c", false, "a%2Bb%3Dc"},
	}
	for _, c := range cases {
		if got := awsEscape(c.in, c.keepSlash); got != c.want {
			t.Errorf("awsEscape(%q, %v) = %q, want %q", c.in, c.keepSlash, got, c.want)
		}
	}
}
