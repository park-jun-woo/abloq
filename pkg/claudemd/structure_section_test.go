//ff:func feature=claudemd type=generator control=iteration dimension=1
//ff:what structureSection이 섹션 순서/헤딩 표/front matter 스키마/임계값 4종을 포함하는지 검증
package claudemd

import (
	"strings"
	"testing"
)

func TestStructureSection(t *testing.T) {
	out := structureSection(testBlog())
	wants := []string{
		"body → sources",
		"| sources | 출처 | Sources |",
		"`title`",
		"`lastmod`",
		"min_sources=1 · min_internal_links=2 · freshness_days=90 · min_meaningful_diff=10",
	}
	for _, w := range wants {
		if !strings.Contains(out, w) {
			t.Errorf("want %q in structure section, got:\n%s", w, out)
		}
	}
}
