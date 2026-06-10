//ff:func feature=cli type=command control=sequence topic=report
//ff:what runReportMonthly가 blog.yaml 없는 저장소와 존재하지 않는 로그 소스를 각각 에러로 내는지 검증
package main

import (
	"bytes"
	"testing"
)

func TestRunReportMonthlyErrors(t *testing.T) {
	var out bytes.Buffer
	if err := runReportMonthly(&out, t.TempDir(), "2026-04", ""); err == nil {
		t.Error("repo without blog.yaml must error")
	}
	dir := writeBlogFixture(t)
	if err := runReportMonthly(&out, dir, "2026-04", dir+"/missing-logs"); err == nil {
		t.Error("nonexistent log source must error")
	}
}
