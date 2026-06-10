//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what sourcesLines 케이스 — sources 스팬(헤딩~다음 인식 섹션 직전) 포함, 다음 섹션 제외, sources 미인식이면 빈 집합
package gate

import "testing"

func TestSourcesLines(t *testing.T) {
	b := loadGateBlog(t)
	content := "Intro.\n\n## Sources\n\n- a\n- b\n\n## Changelog\n\n- 2026-01-01 first\n"
	d := ParseArticle(b, "en", content)
	in := sourcesLines(d)
	for ln, want := range map[int]bool{
		0: false, // intro prose
		2: true,  // ## Sources heading
		4: true,  // - a
		5: true,  // - b
		7: false, // ## Changelog heading
		9: false, // changelog entry
	} {
		if in[ln] != want {
			t.Errorf("sourcesLines[%d] = %v, want %v (line %q)", ln, in[ln], want, d.BodyLines[ln])
		}
	}
	t.Run("unrecognized sources heading yields no lines", func(t *testing.T) {
		// "ja" has no heading declared in blog.yaml structure.headings.sources.
		if in := sourcesLines(ParseArticle(b, "ja", content)); len(in) != 0 {
			t.Errorf("want empty set without a recognized sources section, got %v", in)
		}
	})
}
