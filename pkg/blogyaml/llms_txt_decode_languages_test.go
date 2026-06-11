//ff:func feature=blogyaml type=parser control=sequence
//ff:what decodeLlmsLanguages가 시퀀스는 그대로, 스칼라는 1원소 리스트로 디코드하고 그 외 노드는 거부하는지 검증
package blogyaml

import (
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestDecodeLlmsLanguages(t *testing.T) {
	var seqDoc yaml.Node
	if err := yaml.Unmarshal([]byte("[ko, en]"), &seqDoc); err != nil {
		t.Fatalf("unmarshal sequence fixture: %v", err)
	}
	var s LlmsTxtSpec
	if err := decodeLlmsLanguages(&s, seqDoc.Content[0]); err != nil {
		t.Fatalf("sequence: %v", err)
	}
	if !reflect.DeepEqual(s.Languages, []string{"ko", "en"}) {
		t.Errorf("sequence Languages = %v, want [ko en]", s.Languages)
	}
	var scalarDoc yaml.Node
	if err := yaml.Unmarshal([]byte("base"), &scalarDoc); err != nil {
		t.Fatalf("unmarshal scalar fixture: %v", err)
	}
	var sc LlmsTxtSpec
	if err := decodeLlmsLanguages(&sc, scalarDoc.Content[0]); err != nil {
		t.Fatalf("scalar: %v", err)
	}
	if !reflect.DeepEqual(sc.Languages, []string{"base"}) {
		t.Errorf("scalar Languages = %v, want [base]", sc.Languages)
	}
	var mapDoc yaml.Node
	if err := yaml.Unmarshal([]byte("a: 1"), &mapDoc); err != nil {
		t.Fatalf("unmarshal mapping fixture: %v", err)
	}
	err := decodeLlmsLanguages(&sc, mapDoc.Content[0])
	if err == nil || !strings.Contains(err.Error(), "cannot unmarshal geo.llms_txt.languages: want base, all, or a sequence") {
		t.Errorf("mapping node error = %v, want union-shape message", err)
	}
}
