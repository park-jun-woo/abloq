//ff:func feature=cli type=command control=sequence topic=diagnostics
//ff:what runGate가 정규 저장소를 통과시키고 --rule 지정 시 해당 룰만 실행하는지 검증
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunGate(t *testing.T) {
	dir := writeGateFixture(t, true)
	var out bytes.Buffer
	if err := runGate(&out, dir, "", false, false); err != nil {
		t.Fatalf("runGate on canonical fixture: %v\noutput: %s", err, out.String())
	}
	if !strings.Contains(out.String(), "pass the gate") {
		t.Errorf("want pass summary, got %q", out.String())
	}
	bad := writeGateFixture(t, false)
	out.Reset()
	if err := runGate(&out, bad, "section-order", false, false); err != nil {
		t.Fatalf("rule filter: section-order must pass on image-less fixture, got %v", err)
	}
	out.Reset()
	if err := runGate(&out, bad, "image-first", false, false); err == nil {
		t.Fatalf("rule filter: image-first must fail, output: %s", out.String())
	}
}
