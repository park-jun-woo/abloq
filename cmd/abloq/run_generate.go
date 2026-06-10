//ff:func feature=cli type=command control=sequence
//ff:what generate 실행 본체 — blog.yaml 검증 후 파생물 4종을 생성·기록하고 경로를 출력 (멱등)
package main

import (
	"io"

	"github.com/park-jun-woo/abloq/pkg/gen"
)

// runGenerate builds and writes all derived files for the blog rooted at dir.
func runGenerate(out io.Writer, dir string) error {
	b, err := loadValidBlog(out, dir)
	if err != nil {
		return err
	}
	outs := gen.Build(dir, b)
	if err := gen.Write(dir, outs); err != nil {
		return err
	}
	printOutputs(out, dir, outs)
	return nil
}
