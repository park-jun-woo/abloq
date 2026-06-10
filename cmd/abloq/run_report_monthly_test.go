//ff:func feature=cli type=command control=sequence topic=report
//ff:what runReportMonthly가 부분 리포트 markdown(PARTIAL 명시·글별 표)을 출력하고 ym 형식 위반은 에러인지 검증
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunReportMonthly(t *testing.T) {
	dir := writeBlogFixture(t)
	var out bytes.Buffer
	if err := runReportMonthly(&out, dir, "2026-04", ""); err != nil {
		t.Fatalf("runReportMonthly: %v", err)
	}
	got := out.String()
	if !strings.Contains(got, "PARTIAL REPORT") {
		t.Error("output must declare itself partial")
	}
	if !strings.Contains(got, "# Visibility report 2026-04") {
		t.Errorf("report header missing:\n%s", got)
	}
	// The fixture post has no measurements — the cold-start date score.
	if !strings.Contains(got, "| ko/opinion/hello | 2026-01-02 |") {
		t.Errorf("article row missing:\n%s", got)
	}
	if err := runReportMonthly(&out, dir, "2026-4", ""); err == nil {
		t.Error("malformed ym must error")
	}
}
