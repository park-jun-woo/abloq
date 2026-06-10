//ff:func feature=cli type=output control=sequence
//ff:what printRotReport가 실패 인용만 출력하고 실패 수를 반환하는지 검증 (ok는 침묵)
package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/scan/evidence"
)

func TestPrintRotReport(t *testing.T) {
	var out bytes.Buffer
	n := printRotReport(&out, []evidence.Check{
		{URL: "https://ok.example/x", Lang: "ko", Section: "tech", Slug: "a", Status: "ok"},
		{URL: "https://gone.example/y", Lang: "ko", Section: "tech", Slug: "b", Status: "hard", ConsecutiveFailures: 1},
	})
	if n != 1 {
		t.Errorf("failing = %d, want 1", n)
	}
	if !strings.Contains(out.String(), "rot-check: https://gone.example/y hard (ko/tech/b)") {
		t.Errorf("failing line missing: %q", out.String())
	}
	if strings.Contains(out.String(), "ok.example") {
		t.Errorf("healthy citations must stay silent: %q", out.String())
	}
}
