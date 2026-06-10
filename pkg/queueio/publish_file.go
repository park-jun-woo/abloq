//ff:func feature=queueio type=client control=sequence
//ff:what 일반 파일 1개 발행 — 클론 최신화 → 쓰기 → 커밋·푸시, 동일 내용이면 no-op(false) (리포트 발행 사본, Phase014)
package queueio

import (
	"os"
	"path/filepath"
)

// PublishFile publishes one file at relPath inside the work clone: pull,
// write, commit and push. Identical content is the idempotent no-op
// (committed=false) — the report's now-free markdown makes regeneration
// byte-identical. The caller decides the path (e.g. reports/<ym>.md); the
// queue export contract is untouched.
func PublishFile(cfg Config, relPath string, data []byte) (bool, error) {
	if err := ensureClone(cfg); err != nil {
		return false, err
	}
	path := filepath.Join(cfg.Workdir, relPath)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return false, err
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return false, err
	}
	return commitPushPath(cfg, relPath, "abloqd: publish "+relPath)
}
