//ff:func feature=quest type=rule control=sequence
//ff:what hugo-build 룰 검증 — 가짜 hugo(성공/실패 스크립트)로 0 에러 무발동·비0 종료 Fact(출력 인용) 확인
package translation

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

func TestRuleHugoBuild(t *testing.T) {
	r := ruleHugoBuild()
	if r.Meta.ID != "hugo-build" || r.Meta.Level != rgate.LevelFail {
		t.Errorf("Meta = %+v", r.Meta)
	}
	origin, ko := passPair()
	sub := subWith(t, writeInstance(t), origin, ko)
	bin := filepath.Join(t.TempDir(), "hugo")
	orig := hugoLook
	hugoLook = func(string) (string, error) { return bin, nil }
	defer func() { hugoLook = orig }()
	if err := os.WriteFile(bin, []byte("#!/bin/sh\nexit 0\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	if fired, fact := fireRule(t, r, sub); fired {
		t.Errorf("zero-error build: fired with %+v", fact)
	}
	if err := os.WriteFile(bin, []byte("#!/bin/sh\necho 'ERROR boom' >&2\nexit 1\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	fired, fact := fireRule(t, r, sub)
	if !fired || !strings.Contains(fact.Actual, "ERROR boom") {
		t.Errorf("failing build: fired=%v fact=%+v", fired, fact)
	}
	hugoLook = func(string) (string, error) { return "", fmt.Errorf("absent") }
	if fired, fact := fireRule(t, r, sub); !fired || fact.Expected != "hugo binary in PATH" {
		t.Errorf("hugo absent: fired=%v fact=%+v", fired, fact)
	}
	hugoLook = func(string) (string, error) { return bin, nil }
	t.Setenv("TMPDIR", filepath.Join(t.TempDir(), "missing-subdir"))
	if fired, fact := fireRule(t, r, sub); !fired || fact.Expected != "writable temp build dir" {
		t.Errorf("tempdir failure: fired=%v fact=%+v", fired, fact)
	}
}
