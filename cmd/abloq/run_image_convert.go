//ff:func feature=cli type=command control=sequence
//ff:what image convert 실행 본체 — slug 기본값(원본 파일명) 결정, 변환 후 절감률과 참조 경로 출력
package main

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/img"
)

// runImageConvert converts src to WebP and reports the size change.
func runImageConvert(out io.Writer, src, slug, outDir string, maxWidth int) error {
	if slug == "" {
		base := filepath.Base(src)
		slug = strings.TrimSuffix(base, filepath.Ext(base))
	}
	dst := filepath.Join(outDir, slug+".webp")
	res, err := img.Convert(src, dst, maxWidth)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s (%dKB) -> %s (%dKB)\n", src, res.SrcBytes/1024, res.Dst, res.DstBytes/1024)
	fmt.Fprintf(out, "  markdown: ![alt text](/images/%s.webp)\n", slug)
	return nil
}
