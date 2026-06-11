//ff:func feature=blogyaml type=parser control=sequence
//ff:what geo.llms_txt.languages 스칼라 디코드 — 값 1개를 1원소 리스트로 정규화, null이면 기본값(base) 유지
package blogyaml

import "gopkg.in/yaml.v3"

// decodeLlmsLanguagesScalar wraps a scalar languages value into the
// normalized one-element list. A null scalar keeps the default scope.
func decodeLlmsLanguagesScalar(s *LlmsTxtSpec, val *yaml.Node) error {
	if val.Tag == "!!null" {
		return nil
	}
	s.Languages = []string{val.Value}
	return nil
}
