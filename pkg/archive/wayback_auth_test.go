//ff:func feature=archive type=client control=sequence
//ff:what waybackAuth가 두 키가 모두 있을 때만 "LOW key:secret"을 만들고 아니면 빈 문자열인지 검증
package archive

import "testing"

func TestWaybackAuth(t *testing.T) {
	t.Setenv("WAYBACK_ACCESS_KEY", "ak")
	t.Setenv("WAYBACK_SECRET_KEY", "sk")
	if got := waybackAuth(); got != "LOW ak:sk" {
		t.Errorf("waybackAuth = %q, want LOW ak:sk", got)
	}
	t.Setenv("WAYBACK_SECRET_KEY", "")
	if got := waybackAuth(); got != "" {
		t.Errorf("waybackAuth without secret = %q, want empty", got)
	}
	t.Setenv("WAYBACK_ACCESS_KEY", "")
	if got := waybackAuth(); got != "" {
		t.Errorf("waybackAuth without keys = %q, want empty", got)
	}
}
