//ff:func feature=blogyaml type=parser control=sequence
//ff:what decodeLlmsLanguagesScalar가 값 스칼라를 1원소 리스트로 정규화하고 null 스칼라는 기본값을 유지하는지 검증
package blogyaml

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestDecodeLlmsLanguagesScalar(t *testing.T) {
	var s LlmsTxtSpec
	val := &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: "all"}
	if err := decodeLlmsLanguagesScalar(&s, val); err != nil {
		t.Fatalf("value scalar: %v", err)
	}
	if !reflect.DeepEqual(s.Languages, []string{"all"}) {
		t.Errorf("Languages = %v, want [all]", s.Languages)
	}
	null := &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!null", Value: "null"}
	keep := LlmsTxtSpec{Languages: []string{"base"}}
	if err := decodeLlmsLanguagesScalar(&keep, null); err != nil {
		t.Fatalf("null scalar: %v", err)
	}
	if !reflect.DeepEqual(keep.Languages, []string{"base"}) {
		t.Errorf("null scalar must keep default Languages, got %v", keep.Languages)
	}
}
