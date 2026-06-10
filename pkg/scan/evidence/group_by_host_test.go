//ff:func feature=scan type=client control=sequence topic=evidence
//ff:what groupByHost가 호스트별로 묶고 중복 URL을 제거하며 파싱 불가 URL을 "" 버킷에 두는지 검증
package evidence

import "testing"

func TestGroupByHost(t *testing.T) {
	groups := groupByHost([]string{
		"https://a.example/one",
		"https://a.example/two",
		"https://a.example/one", // duplicate
		"https://b.example/x",
		"://bad",
	})
	if len(groups["a.example"]) != 2 {
		t.Errorf("a.example = %v, want 2 unique URLs", groups["a.example"])
	}
	if len(groups["b.example"]) != 1 {
		t.Errorf("b.example = %v", groups["b.example"])
	}
	if len(groups[""]) != 1 {
		t.Errorf("unparseable bucket = %v, want the bad URL kept", groups[""])
	}
}
