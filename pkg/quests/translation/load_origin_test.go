//ff:func feature=quest type=parser control=sequence
//ff:what loadOrigin 검증 — 원문 읽기·파싱과 Article 메타 채움, 파일 부재는 에러
package translation

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestLoadOrigin(t *testing.T) {
	root := writeInstance(t)
	origin, _ := passPair()
	writeFile(t, root, "content/en/posts/fixture.md", origin)
	b, _, err := blogyaml.Load(root + "/blog.yaml")
	if err != nil {
		t.Fatalf("blog.yaml: %v", err)
	}
	p := Payload{Root: root, Origin: "content/en/posts/fixture.md",
		OriginLang: "en", Section: "posts", Slug: "fixture"}
	art, err := loadOrigin(b, p)
	if err != nil {
		t.Fatalf("loadOrigin: %v", err)
	}
	if art.Lang != "en" || art.Path != p.Origin || len(art.Doc.Sections) != 1 {
		t.Errorf("art = %+v", art)
	}
	p.Origin = "content/en/posts/missing.md"
	if _, err := loadOrigin(b, p); err == nil {
		t.Error("missing origin: want error")
	}
}
