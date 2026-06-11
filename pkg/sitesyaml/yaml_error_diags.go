//ff:func feature=sitesyaml type=parser control=iteration dimension=1 topic=diagnostics
//ff:what yaml.v3 에러(단일 또는 TypeError 다건)를 진단 목록으로 변환
package sitesyaml

import (
	"errors"

	"gopkg.in/yaml.v3"
)

// yamlErrorDiags converts a yaml.v3 decode error into one diagnostic per message.
func yamlErrorDiags(filename string, err error) []Diagnostic {
	msgs := []string{err.Error()}
	var te *yaml.TypeError
	if errors.As(err, &te) {
		msgs = te.Errors
	}
	diags := make([]Diagnostic, 0, len(msgs))
	for _, msg := range msgs {
		diags = append(diags, yamlErrorDiag(filename, msg))
	}
	return diags
}
