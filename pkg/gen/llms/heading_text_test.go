//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what headingText가 section_labels 라벨을 우선 적용하고, 다언어 스코프에서만 "{lang}/" 접두를 붙이는지 검증
package llms

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestHeadingText(t *testing.T) {
	b := &blogyaml.Blog{}
	b.Geo.LlmsTxt.SectionLabels = map[string]string{"opinion": "Concept"}
	cases := []struct {
		name, lang, section string
		multi               bool
		want                string
	}{
		{"label override", "ko", "opinion", false, "Concept"},
		{"raw section fallback", "ko", "tech", false, "tech"},
		{"multi-language prefix", "ko", "opinion", true, "ko/Concept"},
		{"multi-language raw", "en", "tech", true, "en/tech"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := headingText(b, tc.lang, tc.section, tc.multi); got != tc.want {
				t.Errorf("headingText(%q, %q, %v) = %q, want %q", tc.lang, tc.section, tc.multi, got, tc.want)
			}
		})
	}
}
