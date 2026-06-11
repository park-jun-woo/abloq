//ff:func feature=quest type=parser control=sequence
//ff:what 원문 글 경로 1개 → seedSrc — 인스턴스 루트 탐색, Key 부품 파생, 기본 언어 검증, blog.yaml 언어 목록과 원문 lastmod 적재
package translation

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	agate "github.com/park-jun-woo/abloq/pkg/gate"
	"github.com/park-jun-woo/abloq/pkg/quests/common"
)

// seedOrigin resolves one origin article argument: the instance root is the
// nearest ancestor holding blog.yaml; the path must be a content article in
// the blog's default language (translations are seeded FROM the original);
// the origin's lastmod is captured for the staleness comparison.
func seedOrigin(arg string) (seedSrc, error) {
	abs, err := filepath.Abs(arg)
	if err != nil {
		return seedSrc{}, err
	}
	root, err := common.FindRoot(filepath.Dir(abs))
	if err != nil {
		return seedSrc{}, err
	}
	rel, err := filepath.Rel(root, abs)
	if err != nil {
		return seedSrc{}, err
	}
	origin := filepath.ToSlash(rel)
	lang, section, slug, ok := common.KeyParts(origin)
	if !ok {
		return seedSrc{}, fmt.Errorf("%s: not under content/<lang>/<section>/", arg)
	}
	b, diags, err := blogyaml.Load(filepath.Join(root, "blog.yaml"))
	if err != nil {
		return seedSrc{}, err
	}
	if len(diags) > 0 {
		return seedSrc{}, fmt.Errorf("blog.yaml: %s", diags[0].String())
	}
	if len(b.Languages) < 2 {
		return seedSrc{}, fmt.Errorf("blog.yaml declares no translation languages (languages: %v)", b.Languages)
	}
	if lang != b.Languages[0] {
		return seedSrc{}, fmt.Errorf("%s: lang %s is not the default language %s — seed the original, not a translation", arg, lang, b.Languages[0])
	}
	body, err := os.ReadFile(abs)
	if err != nil {
		return seedSrc{}, fmt.Errorf("origin article unreadable: %w", err)
	}
	doc := agate.ParseArticle(b, lang, string(body))
	lastmod, has := fmTime(doc, "lastmod")
	return seedSrc{root: root, origin: origin, originLang: lang, section: section,
		slug: slug, blog: b, lastmod: lastmod, hasLastmod: has}, nil
}
