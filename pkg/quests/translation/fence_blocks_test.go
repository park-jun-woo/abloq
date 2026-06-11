//ff:func feature=quest type=parser control=sequence
//ff:what fenceBlocks 검증 — 빈 줄 경계 블록 수, 코드 펜스 안 빈 줄은 블록을 쪼개지 않음
package translation

import "testing"

func TestFenceBlocks(t *testing.T) {
	md := "---\ntitle: x\n---\n\npara one\nstill one\n\npara two\n\n```sh\ncode\n\nsame block\n```\n\nlast\n"
	if n := fenceBlocks(docOf(t, "en", md)); n != 4 {
		t.Errorf("blocks = %d, want 4 (fence-internal blank must not split)", n)
	}
	origin, _ := passPair()
	if o, k := fenceBlocks(docOf(t, "en", origin)), fenceBlocks(docOf(t, "ko", cleanKoMD)); o != k {
		t.Errorf("fixture pair blocks differ: %d vs %d", o, k)
	}
}
