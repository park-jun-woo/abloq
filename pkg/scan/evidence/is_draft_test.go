//ff:func feature=scan type=parser control=sequence topic=evidence
//ff:what isDraft 케이스 — draft: true만 참, 키 없음·false·깨진 front matter는 발행 취급
package evidence

import "testing"

func TestIsDraft(t *testing.T) {
	b := testBlog(t)
	if !isDraft(testArticle(t, b, "---\ntitle: D\ndraft: true\n---\n\nBody.\n")) {
		t.Error("draft: true must be a draft")
	}
	if isDraft(testArticle(t, b, "---\ntitle: P\ndraft: false\n---\n\nBody.\n")) {
		t.Error("draft: false is published")
	}
	if isDraft(testArticle(t, b, "---\ntitle: P\n---\n\nBody.\n")) {
		t.Error("absent key is published")
	}
	if isDraft(testArticle(t, b, "---\ndraft: [broken\n---\n\nBody.\n")) {
		t.Error("broken front matter stays in scope")
	}
}
