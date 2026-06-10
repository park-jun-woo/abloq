//ff:func feature=cli type=command control=sequence topic=diagnostics
//ff:what gate 실행 본체 — blog.yaml 로드, 대상 글 수집, 룰 실행(전체 또는 --rule 1개), 진단 출력과 에러 반환
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/abloq/pkg/gate"
)

// runGate runs the structure gate rules against the blog rooted at dir.
func runGate(out io.Writer, dir, ruleID string, jsonOut bool) error {
	b, err := loadValidBlog(out, dir)
	if err != nil {
		return err
	}
	var ruleIDs []string
	if ruleID != "" {
		ruleIDs = append(ruleIDs, ruleID)
	}
	arts := gate.Discover(dir, b)
	diags := gate.Run(gate.NewTarget(dir, b, arts), ruleIDs...)
	if jsonOut {
		if err := printDiagsJSON(out, diags); err != nil {
			return err
		}
	} else if len(diags) == 0 {
		fmt.Fprintf(out, "%s: %d article(s) pass the gate\n", dir, len(arts))
	} else {
		printDiagsText(out, diags)
	}
	if len(diags) > 0 {
		return fmt.Errorf("%s: %d gate violation(s) found", dir, len(diags))
	}
	return nil
}
