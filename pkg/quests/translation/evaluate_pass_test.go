//ff:func feature=quest type=frame control=sequence
//ff:what 게이트 전 룰 통과 검증 — 클린 번역 쌍을 reins 레벨집계로 평가해 PASS (hugo는 성공 스크립트 주입)
package translation

import (
	"os"
	"path/filepath"
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestEvaluatePass(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "hugo")
	if err := os.WriteFile(bin, []byte("#!/bin/sh\nexit 0\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	orig := hugoLook
	hugoLook = func(string) (string, error) { return bin, nil }
	defer func() { hugoLook = orig }()
	origin, ko := passPair()
	sub := subWith(t, writeInstance(t), origin, ko)
	v := rgate.Evaluate(Definition{}.Rules(), rgate.Context{Submission: sub})
	if v.Outcome != quest.OutPass {
		t.Fatalf("Outcome = %s, facts = %+v — want PASS", v.Outcome, v.Facts)
	}
}
