//ff:func feature=queueio type=client control=sequence
//ff:what 테스트 헬퍼 — git 명령 실행 실패 시 즉시 t.Fatal
package queueio

import "testing"

// mustGit runs one git command and fails the test on error.
func mustGit(t *testing.T, dir string, args ...string) string {
	t.Helper()
	out, err := gitRun(dir, args...)
	if err != nil {
		t.Fatalf("git %v: %v", args, err)
	}
	return out
}
