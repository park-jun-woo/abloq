//ff:func feature=sitesyaml type=parser control=sequence topic=diagnostics
//ff:what lineOfSite가 키 라인 → 항목 라인 → 1 순서로 폴백하는지 검증
package sitesyaml

import "testing"

func TestLineOfSite(t *testing.T) {
	idx := lineIndex{"sites[0]": 2, "sites[0].name": 3}
	if got := lineOfSite(idx, 0, "name"); got != 3 {
		t.Errorf("key line = %d, want 3", got)
	}
	if got := lineOfSite(idx, 0, "repo_path"); got != 2 {
		t.Errorf("absent key must fall back to the item line: %d, want 2", got)
	}
	if got := lineOfSite(idx, 5, "name"); got != 1 {
		t.Errorf("absent item must fall back to 1: %d", got)
	}
}
