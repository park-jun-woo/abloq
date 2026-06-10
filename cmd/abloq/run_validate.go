//ff:func feature=cli type=command control=sequence topic=diagnostics
//ff:what validate 실행 본체 — dir/blog.yaml을 Load하고 진단을 text 또는 JSON으로 출력, 진단 존재 시 에러 반환
package main

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// runValidate validates dir/blog.yaml and reports diagnostics on out.
func runValidate(out io.Writer, dir string, jsonOut bool) error {
	path := filepath.Join(dir, "blog.yaml")
	_, diags, err := blogyaml.Load(path)
	if err != nil {
		return err
	}
	if jsonOut {
		if err := printDiagsJSON(out, diags); err != nil {
			return err
		}
	} else if len(diags) == 0 {
		fmt.Fprintf(out, "%s: OK\n", path)
	} else {
		printDiagsText(out, diags)
	}
	if len(diags) > 0 {
		return fmt.Errorf("%s: %d issue(s) found", path, len(diags))
	}
	return nil
}
