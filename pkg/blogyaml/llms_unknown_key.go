//ff:func feature=blogyaml type=parser control=sequence topic=diagnostics
//ff:what 객체 폼 미지 키를 KnownFields(true)와 동일한 메시지 형식의 TypeError로 변환 — yaml_error_diag가 unknown-key 룰·라인으로 분류
package blogyaml

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// llmsUnknownKey builds the strict-mode-equivalent unknown-key error so the
// existing diagnostic classifier yields the same rule ID and line.
func llmsUnknownKey(key *yaml.Node, typeName string) error {
	return &yaml.TypeError{Errors: []string{fmt.Sprintf(
		"line %d: field %s not found in type %s", key.Line, key.Value, typeName)}}
}
