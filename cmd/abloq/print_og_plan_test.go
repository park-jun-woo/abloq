//ff:func feature=cli type=output control=sequence
//ff:what printOGPlan 검증 — 총 건수(안 수×샘플 수)와 안별 ×샘플·모델 내역 echo
package main

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/img"
)

func TestPrintOGPlan(t *testing.T) {
	var out bytes.Buffer
	runs := []img.OGVariant{{Name: "default", Model: "m1"}, {Name: "bold", Model: "m2"}}
	printOGPlan(&out, runs, 3)
	want := "planned: 6 image(s)\n  default x3 (model m1)\n  bold x3 (model m2)\n"
	if out.String() != want {
		t.Errorf("printOGPlan = %q, want %q", out.String(), want)
	}
}
