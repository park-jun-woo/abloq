//ff:func feature=queueio type=client control=sequence
//ff:what 전용 작업 클론 준비 — 없으면 clone, 있으면 pull --rebase로 최신화 (ro 마운트 /blog는 절대 건드리지 않음)
package queueio

import (
	"os"
	"path/filepath"
)

// ensureClone makes cfg.Workdir a fresh clone of cfg.RepoURL. The exporter
// owns this dedicated work clone; the read-only /blog index mount stays
// untouched. An existing clone is fast-forwarded with pull --rebase so the
// consumed-sync sees the agents' latest deletion commits.
func ensureClone(cfg Config) error {
	if _, err := os.Stat(filepath.Join(cfg.Workdir, ".git")); err == nil {
		_, pullErr := gitRun(cfg.Workdir, "pull", "--rebase")
		return pullErr
	}
	if err := os.MkdirAll(filepath.Dir(cfg.Workdir), 0o755); err != nil {
		return err
	}
	_, err := gitRun("", "clone", cfg.RepoURL, cfg.Workdir)
	return err
}
