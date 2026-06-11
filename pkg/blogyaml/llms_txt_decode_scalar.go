//ff:func feature=blogyaml type=parser control=sequence
//ff:what geo.llms_txt 문자열 단축형 디코드 — 스칼라 값을 mode로 넣고 나머지 필드는 기본값 유지 (null이면 전체 기본값)
package blogyaml

import "gopkg.in/yaml.v3"

// decodeLlmsTxtScalar maps the string shorthand onto Mode. A null scalar
// (bare "llms_txt:") keeps every default.
func decodeLlmsTxtScalar(s *LlmsTxtSpec, value *yaml.Node) error {
	if value.Tag == "!!null" {
		return nil
	}
	s.Mode = value.Value
	return nil
}
