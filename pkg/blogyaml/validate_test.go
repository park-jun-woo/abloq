//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what Validate가 룰 6종을 정해진 순서로 실행해 진단을 모으는지 검증
package blogyaml

import (
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {
	b := &Blog{
		Geo: Geo{Crawlers: map[string]string{"gptbot": "deny"}},
	}
	diags := Validate("blog.yaml", b, lineIndex{})
	var rules []string
	for _, d := range diags {
		rules = append(rules, d.Rule)
	}
	want := []string{"lang-bcp47", "sections-empty", "threshold-range", "threshold-range", "baseurl-format", "crawlers-policy"}
	if !reflect.DeepEqual(rules, want) {
		t.Errorf("want rules %v in order, got %v", want, rules)
	}

	ok := defaultBlog()
	ok.Site.BaseURL = "https://example.com"
	ok.Languages = []string{"ko"}
	ok.Sections = []string{"tech"}
	if diags := Validate("blog.yaml", &ok, lineIndex{}); len(diags) != 0 {
		t.Errorf("want 0 diagnostics for valid blog, got %v", diags)
	}
}
