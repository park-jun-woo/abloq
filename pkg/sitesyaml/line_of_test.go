//ff:func feature=sitesyaml type=parser control=sequence topic=diagnostics
//ff:what lineOf가 등록된 경로의 라인을 주고 미등록 경로는 1로 폴백하는지 검증
package sitesyaml

import "testing"

func TestLineOf(t *testing.T) {
	idx := lineIndex{"sites[0].name": 4}
	if got := lineOf(idx, "sites[0].name"); got != 4 {
		t.Errorf("known path = %d, want 4", got)
	}
	if got := lineOf(idx, "sites[9].name"); got != 1 {
		t.Errorf("unknown path = %d, want fallback 1", got)
	}
}
