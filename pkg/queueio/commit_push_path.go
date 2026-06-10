//ff:func feature=queueio type=client control=sequence
//ff:what 지정 경로 변경분 커밋·푸시 — 동일 내용이면 no-op(false), 푸시 충돌 시 pull --rebase 후 1회 재시도 (큐 전용에서 일반 파일 쓰기로 일반화, Phase014)
package queueio

import "fmt"

// commitPushPath commits and pushes the changes under relPath in the work
// clone. A clean tree is the idempotent no-op (deterministic content makes
// re-runs byte-identical) and reports committed=false. A rejected push gets
// one pull --rebase retry; a second failure surfaces as an error and the
// next cycle tries again. Phase014 split this out of commitPush so the
// report publisher shares the git labor without the queue serialization.
func commitPushPath(cfg Config, relPath, msg string) (bool, error) {
	status, err := gitRun(cfg.Workdir, "status", "--porcelain", "--", relPath)
	if err != nil {
		return false, err
	}
	if status == "" {
		return false, nil
	}
	if _, err := gitRun(cfg.Workdir, "add", "--", relPath); err != nil {
		return false, err
	}
	_, err = gitRun(cfg.Workdir,
		"-c", "user.name="+cfg.AuthorName,
		"-c", "user.email="+cfg.AuthorEmail,
		"commit", "-m", msg)
	if err != nil {
		return false, err
	}
	if _, err := gitRun(cfg.Workdir, "push", "origin", "HEAD"); err == nil {
		return true, nil
	}
	if _, err := gitRun(cfg.Workdir, "pull", "--rebase"); err != nil {
		return false, fmt.Errorf("push rejected and rebase failed: %w", err)
	}
	if _, err := gitRun(cfg.Workdir, "push", "origin", "HEAD"); err != nil {
		return false, err
	}
	return true, nil
}
