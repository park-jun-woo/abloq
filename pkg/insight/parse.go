//ff:func feature=insight type=parser control=sequence
//ff:what insight.yaml 바이트를 strict 디코드(unknown key = 에러) — 실패 시 yaml-syntax/unknown-key 진단
package insight

import (
	"bytes"
	"errors"
	"io"

	"gopkg.in/yaml.v3"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// Parse decodes insight.yaml strictly (unknown keys are errors).
// On failure it returns a nil Insight with diagnostics.
func Parse(filename string, data []byte) (*Insight, []blogyaml.Diagnostic) {
	var ins Insight
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	err := dec.Decode(&ins)
	if errors.Is(err, io.EOF) {
		return nil, []blogyaml.Diagnostic{{File: filename, Line: 1, Rule: "yaml-syntax", Message: "insight.yaml is empty"}}
	}
	if err != nil {
		return nil, []blogyaml.Diagnostic{{File: filename, Line: 1, Rule: "yaml-syntax", Message: err.Error()}}
	}
	return &ins, nil
}
