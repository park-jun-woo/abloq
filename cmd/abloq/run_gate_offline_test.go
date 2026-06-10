//ff:func feature=cli type=command control=sequence topic=diagnostics
//ff:what runGate offline 모드와 gate --offline 플래그가 정규 픽스처에서 네트워크 룰 없이 통과하는지 검증
package main

import (
	"bytes"
	"testing"
)

func TestRunGateOffline(t *testing.T) {
	dir := writeGateFixture(t, true)
	var out bytes.Buffer
	if err := runGate(&out, dir, "", false, true); err != nil {
		t.Fatalf("offline runGate on canonical fixture: %v\noutput: %s", err, out.String())
	}
	cmd := newGateCmd()
	out.Reset()
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"--offline", dir})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("gate --offline: %v\noutput: %s", err, out.String())
	}
}
