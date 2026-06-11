//ff:func feature=gen type=generator control=sequence
//ff:what scopeLangs가 미지정/base→기본 언어 1개, all→전 언어, 명시 목록→그대로를 반환하는지 검증
package llms

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestScopeLangs(t *testing.T) {
	blogWith := func(spec []string) *blogyaml.Blog {
		return &blogyaml.Blog{
			Languages: []string{"en", "ko", "ja"},
			Geo:       blogyaml.Geo{LlmsTxt: blogyaml.LlmsTxtSpec{Languages: spec}},
		}
	}
	if got := scopeLangs(blogWith(nil)); !reflect.DeepEqual(got, []string{"en"}) {
		t.Errorf("unset scope = %v, want [en] (base default)", got)
	}
	if got := scopeLangs(blogWith([]string{"base"})); !reflect.DeepEqual(got, []string{"en"}) {
		t.Errorf("base scope = %v, want [en]", got)
	}
	if got := scopeLangs(blogWith([]string{"all"})); !reflect.DeepEqual(got, []string{"en", "ko", "ja"}) {
		t.Errorf("all scope = %v, want every declared language", got)
	}
	if got := scopeLangs(blogWith([]string{"ko", "ja"})); !reflect.DeepEqual(got, []string{"ko", "ja"}) {
		t.Errorf("explicit scope = %v, want [ko ja]", got)
	}
	empty := &blogyaml.Blog{}
	if got := scopeLangs(empty); len(got) != 0 {
		t.Errorf("no declared languages must yield empty scope, got %v", got)
	}
}
