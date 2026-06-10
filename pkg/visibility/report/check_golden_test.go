//ff:func feature=visibility type=generator control=sequence topic=report
//ff:what 테스트 헬퍼 — 산출물이 testdata 골든과 바이트 동일한지 비교, UPDATE_GOLDEN=1이면 골든 갱신
package report

import (
	"os"
	"testing"
)

func checkGolden(t *testing.T, path string, got []byte) {
	t.Helper()
	if os.Getenv("UPDATE_GOLDEN") == "1" {
		if err := os.WriteFile(path, got, 0o644); err != nil {
			t.Fatal(err)
		}
		return
	}
	want, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("golden missing (run with UPDATE_GOLDEN=1): %v", err)
	}
	if string(want) != string(got) {
		t.Errorf("%s drifted:\n--- want ---\n%s\n--- got ---\n%s", path, want, got)
	}
}
