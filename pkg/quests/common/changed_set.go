//ff:func feature=quest type=parser control=iteration dimension=1 topic=queue
//ff:what git status --porcelain -z로 작업트리 변경 파일 집합 수집 — untracked 포함, rename은 양쪽 경로, git 불가 시 에러 (퀘스트 공용)
//ff:why git diff HEAD는 untracked 신규 파일을 누락한다 — 범위 밖 파일을 새로 만드는 치즈가 queue-scope를 침묵 통과하므로 porcelain(untracked 포함)이 변경 집합의 단일 출처다 (Phase018 계획)
package common

import (
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

// ChangedSet returns the repository-relative paths of every working-tree
// change (modified, deleted, renamed both sides, and untracked files) at the
// blog instance root. The instance root must be the git toplevel (the abloq
// layout: blog.yaml at the repository root) — porcelain paths are toplevel-
// relative. A non-git root is an error: the queue quests judge against the
// HEAD baseline, so no repository means no gate.
func ChangedSet(root string) ([]string, error) {
	out, err := exec.Command("git", "-C", root, "status", "--porcelain", "-z").Output()
	if err != nil {
		return nil, fmt.Errorf("git status --porcelain at %s (the queue quests require a git repository): %w", root, err)
	}
	var paths []string
	tokens := strings.Split(string(out), "\x00")
	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]
		if len(tok) < 4 {
			continue
		}
		paths = append(paths, tok[3:])
		if tok[0] == 'R' || tok[0] == 'C' {
			i++ // the next token is the rename/copy origin — it changed too
			paths = append(paths, tokens[i])
		}
	}
	sort.Strings(paths)
	return paths, nil
}
