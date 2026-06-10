//ff:func feature=gate type=rule control=sequence
//ff:what missingLangDiags가 누락 언어마다 진단 1건을 만들고 완전한 그룹에 0건을 반환하는지 검증
package gate

import "testing"

func TestMissingLangDiags(t *testing.T) {
	group := []*Article{{Lang: "ko", Path: "content/ko/tech/a.md"}}
	diags := missingLangDiags([]string{"ko", "en", "ja"}, "tech/a", group)
	if len(diags) != 2 {
		t.Fatalf("want 2 diagnostics, got %v", diags)
	}
	checkDiags(t, diags[:1], 1, "slug-consistency", "no en version")
	full := []*Article{{Lang: "ko"}, {Lang: "en"}, {Lang: "ja"}}
	if got := missingLangDiags([]string{"ko", "en", "ja"}, "tech/a", full); len(got) != 0 {
		t.Errorf("complete group: want 0, got %v", got)
	}
}
