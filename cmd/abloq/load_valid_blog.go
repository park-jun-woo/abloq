//ff:func feature=cli type=command control=sequence topic=diagnostics
//ff:what dir/blog.yaml을 Load하고 진단이 있으면 출력 후 에러 — generate/check의 공통 전제(유효한 SSOT)
package main

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// loadValidBlog loads dir/blog.yaml and fails when validation finds issues.
func loadValidBlog(out io.Writer, dir string) (*blogyaml.Blog, error) {
	path := filepath.Join(dir, "blog.yaml")
	b, diags, err := blogyaml.Load(path)
	if err != nil {
		return nil, err
	}
	if len(diags) > 0 {
		printDiagsText(out, diags)
		return nil, fmt.Errorf("%s: %d issue(s) found", path, len(diags))
	}
	return b, nil
}
