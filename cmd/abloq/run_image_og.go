//ff:func feature=cli type=command control=sequence
//ff:what image og 실행 본체 — OGSpec 조립 후 렌더·기록, 마크다운/front matter 참조 경로 안내 출력
package main

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/img"
)

// runImageOG renders the OG card and prints how to reference it.
func runImageOG(out io.Writer, slug, title, brand, fontPath, outDir string) error {
	dst := filepath.Join(outDir, slug+".webp")
	spec := img.OGSpec{Title: title, Brand: brand, FontPath: fontPath, Out: dst}
	if err := img.OG(spec); err != nil {
		return err
	}
	fmt.Fprintln(out, dst)
	fmt.Fprintf(out, "  front matter: image: \"/images/%s.webp\"\n", slug)
	return nil
}
