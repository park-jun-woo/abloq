//ff:func feature=claudemd type=generator control=sequence
//ff:what headingTable이 언어 순서대로 헤더와 행을 내고 헤딩이 없으면 빈 문자열을 반환하는지 검증
package claudemd

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestHeadingTable(t *testing.T) {
	out := headingTable(testBlog())
	if !strings.Contains(out, "| 키 | ko | en |") {
		t.Errorf("want language header row, got:\n%s", out)
	}
	if !strings.Contains(out, "| sources | 출처 | Sources |") {
		t.Errorf("want sources row, got:\n%s", out)
	}
	if got := headingTable(&blogyaml.Blog{}); got != "" {
		t.Errorf("empty headings must render nothing, got %q", got)
	}
}
