//ff:func feature=sitesyaml type=rule control=iteration dimension=1 topic=gsc
//ff:what gscSiteURLProblem이 빈 값·sc-domain:도메인·http(s) URL을 통과시키고 빈 sc-domain·비http 스킴·호스트 없음을 거부하는지 검증
package sitesyaml

import "testing"

func TestGSCSiteURLProblem(t *testing.T) {
	for _, ok := range []string{"", "sc-domain:example.com", "https://example.com/", "http://example.com"} {
		if msg := gscSiteURLProblem(ok); msg != "" {
			t.Errorf("gscSiteURLProblem(%q) = %q, want legal", ok, msg)
		}
	}
	for _, bad := range []string{"sc-domain:", "ftp://example.com", "example.com/path", "https://", "://broken"} {
		if msg := gscSiteURLProblem(bad); msg == "" {
			t.Errorf("gscSiteURLProblem(%q) = legal, want a problem", bad)
		}
	}
}
