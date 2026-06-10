//ff:func feature=blogyaml type=parser control=iteration dimension=1
//ff:what indexNode가 노드 종류(document/빈 document/mapping/sequence/scalar)별로 올바르게 분기하는지 검증
package blogyaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestIndexNode(t *testing.T) {
	parse := func(t *testing.T, src string) *yaml.Node {
		t.Helper()
		var doc yaml.Node
		if err := yaml.Unmarshal([]byte(src), &doc); err != nil {
			t.Fatalf("yaml.Unmarshal: %v", err)
		}
		return &doc
	}
	cases := []struct {
		name string
		node func(t *testing.T) *yaml.Node
		want map[string]int
	}{
		{
			name: "document",
			node: func(t *testing.T) *yaml.Node { return parse(t, "a: 1\n") },
			want: map[string]int{"a": 1},
		},
		{
			name: "empty document",
			node: func(t *testing.T) *yaml.Node { return &yaml.Node{Kind: yaml.DocumentNode} },
			want: map[string]int{},
		},
		{
			name: "mapping",
			node: func(t *testing.T) *yaml.Node { return parse(t, "a: 1\nb: 2\n").Content[0] },
			want: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "sequence",
			node: func(t *testing.T) *yaml.Node { return parse(t, "- x\n- y\n").Content[0] },
			want: map[string]int{"[0]": 1, "[1]": 2},
		},
		{
			name: "scalar",
			node: func(t *testing.T) *yaml.Node { return parse(t, "just-a-scalar\n").Content[0] },
			want: map[string]int{},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkIndexNode(t, tc.node(t), tc.want) })
	}
}
