//ff:func feature=cli type=command control=sequence
//ff:what postbuild md 실행 본체 — content/ 전 글을 public/ 옆 .md로 기록하고 생성 수 출력
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/abloq/pkg/postbuild"
)

// runPostbuildMD writes the parallel .md files and reports the count.
func runPostbuildMD(out io.Writer, dir string) error {
	n, err := postbuild.MD(dir)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "postbuild md: %d file(s)\n", n)
	return nil
}
