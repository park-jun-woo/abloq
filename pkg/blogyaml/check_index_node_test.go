//ff:func feature=blogyaml type=parser control=iteration dimension=1
//ff:what indexNode 케이스 하나를 실행해 라인 인덱스 엔트리 수와 각 경로의 라인을 검증
package blogyaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func checkIndexNode(t *testing.T, node *yaml.Node, want map[string]int) {
	t.Helper()
	idx := lineIndex{}
	indexNode(idx, "", node)
	if len(idx) != len(want) {
		t.Errorf("want %d entries, got %d: %v", len(want), len(idx), idx)
	}
	for path, line := range want {
		if idx[path] != line {
			t.Errorf("idx[%q]: want %d, got %d", path, line, idx[path])
		}
	}
}
