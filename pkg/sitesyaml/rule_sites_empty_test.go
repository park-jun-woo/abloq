//ff:func feature=sitesyaml type=rule control=sequence
//ff:what ruleSitesEmpty가 빈 리스트·키 부재를 거부하고 1개 이상이면 통과시키는지 검증
package sitesyaml

import "testing"

func TestRuleSitesEmpty(t *testing.T) {
	s, idx, diags := Parse("sites.yaml", []byte("sites: []\n"))
	if len(diags) != 0 {
		t.Fatalf("parse: %v", diags)
	}
	got := ruleSitesEmpty("sites.yaml", s, idx)
	if len(got) != 1 || got[0].Rule != "sites-empty" || got[0].Line != 1 {
		t.Errorf("empty list = %v, want one sites-empty diagnostic at line 1", got)
	}

	ok := &Sites{Sites: []Site{{Name: "a"}}}
	if got := ruleSitesEmpty("sites.yaml", ok, lineIndex{}); got != nil {
		t.Errorf("non-empty list = %v, want nil", got)
	}
}
