//ff:func feature=blogyaml type=parser control=sequence topic=diagnostics
//ff:what llmsUnknownKey가 KnownFields(true)와 동일한 형식("line N: field K not found in type T")의 TypeError를 만드는지 검증
package blogyaml

import (
	"errors"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestLlmsUnknownKey(t *testing.T) {
	key := &yaml.Node{Kind: yaml.ScalarNode, Value: "foo", Line: 6}
	err := llmsUnknownKey(key, "blogyaml.LlmsTxtSpec")
	var te *yaml.TypeError
	if !errors.As(err, &te) {
		t.Fatalf("want *yaml.TypeError, got %T (%v)", err, err)
	}
	if len(te.Errors) != 1 {
		t.Fatalf("want 1 error, got %v", te.Errors)
	}
	want := "line 6: field foo not found in type blogyaml.LlmsTxtSpec"
	if te.Errors[0] != want {
		t.Errorf("message = %q, want %q", te.Errors[0], want)
	}
}
