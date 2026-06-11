//ff:func feature=blogyaml type=parser control=sequence
//ff:what geo.llms_txt 객체 폼 파싱 검증 — mode/languages 시퀀스/header/pinned/section_labels/max_summary 전 필드 디코드
package blogyaml

import (
	"reflect"
	"testing"
)

func TestParseLlmsTxtObject(t *testing.T) {
	src := []byte(`languages: [ko, en]
sections: [tech]
geo:
  llms_txt:
    mode: auto
    languages: [en, ko]
    header: |
      Positioning paragraph.
    pinned:
      - title: Master Index
        url: /reins.md
        desc: Index of everything
        group: Core Content
    section_labels:
      tech: Pattern
    max_summary: 120
`)
	b, _, diags := Parse("blog.yaml", src)
	if len(diags) != 0 {
		t.Fatalf("want 0 diagnostics, got %v", diags)
	}
	got := b.Geo.LlmsTxt
	want := LlmsTxtSpec{
		Mode:          "auto",
		Languages:     []string{"en", "ko"},
		Header:        "Positioning paragraph.\n",
		Pinned:        []LlmsPinned{{Title: "Master Index", URL: "/reins.md", Desc: "Index of everything", Group: "Core Content"}},
		SectionLabels: map[string]string{"tech": "Pattern"},
		MaxSummary:    120,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("LlmsTxt = %+v, want %+v", got, want)
	}
}
