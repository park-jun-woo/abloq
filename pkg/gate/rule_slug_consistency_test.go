//ff:func feature=gate type=rule control=sequence
//ff:what [slug-consistency] 전 언어 존재+slug 일치 PASS, 누락 언어 FAIL, slug 불일치 FAIL 검증
package gate

import "testing"

func TestRuleSlugConsistency(t *testing.T) {
	b := loadGateBlog(t)
	ko := artFromMD(t, b, "ko", "tech", "hello", "repo-pass/content/ko/tech/hello.md")
	en := artFromMD(t, b, "en", "tech", "hello", "repo-pass/content/en/tech/hello.md")
	if diags := ruleSlugConsistency(NewTarget("testdata", b, []*Article{ko, en})); len(diags) != 0 {
		t.Fatalf("complete pair: want 0 diagnostics, got %v", diags)
	}
	missing := ruleSlugConsistency(NewTarget("testdata", b, []*Article{ko}))
	checkDiags(t, missing, 1, "slug-consistency", "no en version")

	en.Doc.FrontMatter += "\nslug: \"other-slug\""
	mismatch := ruleSlugConsistency(NewTarget("testdata", b, []*Article{ko, en}))
	checkDiags(t, mismatch, 1, "slug-consistency", "differs from")
}
