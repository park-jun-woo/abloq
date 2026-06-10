//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what 골든 테스트 — 오류 예제 6종이 각각 의도한 룰ID 1건으로 실패하는지 검증
package blogyaml

import (
	"path/filepath"
	"testing"
)

func TestLoadInvalid(t *testing.T) {
	cases := []struct{ dir, rule string }{
		{"invalid-lang", "lang-bcp47"},
		{"invalid-heading", "heading-default-lang"},
		{"invalid-sections", "sections-empty"},
		{"invalid-threshold", "threshold-range"},
		{"invalid-baseurl", "baseurl-format"},
		{"invalid-crawlers", "crawlers-policy"},
	}
	for _, tc := range cases {
		_, diags, err := Load(filepath.Join("testdata", tc.dir, "blog.yaml"))
		if err != nil {
			t.Fatalf("%s: Load: %v", tc.dir, err)
		}
		if len(diags) != 1 {
			t.Fatalf("%s: want exactly 1 diagnostic, got %d: %v", tc.dir, len(diags), diags)
		}
		if diags[0].Rule != tc.rule {
			t.Errorf("%s: want rule %q, got %q (%s)", tc.dir, tc.rule, diags[0].Rule, diags[0])
		}
		if diags[0].Line < 1 {
			t.Errorf("%s: want line >= 1, got %d", tc.dir, diags[0].Line)
		}
	}
}
