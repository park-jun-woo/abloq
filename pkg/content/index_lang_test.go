//ff:func feature=content type=parser control=sequence
//ff:what indexLang이 blog.yaml 선언 섹션 순서대로 순회하고(없는 섹션은 건너뜀) 언어별 발행 글만 수집하는지 검증
package content

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestIndexLang(t *testing.T) {
	root := filepath.Join("testdata", "fixture")
	b, diags, err := blogyaml.Load(filepath.Join(root, "blog.yaml"))
	if err != nil || len(diags) > 0 {
		t.Fatalf("blog.yaml load: err=%v diags=%v", err, diags)
	}
	ko := indexLang(root, b, "ko")
	if len(ko) != 2 {
		t.Fatalf("ko entries = %d, want 2 (opinion section dir is absent): %+v", len(ko), ko)
	}
	if ko[0].Slug != "post-a" || ko[1].Slug != "post-b" {
		t.Errorf("ko order = [%s %s], want declared section then name order", ko[0].Slug, ko[1].Slug)
	}
	en := indexLang(root, b, "en")
	if len(en) != 1 || en[0].Slug != "custom-en" {
		t.Errorf("en entries = %+v, want only custom-en", en)
	}
	if none := indexLang(root, b, "ja"); len(none) != 0 {
		t.Errorf("undeclared lang dir = %+v, want empty", none)
	}
}
