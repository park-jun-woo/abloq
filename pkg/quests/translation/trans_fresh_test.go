//ff:func feature=quest type=parser control=sequence
//ff:what transFresh 검증 — 번역 부재/구판 lastmod는 stale, 동일·이후 lastmod는 fresh, 원문 lastmod 미상은 무조건 stale
package translation

import (
	"strings"
	"testing"
)

func TestTransFresh(t *testing.T) {
	root := writeInstance(t)
	origin, ko := passPair()
	originPath := writeFile(t, root, "content/en/posts/fixture.md", origin)
	src, err := seedOrigin(originPath)
	if err != nil {
		t.Fatalf("seedOrigin: %v", err)
	}
	if transFresh(src, "ko", "content/ko/posts/fixture.md") {
		t.Error("missing translation: want stale")
	}
	writeFile(t, root, "content/ko/posts/fixture.md", strings.Replace(ko, "lastmod: 2026-06-03", "lastmod: 2026-06-02", 1))
	if transFresh(src, "ko", "content/ko/posts/fixture.md") {
		t.Error("older translation: want stale")
	}
	writeFile(t, root, "content/ko/posts/fixture.md", ko)
	if !transFresh(src, "ko", "content/ko/posts/fixture.md") {
		t.Error("same-lastmod translation: want fresh")
	}
	writeFile(t, root, "content/ko/posts/fixture.md", removeLine(ko, "lastmod:"))
	if transFresh(src, "ko", "content/ko/posts/fixture.md") {
		t.Error("translation lastmod unparseable: want stale")
	}
	src.hasLastmod = false
	if transFresh(src, "ko", "content/ko/posts/fixture.md") {
		t.Error("origin lastmod unknown: want stale")
	}
}
