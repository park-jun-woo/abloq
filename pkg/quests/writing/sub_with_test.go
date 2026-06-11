//ff:func feature=quest type=parser control=sequence
//ff:what 테스트 헬퍼 — 글 원문(+선택 기준선 원문)으로 Submission 조립, baseMD 빈 문자열이면 Base nil(집필 규약)
package writing

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

func subWith(t *testing.T, root, md, baseMD string) *Submission {
	t.Helper()
	b, diags, err := blogyaml.Load(root + "/blog.yaml")
	if err != nil {
		t.Fatalf("load blog.yaml: %v", err)
	}
	if len(diags) > 0 {
		t.Fatalf("fixture blog.yaml diagnostics: %v", diags)
	}
	art := &agate.Article{Lang: "en", Section: "posts", Slug: "fixture",
		Path: "content/en/posts/fixture.md", Doc: agate.ParseArticle(b, "en", md)}
	if baseMD != "" {
		art.Base = agate.ParseArticle(b, "en", baseMD)
	}
	return &Submission{
		Target:  agate.NewTarget(root, b, []*agate.Article{art}),
		Article: art.Path,
	}
}
