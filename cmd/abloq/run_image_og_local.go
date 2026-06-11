//ff:func feature=cli type=command control=sequence
//ff:what local OG 실행 — 현행 결정론 경로 그대로(OGSpec 조립→렌더·기록→참조 안내), Provider 비경유로 바이트 동일을 구조로 보장
package main

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/img"
)

// runImageOGLocal renders the deterministic OG card and prints how to
// reference it — byte-for-byte the pre-Phase022 behavior.
func runImageOGLocal(out io.Writer, opts imageOGOpts) error {
	dst := filepath.Join(opts.OutDir, opts.Slug+".webp")
	spec := img.OGSpec{Title: opts.Title, Brand: opts.Brand, FontPath: opts.FontPath, Out: dst}
	if err := img.OG(spec); err != nil {
		return err
	}
	fmt.Fprintln(out, dst)
	fmt.Fprintf(out, "  front matter: image: \"/images/%s.webp\"\n", opts.Slug)
	return nil
}
