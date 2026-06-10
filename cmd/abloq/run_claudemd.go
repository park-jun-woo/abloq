//ff:func feature=cli type=command control=sequence
//ff:what claudemd 실행 본체 — 유효한 blog.yaml에서 CLAUDE.md를 렌더·기록하고 경로 출력 (멱등)
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/claudemd"
)

// runClaudeMD regenerates dir/CLAUDE.md from a valid blog.yaml.
func runClaudeMD(out io.Writer, dir string) error {
	b, err := loadValidBlog(out, dir)
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "CLAUDE.md")
	if err := os.WriteFile(path, claudemd.Render(b), 0o644); err != nil {
		return err
	}
	fmt.Fprintln(out, path)
	return nil
}
