//ff:func feature=cli type=command control=sequence topic=diagnostics
//ff:what runGate --json이 진단을 JSON 배열로 출력하고 위반 시에도 파싱 가능한지 검증
package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestRunGateJSON(t *testing.T) {
	dir := writeGateFixture(t, false)
	var out bytes.Buffer
	if err := runGate(&out, dir, "image-first", true, false); err == nil {
		t.Fatal("want error for violations")
	}
	var diags []blogyaml.Diagnostic
	if err := json.Unmarshal(out.Bytes(), &diags); err != nil {
		t.Fatalf("output is not JSON: %v\n%s", err, out.String())
	}
	if len(diags) != 2 || diags[0].Rule != "image-first" {
		t.Errorf("diags = %+v, want 2 image-first entries", diags)
	}
}
