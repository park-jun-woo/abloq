//ff:func feature=scan type=parser control=sequence topic=cluster
//ff:what 링크 대상 1개 → 기본 언어 글 키 — baseURL/사이트 절대경로만, 언어 세그먼트는 Blog.URLLang 규칙, /<section>/<slug>/ 형태만 해석
//ff:why 번역 URL(언어 세그먼트 보유)은 그래프 밖이다 — 그래프는 기본 언어 1회 구성이고 태그·섹션 페이지 등 비글 경로는 세그먼트 수·섹션 대조로 걸러진다
package cluster

import (
	"slices"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// resolveTarget maps one markdown link target onto a default-language
// article key (<section>/<slug>). Site-absolute paths and absolute URLs
// under baseURL qualify; the language segment follows the Blog.URLLang
// contract (the root-served default language has none, so any declared
// language segment marks a translation URL — out of the graph).
func resolveTarget(b *blogyaml.Blog, lang, target string) (string, bool) {
	if i := strings.IndexAny(target, "#?"); i >= 0 {
		target = target[:i]
	}
	base := strings.TrimRight(b.Site.BaseURL, "/")
	if base != "" && strings.HasPrefix(target, base) {
		target = target[len(base):]
	}
	if !strings.HasPrefix(target, "/") {
		return "", false
	}
	segs := strings.Split(strings.Trim(target, "/"), "/")
	if seg := b.URLLang(lang); seg != "" {
		if len(segs) == 0 || segs[0] != seg {
			return "", false
		}
		segs = segs[1:]
	} else if len(segs) > 0 && slices.Contains(b.Languages, segs[0]) {
		return "", false
	}
	if len(segs) != 2 || !slices.Contains(b.Sections, segs[0]) || segs[1] == "" {
		return "", false
	}
	return PostKey(segs[0], segs[1]), true
}
