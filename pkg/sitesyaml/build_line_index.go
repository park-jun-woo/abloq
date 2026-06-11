//ff:func feature=sitesyaml type=parser control=sequence
//ff:what 파싱된 yaml.Node 트리에서 키 경로→라인 인덱스를 생성
package sitesyaml

import "gopkg.in/yaml.v3"

// buildLineIndex walks the YAML document and records the source line of every key path.
func buildLineIndex(doc *yaml.Node) lineIndex {
	idx := lineIndex{}
	indexNode(idx, "", doc)
	return idx
}
