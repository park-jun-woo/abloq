//ff:func feature=blogyaml type=parser control=sequence
//ff:what decodeLlmsTxtMap이 객체 폼 키-값 쌍을 전부 디코드하고, 키별 에러를 모아 하나의 TypeError로 반환하는지 검증
package blogyaml

import (
	"errors"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestDecodeLlmsTxtMap(t *testing.T) {
	var okDoc yaml.Node
	if err := yaml.Unmarshal([]byte("mode: manual\nmax_summary: 7\n"), &okDoc); err != nil {
		t.Fatalf("unmarshal valid fixture: %v", err)
	}
	var s LlmsTxtSpec
	if err := decodeLlmsTxtMap(&s, okDoc.Content[0]); err != nil {
		t.Fatalf("valid map: %v", err)
	}
	if s.Mode != "manual" || s.MaxSummary != 7 {
		t.Errorf("decoded = %+v, want Mode manual, MaxSummary 7", s)
	}
	var badDoc yaml.Node
	if err := yaml.Unmarshal([]byte("mode: auto\nfoo: 1\nbar: 2\n"), &badDoc); err != nil {
		t.Fatalf("unmarshal invalid fixture: %v", err)
	}
	err := decodeLlmsTxtMap(&s, badDoc.Content[0])
	var te *yaml.TypeError
	if !errors.As(err, &te) || len(te.Errors) != 2 {
		t.Fatalf("two unknown keys: want TypeError with 2 errors, got %v", err)
	}
	if !strings.Contains(te.Errors[0], "field foo") || !strings.Contains(te.Errors[1], "field bar") {
		t.Errorf("aggregated errors = %v, want foo then bar", te.Errors)
	}
}
