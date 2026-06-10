//ff:func feature=cli type=output control=sequence topic=diagnostics
//ff:what 진단 목록을 JSON 배열로 출력 (--json), 진단 없음 = []
package main

import (
	"encoding/json"
	"io"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// printDiagsJSON writes the diagnostics as a JSON array (empty slice -> []).
func printDiagsJSON(out io.Writer, diags []blogyaml.Diagnostic) error {
	if diags == nil {
		diags = []blogyaml.Diagnostic{}
	}
	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	return enc.Encode(diags)
}
