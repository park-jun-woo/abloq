//ff:func feature=cli type=command control=sequence topic=diagnostics
//ff:what validate 실행 본체 — dir/blog.yaml을 파싱·검증해 진단을 text 또는 JSON으로 출력(진단 존재 시 에러), text OK 시 비차단 경고(og-local-variants)도 표시
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// runValidate validates dir/blog.yaml and reports diagnostics on out.
// Non-blocking advisories (OGWarnings) print after the OK line in text mode
// and never affect the exit code.
func runValidate(out io.Writer, dir string, jsonOut bool) error {
	path := filepath.Join(dir, "blog.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	b, idx, diags := blogyaml.Parse(path, data)
	if len(diags) == 0 {
		diags = blogyaml.Validate(path, b, idx)
	}
	if jsonOut {
		if err := printDiagsJSON(out, diags); err != nil {
			return err
		}
	} else if len(diags) == 0 {
		fmt.Fprintf(out, "%s: OK\n", path)
		printDiagsText(out, blogyaml.OGWarnings(path, b, idx))
	} else {
		printDiagsText(out, diags)
	}
	if len(diags) > 0 {
		return fmt.Errorf("%s: %d issue(s) found", path, len(diags))
	}
	return nil
}
