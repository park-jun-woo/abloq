//ff:func feature=insight type=rule control=sequence
//ff:what Match 검증 — front matter 속 어휘는 미출현, 본문 출현은 Found, 미출현은 Missing, 섹션 도출
package insight

import "testing"

func TestMatch(t *testing.T) {
	article := []byte("---\ntitle: \"only-in-front-matter\"\n---\n\nThe RATCHET never moves backward.\n")
	ins := &Insight{Section: "tech", Claims: []Claim{
		{ID: "hit", Anchors: []string{"ratchet never moves"}},
		{ID: "fm-only", Anchors: []string{"only-in-front-matter"}},
	}}
	res := Match(ins, "content/en/tech/post.md", article)
	if res.Section != "tech" {
		t.Errorf("want section tech from path, got %q", res.Section)
	}
	if len(res.Found) != 1 || res.Found[0] != "hit" {
		t.Errorf("want only claim hit found, got %v", res.Found)
	}
	if len(res.Missing) != 1 || res.Missing[0].ID != "fm-only" {
		t.Errorf("want front-matter-only claim in missing (front matter excluded), got %v", res.Missing)
	}
}
