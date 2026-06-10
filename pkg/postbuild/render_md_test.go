//ff:func feature=postbuild type=generator control=sequence
//ff:what RenderMD가 front matter를 제거하고 "# title" 헤더를 붙이는지, 제목 없으면 본문만 내는지 검증
package postbuild

import "testing"

func TestRenderMD(t *testing.T) {
	src := "---\ntitle: \"Hello World\"\ndate: 2026-01-01\n---\n\nFirst line.\n"
	got := string(RenderMD([]byte(src)))
	want := "# Hello World\n\nFirst line.\n"
	if got != want {
		t.Errorf("RenderMD = %q, want %q", got, want)
	}
	noTitle := string(RenderMD([]byte("---\ndate: 2026-01-01\n---\nbody\n")))
	if noTitle != "body\n" {
		t.Errorf("RenderMD without title = %q, want %q", noTitle, "body\n")
	}
}
