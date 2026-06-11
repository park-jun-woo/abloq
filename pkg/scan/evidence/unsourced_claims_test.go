//ff:func feature=scan type=parser control=sequence topic=evidence
//ff:what unsourcedClaims 케이스 — 무출처 주장만 해시·path:line·원문으로 수집, 출처 있는 문단·claims_ignore 예외는 제외
package evidence

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/gate"
)

func TestUnsourcedClaims(t *testing.T) {
	b := testBlog(t)
	body := "---\ntitle: T\n---\n\n" +
		"처리량이 40% 증가했다.\n\n" +
		"지연이 120ms 단축됐다. [근거](https://src.example/x)\n"
	refs := unsourcedClaims(testArticle(t, b, body))
	if len(refs) != 1 {
		t.Fatalf("want 1 unsourced claim, got %d: %+v", len(refs), refs)
	}
	if refs[0].Text != "처리량이 40% 증가했다." {
		t.Errorf("text = %q", refs[0].Text)
	}
	if refs[0].Loc != "content/ko/tech/fixture.md:5" {
		t.Errorf("loc must be repo-relative path:line, got %q", refs[0].Loc)
	}
	if refs[0].Hash != gate.HashText(refs[0].Text) {
		t.Errorf("hash must key on the claim text")
	}
	exempt := "---\ntitle: T\nclaims_ignore:\n  - \"own benchmark\"\n---\n\n처리량이 40% 증가했다.\n"
	if got := unsourcedClaims(testArticle(t, b, exempt)); got != nil {
		t.Errorf("claims_ignore must exempt the article: %+v", got)
	}
}
