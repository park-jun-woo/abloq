//ff:func feature=insight type=rule control=iteration dimension=1
//ff:what 역명세 골든 — parkjunwoo 사본 2편(기술·의견)에서 대응 claim 전부 검출, 심어둔 무관 claim만 미출현, 섹션 일치
package insight

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMatchGolden(t *testing.T) {
	root := filepath.Join("testdata", "parkjunwoo")
	articles := []string{
		filepath.Join(root, "content", "en", "tech", "ratchet-pattern.md"),
		filepath.Join(root, "content", "en", "opinion", "limits-of-natural-language.md"),
	}
	for _, articlePath := range articles {
		ins, errs, warns, err := Load(PathFor(articlePath))
		if err != nil || len(errs) != 0 || len(warns) != 0 {
			t.Fatalf("%s: want clean insight load, got errs=%v warns=%v err=%v", articlePath, errs, warns, err)
		}
		article, err := os.ReadFile(articlePath)
		if err != nil {
			t.Fatal(err)
		}
		res := Match(ins, articlePath, article)
		if res.Section != ins.Section {
			t.Errorf("%s: want section %q, got %q", articlePath, ins.Section, res.Section)
		}
		if len(res.Found) != len(ins.Claims)-1 {
			t.Errorf("%s: want all %d real claims found, got %v", articlePath, len(ins.Claims)-1, res.Found)
		}
		if len(res.Missing) != 1 || res.Missing[0].ID != "planted-unrelated" {
			t.Errorf("%s: want only planted-unrelated missing, got %v", articlePath, res.Missing)
		}
	}
}
