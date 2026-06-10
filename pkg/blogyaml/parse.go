//ff:func feature=blogyaml type=parser control=sequence
//ff:what blog.yaml 바이트를 strict 디코드(unknown key = 에러) + 기본값 주입 + 라인 인덱스 생성
package blogyaml

import (
	"bytes"
	"errors"
	"io"

	"gopkg.in/yaml.v3"
)

// Parse decodes blog.yaml strictly over schema v1 defaults.
// On parse failure it returns nil Blog with yaml-syntax / unknown-key diagnostics.
func Parse(filename string, data []byte) (*Blog, lineIndex, []Diagnostic) {
	var doc yaml.Node
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return nil, nil, yamlErrorDiags(filename, err)
	}
	idx := buildLineIndex(&doc)

	b := defaultBlog()
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	err := dec.Decode(&b)
	if errors.Is(err, io.EOF) {
		return nil, idx, []Diagnostic{{File: filename, Line: 1, Rule: "yaml-syntax", Message: "blog.yaml is empty"}}
	}
	if err != nil {
		return nil, idx, yamlErrorDiags(filename, err)
	}
	return &b, idx, nil
}
