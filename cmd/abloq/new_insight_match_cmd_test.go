//ff:func feature=cli type=command control=sequence
//ff:what insight match 명령 검증 — 인자 2개 고정, 픽스처 실행 시 anchored 요약 출력
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewInsightMatchCmd(t *testing.T) {
	cmd := newInsightMatchCmd()
	if err := cmd.Args(cmd, []string{"one"}); err == nil {
		t.Errorf("want arg error with 1 arg, got nil")
	}
	insightPath, articlePath := writeInsightFixture(t,
		"topic: t\nsection: tech\nclaims:\n  - id: a\n    text: x\n    kind: claim\n    anchors: [\"ratchet never moves\"]\n")
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetArgs([]string{insightPath, articlePath})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("want clean run, got %v (output: %s)", err, out.String())
	}
	if !strings.Contains(out.String(), "anchored claims: 1/1") {
		t.Errorf("want anchored summary in output, got %q", out.String())
	}
}
