//ff:func feature=quest type=parser control=sequence
//ff:what 원문(기본 언어) 글을 읽고 파싱해 비교용 Article로 적재 — Prepare가 매 제출 호출, 읽기 실패는 에러
package translation

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// loadOrigin reads and parses the default-language origin article the parity
// rule compares against. It is loaded fresh on every submit, so editing the
// origin to dodge the gate just changes what the gate compares to.
func loadOrigin(b *blogyaml.Blog, p Payload) (*agate.Article, error) {
	body, err := os.ReadFile(filepath.Join(p.Root, filepath.FromSlash(p.Origin)))
	if err != nil {
		return nil, fmt.Errorf("origin article unreadable: %w", err)
	}
	return &agate.Article{Lang: p.OriginLang, Section: p.Section, Slug: p.Slug,
		Path: p.Origin, Doc: agate.ParseArticle(b, p.OriginLang, string(body))}, nil
}
