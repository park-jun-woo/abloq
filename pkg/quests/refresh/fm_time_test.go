//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what fmTime이 front matter 날짜 키를 해석하고 nil Doc·front matter 부재·키 부재는 실패인지 검증
package refresh

import (
	"testing"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestFMTime(t *testing.T) {
	b := &blogyaml.Blog{Languages: []string{"en"}}
	d := agate.ParseArticle(b, "en", "---\ntitle: T\nlastmod: 2026-06-09\n---\nbody\n")
	ts, ok := fmTime(d, "lastmod")
	if !ok || ts.Format("2006-01-02") != "2026-06-09" {
		t.Errorf("fmTime = %v %v", ts, ok)
	}
	if _, ok := fmTime(d, "date"); ok {
		t.Error("absent key must fail")
	}
	if _, ok := fmTime(nil, "lastmod"); ok {
		t.Error("nil doc must fail")
	}
	noFM := agate.ParseArticle(b, "en", "just a body\n")
	if _, ok := fmTime(noFM, "lastmod"); ok {
		t.Error("missing front matter must fail")
	}
}
