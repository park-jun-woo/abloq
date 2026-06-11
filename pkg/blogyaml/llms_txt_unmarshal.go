//ff:func feature=blogyaml type=parser control=selection
//ff:what geo.llms_txt union 디코드 진입점 — 스칼라(단축형 mode)·매핑(객체 폼)을 분기, 그 외 노드는 yaml-syntax 에러
//ff:why parse.go의 KnownFields(true)는 custom unmarshaler 내부에 미적용 — union 분기와 미지 키 거부를 여기서 직접 구현 (Phase021)
package blogyaml

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// UnmarshalYAML decodes the geo.llms_txt union: a string shorthand sets Mode
// only (defaults stay), a mapping decodes the object form strictly.
func (s *LlmsTxtSpec) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		return decodeLlmsTxtScalar(s, value)
	case yaml.MappingNode:
		return decodeLlmsTxtMap(s, value)
	}
	return &yaml.TypeError{Errors: []string{fmt.Sprintf(
		"line %d: cannot unmarshal geo.llms_txt: want a string (auto|manual|off) or a mapping", value.Line)}}
}
