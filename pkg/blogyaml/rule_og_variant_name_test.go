//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleOGVariantName이 비URL-safe·빈 이름·예약어 default·중복만 거부하는지 검증
package blogyaml

import (
	"strings"
	"testing"
)

func TestRuleOGVariantName(t *testing.T) {
	cases := []struct {
		name      string
		variants  []OGVariant
		wantDiags int
		wantIn    string
	}{
		{"none", nil, 0, ""},
		{"valid", []OGVariant{{Name: "minimal"}, {Name: "photo-2"}, {Name: "a_b"}}, 0, ""},
		{"empty name", []OGVariant{{Name: ""}}, 1, "URL-safe"},
		{"unsafe chars", []OGVariant{{Name: "Phо то!"}}, 1, "URL-safe"},
		{"uppercase", []OGVariant{{Name: "Minimal"}}, 1, "URL-safe"},
		{"reserved default", []OGVariant{{Name: "default"}}, 1, "reserved"},
		{"duplicate", []OGVariant{{Name: "x"}, {Name: "x"}}, 1, "duplicates"},
	}
	for _, tc := range cases {
		b := &Blog{Image: Image{OG: ImageOG{Variants: tc.variants}}}
		diags := ruleOGVariantName("blog.yaml", b, lineIndex{})
		if len(diags) != tc.wantDiags {
			t.Errorf("%s: %d diagnostics (%v), want %d", tc.name, len(diags), diags, tc.wantDiags)
			continue
		}
		if tc.wantDiags > 0 && !strings.Contains(diags[0].Message, tc.wantIn) {
			t.Errorf("%s: message %q, want %q", tc.name, diags[0].Message, tc.wantIn)
		}
	}
}
