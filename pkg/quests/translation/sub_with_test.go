//ff:func feature=quest type=parser control=sequence
//ff:what 테스트 헬퍼 — 원문(en)·번역(ko) 원문 문자열로 Submission 조립 (Base nil, 인스턴스 blog.yaml 로드)
package translation

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

func subWith(t *testing.T, root, originMD, koMD string) *Submission {
	t.Helper()
	b, diags, err := blogyaml.Load(root + "/blog.yaml")
	if err != nil {
		t.Fatalf("load blog.yaml: %v", err)
	}
	if len(diags) > 0 {
		t.Fatalf("fixture blog.yaml diagnostics: %v", diags)
	}
	origin := &agate.Article{Lang: "en", Section: "posts", Slug: "fixture",
		Path: "content/en/posts/fixture.md", Doc: agate.ParseArticle(b, "en", originMD)}
	trans := &agate.Article{Lang: "ko", Section: "posts", Slug: "fixture",
		Path: "content/ko/posts/fixture.md", Doc: agate.ParseArticle(b, "ko", koMD)}
	return &Submission{
		Target:  agate.NewTarget(root, b, []*agate.Article{trans}),
		Origin:  origin,
		Article: trans.Path,
		Root:    root, Lang: "ko", OriginLang: "en",
	}
}
