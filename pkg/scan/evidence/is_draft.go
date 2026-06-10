//ff:func feature=scan type=parser control=sequence topic=evidence
//ff:what 글 1편의 draft 여부 — front matter의 draft: true만 참, 파싱 불가·키 없음은 발행으로 간주
package evidence

import (
	"gopkg.in/yaml.v3"

	"github.com/park-jun-woo/abloq/pkg/gate"
)

// isDraft reads the article's draft flag. Only an explicit draft: true hides
// an article from the scanner; broken front matter stays in scope — its
// decay is still the corpus' decay.
func isDraft(a *gate.Article) bool {
	var fm struct {
		Draft bool `yaml:"draft"`
	}
	if err := yaml.Unmarshal([]byte(a.Doc.FrontMatter), &fm); err != nil {
		return false
	}
	return fm.Draft
}
