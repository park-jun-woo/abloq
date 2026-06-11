//ff:func feature=blogyaml type=parser control=sequence
//ff:what decodeLlmsTxtScalar가 문자열 단축형을 Mode에 넣고 null 스칼라(빈 llms_txt:)는 기본값을 유지하는지 검증
package blogyaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestDecodeLlmsTxtScalar(t *testing.T) {
	var s LlmsTxtSpec
	val := &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: "manual"}
	if err := decodeLlmsTxtScalar(&s, val); err != nil {
		t.Fatalf("value scalar: %v", err)
	}
	if s.Mode != "manual" {
		t.Errorf("Mode = %q, want manual", s.Mode)
	}
	keep := LlmsTxtSpec{Mode: "auto"}
	null := &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!null", Value: "null"}
	if err := decodeLlmsTxtScalar(&keep, null); err != nil {
		t.Fatalf("null scalar: %v", err)
	}
	if keep.Mode != "auto" {
		t.Errorf("null scalar must keep default Mode, got %q", keep.Mode)
	}
}
