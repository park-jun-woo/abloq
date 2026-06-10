//ff:func feature=cli type=output control=iteration dimension=1 topic=diagnostics
//ff:what 진단 목록을 "파일:라인 [룰ID] 메시지" 한 줄씩 출력
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// printDiagsText writes one "file:line [rule] message" line per diagnostic.
func printDiagsText(out io.Writer, diags []blogyaml.Diagnostic) {
	for _, d := range diags {
		fmt.Fprintln(out, d.String())
	}
}
