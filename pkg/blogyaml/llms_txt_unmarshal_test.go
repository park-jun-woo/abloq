//ff:func feature=blogyaml type=parser control=sequence
//ff:what LlmsTxtSpec.UnmarshalYAML가 스칼라는 mode 단축형으로, 매핑은 객체 폼으로 분기하고 그 외 노드는 거부하는지 검증
package blogyaml

import (
	"errors"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestLlmsTxtSpecUnmarshalYAML(t *testing.T) {
	var scalar LlmsTxtSpec
	if err := yaml.Unmarshal([]byte("manual\n"), &scalar); err != nil {
		t.Fatalf("scalar shorthand: %v", err)
	}
	if scalar.Mode != "manual" {
		t.Errorf("shorthand Mode = %q, want manual", scalar.Mode)
	}
	var obj LlmsTxtSpec
	if err := yaml.Unmarshal([]byte("mode: auto\nmax_summary: 9\n"), &obj); err != nil {
		t.Fatalf("object form: %v", err)
	}
	if obj.Mode != "auto" || obj.MaxSummary != 9 {
		t.Errorf("object form = %+v, want Mode auto, MaxSummary 9", obj)
	}
	var seq LlmsTxtSpec
	err := yaml.Unmarshal([]byte("- a\n"), &seq)
	var te *yaml.TypeError
	if !errors.As(err, &te) || !strings.Contains(te.Errors[0], "cannot unmarshal geo.llms_txt: want a string (auto|manual|off) or a mapping") {
		t.Errorf("sequence node error = %v, want union-shape TypeError", err)
	}
}
