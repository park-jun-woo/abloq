//ff:func feature=archive type=client control=sequence
//ff:what targetList가 target URL 목록을 입력 순서대로 추출하는지 검증
package archive

import "testing"

func TestTargetList(t *testing.T) {
	urls := targetList([]Pending{{Target: "https://a/"}, {Target: "https://b/"}})
	if len(urls) != 2 || urls[0] != "https://a/" || urls[1] != "https://b/" {
		t.Errorf("targetList = %v, want [https://a/ https://b/]", urls)
	}
	if len(targetList(nil)) != 0 {
		t.Error("empty input must produce an empty list")
	}
}
