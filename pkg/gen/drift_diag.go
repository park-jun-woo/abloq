//ff:func feature=gen type=rule control=sequence topic=drift
//ff:what 드리프트 파생물 1개의 진단 생성 — 첫 불일치 라인과 기대/실제 줄 내용으로 어긋난 정책을 특정
package gen

import (
	"fmt"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// driftDiag describes one drifted derived file down to the first differing line.
func driftDiag(file, rule string, want, got []byte) blogyaml.Diagnostic {
	line, wantLine, gotLine := firstDiffLine(want, got)
	return blogyaml.Diagnostic{
		File: file, Line: line, Rule: rule,
		Message: fmt.Sprintf("drift from blog.yaml: want %q, got %q — run `abloq generate`", wantLine, gotLine),
	}
}
