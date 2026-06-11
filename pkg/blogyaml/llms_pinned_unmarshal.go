//ff:func feature=blogyaml type=parser control=iteration dimension=1
//ff:what pinned 엔트리 1개를 strict 디코드 — 매핑 필수, 키별 에러(미지 키 포함)를 전부 모아 TypeError로 반환
package blogyaml

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// UnmarshalYAML decodes one pinned entry strictly: unknown keys are rejected
// with the same message shape as KnownFields(true).
func (p *LlmsPinned) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return &yaml.TypeError{Errors: []string{fmt.Sprintf(
			"line %d: cannot unmarshal geo.llms_txt.pinned entry: want a mapping", value.Line)}}
	}
	var errs []string
	for i := 0; i+1 < len(value.Content); i += 2 {
		errs = append(errs, llmsDecodeErrors(decodeLlmsPinnedKey(p, value.Content[i], value.Content[i+1]))...)
	}
	if len(errs) > 0 {
		return &yaml.TypeError{Errors: errs}
	}
	return nil
}
