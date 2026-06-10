//ff:func feature=cli type=output control=sequence topic=gsc
//ff:what printGscRows가 스냅샷 행과 일자·행 합계 한 줄을 출력하는지 검증
package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/visibility/gsc"
)

func TestPrintGscRows(t *testing.T) {
	var out bytes.Buffer
	printGscRows(&out, gsc.Result{
		Rows: []gsc.Snapshot{
			{SnapDate: "2026-06-08", Page: "https://t.example.com/a/", Impressions: 120, Clicks: 3, AvgPosition: 4.2},
		},
		Days: 1,
	})
	got := out.String()
	if !strings.Contains(got, "2026-06-08") || !strings.Contains(got, "https://t.example.com/a/") {
		t.Errorf("row missing:\n%s", got)
	}
	if !strings.Contains(got, "gsc: 1 closed day(s), 1 row(s)") {
		t.Errorf("summary missing:\n%s", got)
	}
}
