//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what 원문 문자열 1개를 게이트 Article로 조립 (en/tech, Base 없음) — 인용·주장 테스트 공용
package gate

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func artFromContent(t *testing.T, b *blogyaml.Blog, content string) *Article {
	t.Helper()
	return &Article{
		Lang: "en", Section: "tech", Slug: "fixture",
		Path: "content/en/tech/fixture.md",
		Doc:  ParseArticle(b, "en", content),
	}
}
