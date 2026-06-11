//ff:func feature=blogyaml type=parser control=selection
//ff:what geo.llms_txt.languages 2차 union 디코드 — base/all 스칼라는 1원소 리스트로, 시퀀스는 그대로 정규화
package blogyaml

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// decodeLlmsLanguages normalizes the languages union: a scalar ("base", "all"
// or one language code) becomes a one-element list, a sequence decodes as-is.
func decodeLlmsLanguages(s *LlmsTxtSpec, val *yaml.Node) error {
	switch val.Kind {
	case yaml.SequenceNode:
		return val.Decode(&s.Languages)
	case yaml.ScalarNode:
		return decodeLlmsLanguagesScalar(s, val)
	}
	return &yaml.TypeError{Errors: []string{fmt.Sprintf(
		"line %d: cannot unmarshal geo.llms_txt.languages: want base, all, or a sequence", val.Line)}}
}
