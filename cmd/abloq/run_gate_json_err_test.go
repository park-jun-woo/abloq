//ff:func feature=cli type=command control=sequence topic=diagnostics
//ff:what runGate --json이 출력 실패(errWriter) 시 쓰기 에러를 반환하는지 검증
package main

import "testing"

func TestRunGateJSONErr(t *testing.T) {
	dir := writeGateFixture(t, true)
	if err := runGate(errWriter{}, dir, "", true); err == nil {
		t.Error("want write error from JSON output, got nil")
	}
}
