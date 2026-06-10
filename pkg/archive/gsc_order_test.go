//ff:func feature=archive type=client control=iteration dimension=1
//ff:what gscOrder가 신규 우선 안정 정렬 사본을 돌려주고 원본을 보존하는지 검증
package archive

import "testing"

func TestGscOrder(t *testing.T) {
	pending := []Pending{
		{Target: "https://u1/", Date: "2026-01-01", Lastmod: "2026-02-01"},
		{Target: "https://n1/", Date: "2026-01-01", Lastmod: "2026-01-01"},
		{Target: "https://x1/"},
		{Target: "https://n2/", Date: "2026-03-01", Lastmod: "2026-03-01"},
	}
	ordered := gscOrder(pending)
	want := []string{"https://n1/", "https://n2/", "https://u1/", "https://x1/"}
	for i, w := range want {
		if ordered[i].Target != w {
			t.Errorf("ordered[%d] = %s, want %s", i, ordered[i].Target, w)
		}
	}
	if pending[0].Target != "https://u1/" {
		t.Error("input slice must not be reordered")
	}
}
