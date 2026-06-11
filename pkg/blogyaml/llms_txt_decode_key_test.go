//ff:func feature=blogyaml type=parser control=iteration dimension=1
//ff:what decodeLlmsTxtKey가 객체 폼 6키(mode/languages/header/pinned/section_labels/max_summary)를 디코드하고 미지 키는 거부하는지 검증
package blogyaml

import (
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestDecodeLlmsTxtKey(t *testing.T) {
	src := "mode: auto\nlanguages: all\nheader: H\npinned:\n  - title: T\n    url: /x\nsection_labels:\n  tech: Pattern\nmax_summary: 9\n"
	var doc yaml.Node
	if err := yaml.Unmarshal([]byte(src), &doc); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	m := doc.Content[0]
	var s LlmsTxtSpec
	for i := 0; i+1 < len(m.Content); i += 2 {
		if err := decodeLlmsTxtKey(&s, m.Content[i], m.Content[i+1]); err != nil {
			t.Fatalf("decodeLlmsTxtKey(%q): %v", m.Content[i].Value, err)
		}
	}
	want := LlmsTxtSpec{
		Mode:          "auto",
		Languages:     []string{"all"},
		Header:        "H",
		Pinned:        []LlmsPinned{{Title: "T", URL: "/x"}},
		SectionLabels: map[string]string{"tech": "Pattern"},
		MaxSummary:    9,
	}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("decoded = %+v, want %+v", s, want)
	}
	err := decodeLlmsTxtKey(&s, &yaml.Node{Kind: yaml.ScalarNode, Value: "foo", Line: 6}, m.Content[1])
	if err == nil || !strings.Contains(err.Error(), "line 6: field foo not found in type blogyaml.LlmsTxtSpec") {
		t.Errorf("unknown key error = %v, want KnownFields-shaped message", err)
	}
}
