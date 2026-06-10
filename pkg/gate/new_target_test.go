//ff:func feature=gate type=frame control=sequence
//ff:what NewTarget이 입력을 보존하고 헤딩 인덱스를 파생해 채우는지 검증
package gate

import "testing"

func TestNewTarget(t *testing.T) {
	b := loadGateBlog(t)
	a := artFromMD(t, b, "en", "tech", "pass", "articles/pass.md")
	tgt := NewTarget("testdata", b, []*Article{a})
	if tgt.Dir != "testdata" || tgt.Blog != b || len(tgt.Articles) != 1 {
		t.Fatalf("target fields not preserved: %+v", tgt)
	}
	if tgt.heads.rank["related"] != 3 || tgt.heads.rank["changelog"] != 6 {
		t.Errorf("heads.rank = %v, want order indexes", tgt.heads.rank)
	}
	if tgt.heads.byLang["ko"]["출처"] != "sources" {
		t.Errorf("heads.byLang missing ko sources: %v", tgt.heads.byLang)
	}
}
