//ff:func feature=gen type=rule control=iteration dimension=1 topic=drift
//ff:what 재생성 파생물과 디스크 파일을 바이트 비교 — 누락·드리프트를 파일·룰ID 단위 진단으로 반환
package gen

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// Check compares freshly built outputs against the files on disk under dir.
// Missing files and byte mismatches each yield one diagnostic; empty means in sync.
func Check(dir string, outs []Output) []blogyaml.Diagnostic {
	var diags []blogyaml.Diagnostic
	for _, o := range outs {
		path := filepath.Join(dir, o.Path)
		got, err := os.ReadFile(path)
		if err != nil {
			diags = append(diags, blogyaml.Diagnostic{File: path, Line: 1, Rule: ruleFor(o.Path),
				Message: "derived file missing — run `abloq generate`"})
			continue
		}
		if !bytes.Equal(got, o.Data) {
			diags = append(diags, driftDiag(path, ruleFor(o.Path), o.Data, got))
		}
	}
	return diags
}
