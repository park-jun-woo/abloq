//ff:func feature=sitesyaml type=parser control=iteration dimension=1
//ff:what indexMapping이 키 라인을 기록하고 값 노드로 재귀하는지 검증
package sitesyaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestIndexMapping(t *testing.T) {
	var doc yaml.Node
	if err := yaml.Unmarshal([]byte("name: a\ngsc:\n  site_url: u\n"), &doc); err != nil {
		t.Fatalf("yaml.Unmarshal: %v", err)
	}
	idx := lineIndex{}
	indexMapping(idx, "sites[0]", doc.Content[0])
	want := map[string]int{
		"sites[0].name":         1,
		"sites[0].gsc":          2,
		"sites[0].gsc.site_url": 3,
	}
	for path, line := range want {
		if got := idx[path]; got != line {
			t.Errorf("idx[%q] = %d, want %d", path, got, line)
		}
	}
}
