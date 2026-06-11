//ff:func feature=sitesyaml type=parser control=iteration dimension=1
//ff:what buildLineIndex가 sites 리스트 항목·중첩 키별 소스 라인을 기록하는지 검증
package sitesyaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestBuildLineIndex(t *testing.T) {
	src := "sites:\n  - name: a\n    repo_path: /blogs/a\n    gsc:\n      site_url: sc-domain:a.com\n"
	var doc yaml.Node
	if err := yaml.Unmarshal([]byte(src), &doc); err != nil {
		t.Fatalf("yaml.Unmarshal: %v", err)
	}
	idx := buildLineIndex(&doc)
	want := map[string]int{
		"sites":                 1,
		"sites[0]":              2,
		"sites[0].name":         2,
		"sites[0].repo_path":    3,
		"sites[0].gsc":          4,
		"sites[0].gsc.site_url": 5,
	}
	for path, line := range want {
		if got, ok := idx[path]; !ok || got != line {
			t.Errorf("idx[%q]: want %d, got %d (present=%v)", path, line, got, ok)
		}
	}
}
