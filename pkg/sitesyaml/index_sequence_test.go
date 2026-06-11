//ff:func feature=sitesyaml type=parser control=iteration dimension=1
//ff:what indexSequence가 항목 라인("prefix[i]")을 기록하고 중첩 노드로 재귀하는지 검증
package sitesyaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestIndexSequence(t *testing.T) {
	var doc yaml.Node
	if err := yaml.Unmarshal([]byte("- name: a\n- name: b\n"), &doc); err != nil {
		t.Fatalf("yaml.Unmarshal: %v", err)
	}
	idx := lineIndex{}
	indexSequence(idx, "sites", doc.Content[0])
	want := map[string]int{
		"sites[0]":      1,
		"sites[0].name": 1,
		"sites[1]":      2,
		"sites[1].name": 2,
	}
	for path, line := range want {
		if got := idx[path]; got != line {
			t.Errorf("idx[%q] = %d, want %d", path, got, line)
		}
	}
}
