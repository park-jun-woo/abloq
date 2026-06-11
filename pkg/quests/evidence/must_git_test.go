//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what 테스트 헬퍼 — git 명령 실행, 실패 시 출력 포함 즉시 실패
package evidence

import (
	"os/exec"
	"testing"
)

func mustGit(t *testing.T, root string, args ...string) {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %v: %v\n%s", args, err, out)
	}
}
