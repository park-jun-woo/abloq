//ff:func feature=gate type=parser control=sequence
//ff:what 픽스처 마크다운 1개를 파싱해 게이트 Article로 조립 (Base 없음)
package gate

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func artFromMD(t *testing.T, b *blogyaml.Blog, lang, section, slug, file string) *Article {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", file))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	return &Article{
		Lang: lang, Section: section, Slug: slug,
		Path: filepath.Join("content", lang, section, slug+".md"),
		Doc:  ParseArticle(b, lang, string(data)),
	}
}
