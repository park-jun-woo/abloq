//ff:func feature=cli type=command control=sequence
//ff:what dir/blog.yaml을 표면화 — 파일이 없으면 (nil blog, 영값 image)=local 완전 하위호환, 있으면 검증된 *Blog와 image 블록을 함께 반환(base 언어·섹션은 summary 결선이 공유)
package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// loadImageOG surfaces the validated blog and its image.og declaration in one
// load. A missing blog.yaml is not an error — `abloq image og` keeps working
// outside a blog root exactly as before (local card), and returns a nil blog so
// summary resolution knows to skip (no IndexEntries(root, nil) dereference). An
// existing blog.yaml must validate, like every other command consuming the SSOT.
// The returned *Blog carries the base language and sections the summary
// resolver needs, so the blog is loaded once and shared.
func loadImageOG(out io.Writer, dir string) (*blogyaml.Blog, blogyaml.Image, error) {
	path := filepath.Join(dir, "blog.yaml")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, blogyaml.Image{}, nil
	}
	b, err := loadValidBlog(out, dir)
	if err != nil {
		return nil, blogyaml.Image{}, err
	}
	return b, b.Image, nil
}
