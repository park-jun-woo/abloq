//ff:func feature=blogyaml type=parser control=iteration dimension=1
//ff:what indexSequence가 항목을 순회하며 "prefix[i]" 경로에 항목 라인을 기록하고 중첩 노드로 재귀하는지 검증
package blogyaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestIndexSequence(t *testing.T) {
	src := "- x\n- y: 2\n"
	var doc yaml.Node
	if err := yaml.Unmarshal([]byte(src), &doc); err != nil {
		t.Fatalf("yaml.Unmarshal: %v", err)
	}
	idx := lineIndex{}
	indexSequence(idx, "items", doc.Content[0])
	want := map[string]int{"items[0]": 1, "items[1]": 2, "items[1].y": 2}
	if len(idx) != len(want) {
		t.Errorf("want %d entries, got %d: %v", len(want), len(idx), idx)
	}
	for path, line := range want {
		if idx[path] != line {
			t.Errorf("idx[%q]: want %d, got %d", path, line, idx[path])
		}
	}
}
