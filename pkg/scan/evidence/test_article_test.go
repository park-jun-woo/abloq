//ff:func feature=scan type=parser control=sequence topic=evidence
//ff:what 원문 문자열 1개를 게이트 Article(ko/tech, 저장소 상대 Path)로 조립 — 주장·인용 단위 테스트 공용
package evidence

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/gate"
)

func testArticle(t *testing.T, b *blogyaml.Blog, content string) *gate.Article {
	t.Helper()
	return &gate.Article{
		Lang: "ko", Section: "tech", Slug: "fixture",
		Path: "content/ko/tech/fixture.md",
		Doc:  gate.ParseArticle(b, "ko", content),
	}
}
