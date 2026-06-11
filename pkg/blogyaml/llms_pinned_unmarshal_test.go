//ff:func feature=blogyaml type=parser control=sequence
//ff:what LlmsPinned.UnmarshalYAML가 매핑은 strict 디코드하고, 비매핑 노드는 거부하며, 키별 에러를 전부 모으는지 검증
package blogyaml

import (
	"errors"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestLlmsPinnedUnmarshalYAML(t *testing.T) {
	var p LlmsPinned
	if err := yaml.Unmarshal([]byte("title: T\nurl: /x\n"), &p); err != nil {
		t.Fatalf("valid mapping: %v", err)
	}
	if p.Title != "T" || p.URL != "/x" {
		t.Errorf("decoded = %+v, want Title T, URL /x", p)
	}
	var te *yaml.TypeError
	var seq LlmsPinned
	err := yaml.Unmarshal([]byte("- a\n"), &seq)
	if !errors.As(err, &te) || !strings.Contains(te.Errors[0], "cannot unmarshal geo.llms_txt.pinned entry: want a mapping") {
		t.Errorf("sequence node error = %v, want want-a-mapping TypeError", err)
	}
	var bad LlmsPinned
	err = yaml.Unmarshal([]byte("foo: 1\nbar: 2\n"), &bad)
	if !errors.As(err, &te) || len(te.Errors) != 2 {
		t.Fatalf("two unknown keys: want TypeError with 2 errors, got %v", err)
	}
	if !strings.Contains(te.Errors[0], "field foo") || !strings.Contains(te.Errors[1], "field bar") {
		t.Errorf("aggregated errors = %v, want foo then bar", te.Errors)
	}
}
