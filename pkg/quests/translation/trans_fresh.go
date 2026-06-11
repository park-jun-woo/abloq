//ff:func feature=quest type=parser control=sequence
//ff:what 기존 번역의 최신성 판정 — 번역 파일이 존재하고 lastmod가 원문 lastmod 이상이면 true(아이템 미생성)
package translation

import (
	"os"
	"path/filepath"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// transFresh reports whether the target-language translation already exists
// and is at least as recent as the origin (lastmod comparison). Missing
// files, unparseable lastmods or an unparseable origin lastmod all mean
// "stale": the language gets an item.
func transFresh(src seedSrc, lang, article string) bool {
	if !src.hasLastmod {
		return false
	}
	body, err := os.ReadFile(filepath.Join(src.root, filepath.FromSlash(article)))
	if err != nil {
		return false
	}
	doc := agate.ParseArticle(src.blog, lang, string(body))
	lastmod, ok := fmTime(doc, "lastmod")
	if !ok {
		return false
	}
	return !lastmod.Before(src.lastmod)
}
