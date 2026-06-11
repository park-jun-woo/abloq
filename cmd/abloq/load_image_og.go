//ff:func feature=cli type=command control=sequence
//ff:what dir/blog.yaml의 image 블록을 읽기 — 파일이 없으면 영값(=local, 완전 하위호환), 있으면 유효한 SSOT여야 한다
package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// loadImageOG fetches the image.og declaration. A missing blog.yaml is not an
// error — `abloq image og` keeps working outside a blog root exactly as
// before (local card). An existing blog.yaml must validate, like every other
// command that consumes the SSOT.
func loadImageOG(out io.Writer, dir string) (blogyaml.Image, error) {
	path := filepath.Join(dir, "blog.yaml")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return blogyaml.Image{}, nil
	}
	b, err := loadValidBlog(out, dir)
	if err != nil {
		return blogyaml.Image{}, err
	}
	return b.Image, nil
}
