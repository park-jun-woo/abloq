//ff:func feature=gate type=parser control=iteration dimension=1 topic=baseline
//ff:what git diff --name-only HEAD로 작업 트리에서 변경된 파일 집합 수집 — git 저장소가 아니면 ok=false
package gate

import (
	"os/exec"
	"strings"
)

// gitChangedSet returns the dir-relative paths of tracked files that differ
// from HEAD. ok is false when dir is not inside a usable git repository.
func gitChangedSet(dir string) (map[string]bool, bool) {
	out, err := exec.Command("git", "-C", dir, "diff", "--relative", "--name-only", "HEAD").Output()
	if err != nil {
		return nil, false
	}
	changed := map[string]bool{}
	for _, ln := range strings.Split(string(out), "\n") {
		if ln != "" {
			changed[ln] = true
		}
	}
	return changed, true
}
