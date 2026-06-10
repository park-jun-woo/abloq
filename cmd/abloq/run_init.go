//ff:func feature=init type=command control=sequence
//ff:what init 실행 본체 — blog.yaml 작성 → 템플릿 복제 → 디렉토리 골격 → generate → CLAUDE.md → validate+gate 통과 확인
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/claudemd"
	"github.com/park-jun-woo/abloq/pkg/scaffold"
	"github.com/park-jun-woo/abloq/template"
)

// runInit scaffolds dir as a fresh abloq blog and leaves it in a state that
// passes validate, check and gate.
func runInit(out io.Writer, in io.Reader, dir string, o initOpts) error {
	if o.Interactive {
		o = promptOpts(out, in, o)
	}
	blogPath := filepath.Join(dir, "blog.yaml")
	if _, err := os.Stat(blogPath); err == nil {
		return fmt.Errorf("%s already exists — refusing to overwrite", blogPath)
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(blogPath, renderBlogYAML(o), 0o644); err != nil {
		return err
	}
	n, err := scaffold.Copy(template.Files(), dir)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "template: %d file(s) copied\n", n)
	if err := initContentDirs(dir, o.Languages, o.Sections); err != nil {
		return err
	}
	if err := runGenerate(out, dir); err != nil {
		return err
	}
	b, err := loadValidBlog(out, dir)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dir, "CLAUDE.md"), claudemd.Render(b), 0o644); err != nil {
		return err
	}
	fmt.Fprintln(out, filepath.Join(dir, "CLAUDE.md"))
	if err := runGate(out, dir, "", false, true); err != nil {
		return err
	}
	fmt.Fprintf(out, "%s: blog initialized — next: write content/%s/..., see CLAUDE.md\n", dir, o.Languages[0])
	return nil
}
