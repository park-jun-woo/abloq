//ff:func feature=quest type=parser control=sequence
//ff:what insight.yaml 경로 1개 → Item — 인스턴스 루트 탐색, 사이드카 규약 역산으로 대상 글 경로 파생, 명세 검증, Payload 채움
package writing

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/insight"
)

// seedItem turns one insight.yaml path into a quest item. The instance root
// is the nearest ancestor directory holding blog.yaml; the target article
// path is derived from the Phase015 sidecar convention (insight.PathFor
// inverted); the key is lang/section/slug from that path.
func seedItem(arg string) (*quest.Item, error) {
	abs, err := filepath.Abs(arg)
	if err != nil {
		return nil, err
	}
	root, err := findRoot(filepath.Dir(abs))
	if err != nil {
		return nil, err
	}
	rel, err := filepath.Rel(root, abs)
	if err != nil {
		return nil, err
	}
	article, ok := articleFor(filepath.ToSlash(rel))
	if !ok {
		return nil, fmt.Errorf("%s: not an insight.yaml sidecar path (want insight.yaml or *.insight.yaml)", arg)
	}
	lang, section, slug, ok := keyParts(article)
	if !ok {
		return nil, fmt.Errorf("%s: target %s is not under content/<lang>/<section>/", arg, article)
	}
	_, errs, _, err := insight.Load(abs)
	if err != nil {
		return nil, err
	}
	if len(errs) > 0 {
		return nil, fmt.Errorf("%s", errs[0].String())
	}
	it := &quest.Item{Key: lang + "/" + section + "/" + slug, State: quest.TODO}
	p := Payload{Root: root, Insight: filepath.ToSlash(rel), Article: article,
		Lang: lang, Section: section, Slug: slug}
	if err := it.SetPayload(p); err != nil {
		return nil, err
	}
	return it, nil
}
