//ff:func feature=image type=generator control=iteration dimension=1
//ff:what WrapText가 좁은 폭에서 여러 줄로 나누고, 넓은 폭에서 한 줄을 유지하며, 각 줄이 폭을 지키는지 검증
package img

import (
	"testing"

	"golang.org/x/image/font"
)

func TestWrapText(t *testing.T) {
	face, err := LoadFace("", 32)
	if err != nil {
		t.Fatalf("LoadFace: %v", err)
	}
	text := "the quick brown fox jumps over the lazy dog"
	narrow := WrapText(face, text, 200)
	if len(narrow) < 2 {
		t.Errorf("narrow wrap lines = %d, want >= 2", len(narrow))
	}
	for _, line := range narrow {
		if w := font.MeasureString(face, line).Ceil(); w > 200 && len(line) > 0 {
			t.Errorf("line %q is %dpx, exceeds 200px", line, w)
		}
	}
	if wide := WrapText(face, text, 5000); len(wide) != 1 {
		t.Errorf("wide wrap lines = %d, want 1", len(wide))
	}
}
