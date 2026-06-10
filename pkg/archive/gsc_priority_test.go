//ff:func feature=archive type=client control=sequence
//ff:what gscPriority가 신규 0·갱신 1·미인덱스 2를 매기는지 검증
package archive

import "testing"

func TestGscPriority(t *testing.T) {
	if got := gscPriority(Pending{Date: "2026-06-01", Lastmod: "2026-06-01"}); got != 0 {
		t.Errorf("new post = %d, want 0", got)
	}
	if got := gscPriority(Pending{Date: "2026-06-01", Lastmod: "2026-06-05"}); got != 1 {
		t.Errorf("updated post = %d, want 1", got)
	}
	if got := gscPriority(Pending{}); got != 2 {
		t.Errorf("unindexed target = %d, want 2", got)
	}
}
