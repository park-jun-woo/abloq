//ff:func feature=gate type=parser control=iteration dimension=1 topic=baseline
//ff:what 임시 디렉토리를 git 저장소로 초기화하고 현재 내용을 커밋 — 베이스라인 테스트 준비
package gate

import (
	"os/exec"
	"testing"
)

func initGitRepo(t *testing.T, dir string) {
	t.Helper()
	cmds := [][]string{
		{"init", "-q"},
		{"config", "user.email", "test@example.com"},
		{"config", "user.name", "Test"},
		{"add", "-A"},
		{"commit", "-q", "-m", "baseline"},
	}
	for _, args := range cmds {
		cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v: %s", args, err, out)
		}
	}
}
