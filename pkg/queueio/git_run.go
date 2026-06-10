//ff:func feature=queueio type=client control=sequence
//ff:what git CLI 1회 실행 — dir 기준(-C), 실패 시 stderr를 에러 메시지로 환원
package queueio

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// gitRun executes one git command inside dir and returns its trimmed stdout.
// stderr is folded into the error so the caller's diagnostics carry git's
// own message (auth failures, non-fast-forward, ...).
func gitRun(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git %s: %v: %s", strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}
	return strings.TrimSpace(stdout.String()), nil
}
