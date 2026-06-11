//ff:func feature=blogyaml type=parser control=selection
//ff:what geo.llms_txt 객체 폼의 키 1개를 디코드 — 알려진 6키만 허용, 그 외는 unknown-key(strict-동등) 에러
package blogyaml

import "gopkg.in/yaml.v3"

// decodeLlmsTxtKey decodes one known object-form key into the spec; unknown
// keys are rejected with the same message shape as KnownFields(true).
func decodeLlmsTxtKey(s *LlmsTxtSpec, key, val *yaml.Node) error {
	switch key.Value {
	case "mode":
		return val.Decode(&s.Mode)
	case "languages":
		return decodeLlmsLanguages(s, val)
	case "header":
		return val.Decode(&s.Header)
	case "pinned":
		return val.Decode(&s.Pinned)
	case "section_labels":
		return val.Decode(&s.SectionLabels)
	case "max_summary":
		return val.Decode(&s.MaxSummary)
	}
	return llmsUnknownKey(key, "blogyaml.LlmsTxtSpec")
}
