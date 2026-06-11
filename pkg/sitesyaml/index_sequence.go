//ff:func feature=sitesyaml type=parser control=iteration dimension=1
//ff:what 시퀀스 노드의 항목을 순회하며 "prefix[i]" 경로에 항목 라인을 기록
package sitesyaml

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// indexSequence records each item's line and recurses into nested nodes.
func indexSequence(idx lineIndex, prefix string, n *yaml.Node) {
	for i, item := range n.Content {
		path := fmt.Sprintf("%s[%d]", prefix, i)
		idx[path] = item.Line
		indexNode(idx, path, item)
	}
}
