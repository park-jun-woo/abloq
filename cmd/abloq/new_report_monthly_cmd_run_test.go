//ff:func feature=cli type=command control=sequence topic=report
//ff:what report monthly 명령 실행이 리포트 헤더를 출력하고 인자 생략 시 blog.yaml 없는 cwd 기본값에서 에러인지 검증
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewReportMonthlyCmdRun(t *testing.T) {
	dir := writeBlogFixture(t)
	cmd := newReportMonthlyCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetArgs([]string{dir, "--ym", "2026-04"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if !strings.Contains(out.String(), "# Visibility report 2026-04") {
		t.Errorf("report output missing:\n%s", out.String())
	}
	// No positional arg — the repository defaults to the current directory,
	// which has no blog.yaml here, so the run errors (default branch covered).
	bad := newReportMonthlyCmd()
	bad.SetOut(&out)
	bad.SetErr(&out)
	bad.SetArgs([]string{"--ym", "2026-04"})
	if err := bad.Execute(); err == nil {
		t.Error("running in a non-blog cwd must error")
	}
}
