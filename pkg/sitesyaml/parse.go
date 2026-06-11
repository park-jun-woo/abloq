//ff:func feature=sitesyaml type=parser control=sequence
//ff:what sites.yaml 바이트를 strict 디코드(unknown key = 에러) + active 기본값 주입 + 라인 인덱스 생성
package sitesyaml

import (
	"bytes"
	"errors"
	"io"

	"gopkg.in/yaml.v3"
)

// Parse decodes sites.yaml strictly over schema v1.
// On parse failure it returns nil Sites with yaml-syntax / unknown-key diagnostics.
func Parse(filename string, data []byte) (*Sites, lineIndex, []Diagnostic) {
	var doc yaml.Node
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return nil, nil, yamlErrorDiags(filename, err)
	}
	idx := buildLineIndex(&doc)

	var s Sites
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	err := dec.Decode(&s)
	if errors.Is(err, io.EOF) {
		return nil, idx, []Diagnostic{{File: filename, Line: 1, Rule: "yaml-syntax", Message: "sites.yaml is empty"}}
	}
	if err != nil {
		return nil, idx, yamlErrorDiags(filename, err)
	}
	defaultActive(&s, idx)
	return &s, idx, nil
}
