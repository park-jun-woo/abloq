//ff:func feature=sitesyaml type=parser control=sequence
//ff:what indexNode가 document/mapping/sequence를 분기 처리하고 스칼라는 무시하는지 검증
package sitesyaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestIndexNode(t *testing.T) {
	var doc yaml.Node
	if err := yaml.Unmarshal([]byte("sites:\n  - name: a\n"), &doc); err != nil {
		t.Fatalf("yaml.Unmarshal: %v", err)
	}
	idx := lineIndex{}
	indexNode(idx, "", &doc)
	if idx["sites"] != 1 || idx["sites[0]"] != 2 || idx["sites[0].name"] != 2 {
		t.Errorf("document dispatch missing entries: %v", idx)
	}

	scalar := &yaml.Node{Kind: yaml.ScalarNode, Value: "x", Line: 9}
	idx = lineIndex{}
	indexNode(idx, "k", scalar)
	if len(idx) != 0 {
		t.Errorf("scalar node must add no entry, got %v", idx)
	}

	empty := &yaml.Node{Kind: yaml.DocumentNode}
	indexNode(idx, "", empty)
	if len(idx) != 0 {
		t.Errorf("empty document must add no entry, got %v", idx)
	}
}
