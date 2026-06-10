//ff:func feature=blogyaml type=parser control=iteration dimension=1
//ff:what 매핑 노드의 키-값 쌍을 순회하며 "prefix.key" 경로에 키 라인을 기록
package blogyaml

import "gopkg.in/yaml.v3"

// indexMapping records each key's line and recurses into its value node.
func indexMapping(idx lineIndex, prefix string, n *yaml.Node) {
	for i := 0; i+1 < len(n.Content); i += 2 {
		key := n.Content[i]
		path := joinPath(prefix, key.Value)
		idx[path] = key.Line
		indexNode(idx, path, n.Content[i+1])
	}
}
