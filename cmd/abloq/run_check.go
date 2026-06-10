//ff:func feature=cli type=command control=sequence topic=drift
//ff:what check 실행 본체 — 파생물을 재생성해 디스크와 바이트 비교, 드리프트를 진단으로 출력하고 에러 반환 (CI 게이트)
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/abloq/pkg/gen"
)

// runCheck verifies that every derived file on disk matches a fresh regeneration.
func runCheck(out io.Writer, dir string) error {
	b, err := loadValidBlog(out, dir)
	if err != nil {
		return err
	}
	diags := gen.Check(dir, gen.Build(dir, b))
	if len(diags) == 0 {
		fmt.Fprintf(out, "%s: derived files in sync\n", dir)
		return nil
	}
	printDiagsText(out, diags)
	return fmt.Errorf("%s: %d derived file issue(s) found", dir, len(diags))
}
