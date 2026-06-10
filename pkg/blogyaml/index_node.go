//ff:func feature=blogyaml type=parser control=selection
//ff:what yaml.Node 종류(document/mapping/sequence)별로 라인 인덱싱을 분기
package blogyaml

import "gopkg.in/yaml.v3"

// indexNode dispatches line indexing by node kind. Scalar nodes need no entry of their own.
func indexNode(idx lineIndex, prefix string, n *yaml.Node) {
	switch n.Kind {
	case yaml.DocumentNode:
		if len(n.Content) > 0 {
			indexNode(idx, prefix, n.Content[0])
		}
	case yaml.MappingNode:
		indexMapping(idx, prefix, n)
	case yaml.SequenceNode:
		indexSequence(idx, prefix, n)
	}
}
