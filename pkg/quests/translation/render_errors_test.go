//ff:func feature=quest type=generator control=sequence
//ff:what Render 에러 경로 검증 — Payload 디코드 불량, blog.yaml 부재/진단, 원문 파일 부재
package translation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRenderErrors(t *testing.T) {
	if _, err := (Definition{}).Render(nil, &quest.Item{Payload: []byte("broken")}); err == nil {
		t.Error("broken payload: want decode error")
	}
	root := writeInstance(t)
	origin, _ := passPair()
	originPath := writeFile(t, root, "content/en/posts/fixture.md", origin)
	items, err := Definition{}.Seed([]string{originPath})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	it := items[0]
	if err := os.Remove(originPath); err != nil {
		t.Fatal(err)
	}
	if _, err := (Definition{}).Render(nil, it); err == nil {
		t.Error("origin removed: want read error")
	}
	bad := []byte("site:\n  baseURL: not-a-url\n  title: T\n  author: A\n\nlanguages: [en, ko]\nsections: [posts]\n")
	if err := os.WriteFile(filepath.Join(root, "blog.yaml"), bad, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := (Definition{}).Render(nil, it); err == nil {
		t.Error("blog.yaml diagnostics: want error")
	}
	if err := os.Remove(filepath.Join(root, "blog.yaml")); err != nil {
		t.Fatal(err)
	}
	if _, err := (Definition{}).Render(nil, it); err == nil {
		t.Error("blog.yaml absent: want load error")
	}
}
