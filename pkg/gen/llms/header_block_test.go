//ff:func feature=gen type=generator control=sequence
//ff:what headerBlock이 header 마크다운을 빈 줄로 감싸 내고, trailing 개행을 정규화하며, 없으면 빈 문자열인지 검증
package llms

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestHeaderBlock(t *testing.T) {
	b := &blogyaml.Blog{}
	if got := headerBlock(b); got != "" {
		t.Errorf("absent header = %q, want empty", got)
	}
	b.Geo.LlmsTxt.Header = "Positioning paragraph.\n\n"
	if got := headerBlock(b); got != "\nPositioning paragraph.\n" {
		t.Errorf("headerBlock = %q, want trailing newlines normalized", got)
	}
	b.Geo.LlmsTxt.Header = "\n"
	if got := headerBlock(b); got != "" {
		t.Errorf("newline-only header = %q, want empty", got)
	}
}
