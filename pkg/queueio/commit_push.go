//ff:func feature=queueio type=client control=sequence
//ff:what quests/queue 변경분 커밋·푸시 — 동일 내용이면 no-op, 푸시 충돌 시 pull --rebase 후 1회 재시도
package queueio

import "fmt"

// commitPush commits and pushes the quests/queue changes of the work clone.
// A clean tree is the idempotent no-op (deterministic serialization makes
// re-exports byte-identical). A rejected push gets one pull --rebase retry;
// a second failure surfaces as an error and the next cycle tries again.
func commitPush(cfg Config) error {
	status, err := gitRun(cfg.Workdir, "status", "--porcelain", "--", "quests/queue")
	if err != nil {
		return err
	}
	if status == "" {
		return nil
	}
	if _, err := gitRun(cfg.Workdir, "add", "--", "quests/queue"); err != nil {
		return err
	}
	_, err = gitRun(cfg.Workdir,
		"-c", "user.name="+cfg.AuthorName,
		"-c", "user.email="+cfg.AuthorEmail,
		"commit", "-m", "abloqd: queue export")
	if err != nil {
		return err
	}
	if _, err := gitRun(cfg.Workdir, "push", "origin", "HEAD"); err == nil {
		return nil
	}
	if _, err := gitRun(cfg.Workdir, "pull", "--rebase"); err != nil {
		return fmt.Errorf("push rejected and rebase failed: %w", err)
	}
	_, err = gitRun(cfg.Workdir, "push", "origin", "HEAD")
	return err
}
