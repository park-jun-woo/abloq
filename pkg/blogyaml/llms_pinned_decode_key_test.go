//ff:func feature=blogyaml type=parser control=iteration dimension=1
//ff:what decodeLlmsPinnedKey가 title/url/desc/group 4키를 디코드하고 미지 키는 strict-동등 메시지로 거부하는지 검증
package blogyaml

import (
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestDecodeLlmsPinnedKey(t *testing.T) {
	var doc yaml.Node
	if err := yaml.Unmarshal([]byte("title: T\nurl: /x\ndesc: D\ngroup: G\n"), &doc); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	m := doc.Content[0]
	var p LlmsPinned
	for i := 0; i+1 < len(m.Content); i += 2 {
		if err := decodeLlmsPinnedKey(&p, m.Content[i], m.Content[i+1]); err != nil {
			t.Fatalf("decodeLlmsPinnedKey(%q): %v", m.Content[i].Value, err)
		}
	}
	want := LlmsPinned{Title: "T", URL: "/x", Desc: "D", Group: "G"}
	if p != want {
		t.Errorf("decoded = %+v, want %+v", p, want)
	}
	err := decodeLlmsPinnedKey(&p, &yaml.Node{Kind: yaml.ScalarNode, Value: "foo", Line: 9}, m.Content[1])
	if err == nil || !strings.Contains(err.Error(), "line 9: field foo not found in type blogyaml.LlmsPinned") {
		t.Errorf("unknown key error = %v, want KnownFields-shaped message", err)
	}
}
