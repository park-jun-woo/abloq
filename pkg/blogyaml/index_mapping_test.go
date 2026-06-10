//ff:func feature=blogyaml type=parser control=iteration dimension=1
//ff:what indexMapping이 키-값 쌍을 순회하며 "prefix.key" 경로에 키 라인을 기록하는지 검증
package blogyaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestIndexMapping(t *testing.T) {
	src := "a: 1\nb:\n  c: 2\n"
	var doc yaml.Node
	if err := yaml.Unmarshal([]byte(src), &doc); err != nil {
		t.Fatalf("yaml.Unmarshal: %v", err)
	}
	idx := lineIndex{}
	indexMapping(idx, "root", doc.Content[0])
	want := map[string]int{"root.a": 1, "root.b": 2, "root.b.c": 3}
	if len(idx) != len(want) {
		t.Errorf("want %d entries, got %d: %v", len(want), len(idx), idx)
	}
	for path, line := range want {
		if idx[path] != line {
			t.Errorf("idx[%q]: want %d, got %d", path, line, idx[path])
		}
	}
}
