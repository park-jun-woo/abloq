//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what 테스트 헬퍼 — writeFixture 인스턴스를 git init + 전량 커밋해 HEAD 기준선이 있는 저장소로 만든다
package common

import (
	"os/exec"
	"testing"
)

func gitFixture(t *testing.T, root string) {
	t.Helper()
	for _, args := range [][]string{
		{"init", "-q", "-b", "main"},
		{"add", "-A"},
		{"-c", "user.name=t", "-c", "user.email=t@test", "commit", "-q", "-m", "fixture"},
	} {
		cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
}
