//ff:func feature=cli type=output control=iteration dimension=1
//ff:what 실행 전 비용 가시화 — "생성 예정 N건"과 안별 ×샘플 수·모델 내역을 echo (추정 비용 자체 계산은 하지 않는다)
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/abloq/pkg/img"
)

// printOGPlan announces what is about to be generated, per model, before any
// API call burns quota. Cost figures are never invented here — only counts
// and model ids.
func printOGPlan(out io.Writer, runs []img.OGVariant, count int) {
	fmt.Fprintf(out, "planned: %d image(s)\n", len(runs)*count)
	for _, r := range runs {
		fmt.Fprintf(out, "  %s x%d (model %s)\n", r.Name, count, r.Model)
	}
}
