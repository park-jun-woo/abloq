//ff:func feature=cli type=output control=iteration dimension=1 topic=gsc
//ff:what GSC 수집 결과 출력 — 스냅샷 행(일자 노출 클릭 평균순위 페이지)과 일자·행 합계 한 줄
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/abloq/pkg/visibility/gsc"
)

// printGscRows prints one stateless GSC collection: the per-(day, page)
// rows and a one-line summary.
func printGscRows(out io.Writer, res gsc.Result) {
	fmt.Fprintln(out, "gsc snapshots (date impressions clicks avg_position page):")
	for _, r := range res.Rows {
		fmt.Fprintf(out, "  %s  %6d %5d %6.1f  %s\n", r.SnapDate, r.Impressions, r.Clicks, r.AvgPosition, r.Page)
	}
	fmt.Fprintf(out, "gsc: %d closed day(s), %d row(s)\n", res.Days, len(res.Rows))
}
