//ff:func feature=scan type=parser control=sequence topic=evidence
//ff:what articleCites가 글 1편의 인용 URL을 등장 순서대로, 글 내 중복 없이 cite로 변환하는지 검증
package evidence

import "testing"

func TestArticleCites(t *testing.T) {
	b := testBlog(t)
	body := "---\ntitle: T\n---\n\n" +
		"[하나](https://a.example/1) 그리고 [둘](https://b.example/2)\n\n" +
		"[하나 again](https://a.example/1)\n"
	cites := articleCites(testArticle(t, b, body))
	if len(cites) != 2 {
		t.Fatalf("want 2 unique cites, got %d: %+v", len(cites), cites)
	}
	if cites[0].URL != "https://a.example/1" || cites[1].URL != "https://b.example/2" {
		t.Errorf("document order broken: %+v", cites)
	}
	if cites[0].Lang != "ko" || cites[0].Section != "tech" || cites[0].Slug != "fixture" {
		t.Errorf("cite coordinates: %+v", cites[0])
	}
}
