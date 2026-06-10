//ff:func feature=gate type=frame control=sequence
//ff:what Run이 전체 룰을 실행하고 ruleIDs 지정 시 해당 룰만 실행하는지 검증
package gate

import "testing"

func TestRun(t *testing.T) {
	b := loadGateBlog(t)
	bad := artFromMD(t, b, "en", "tech", "no-image", "articles/no-image.md")
	ok := artFromMD(t, b, "ko", "tech", "no-image", "articles/pass-ko.md")
	tgt := NewTarget("testdata", b, []*Article{bad, ok})

	all := Run(tgt)
	if len(all) < 2 {
		t.Fatalf("want at least 2 diagnostics from all rules, got %v", all)
	}
	only := Run(tgt, "image-first")
	checkDiags(t, only, 1, "image-first", "first content line")
	if got := Run(tgt, "no-such-rule"); len(got) != 0 {
		t.Errorf("unknown rule id: want 0 diagnostics, got %v", got)
	}
}
