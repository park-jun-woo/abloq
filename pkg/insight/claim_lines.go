//ff:func feature=insight type=parser control=iteration dimension=1
//ff:what insight.yaml 루트 매핑에서 claims 시퀀스 항목별 라인 번호를 수집 — 진단 위치용, 실패 시 nil
package insight

import "gopkg.in/yaml.v3"

// claimLines returns the source line of each claims[] item, for diagnostics.
func claimLines(data []byte) []int {
	var doc yaml.Node
	if yaml.Unmarshal(data, &doc) != nil || len(doc.Content) == 0 {
		return nil
	}
	m := doc.Content[0]
	if m.Kind != yaml.MappingNode {
		return nil
	}
	var lines []int
	for i := 0; i+1 < len(m.Content); i += 2 {
		if m.Content[i].Value != "claims" || m.Content[i+1].Kind != yaml.SequenceNode {
			continue
		}
		for _, item := range m.Content[i+1].Content {
			lines = append(lines, item.Line)
		}
	}
	return lines
}
