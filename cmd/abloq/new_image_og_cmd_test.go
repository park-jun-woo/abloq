//ff:func feature=cli type=command control=sequence
//ff:what image og 명령이 slug/제목 인자와 --out 플래그로 OG WebP를 생성하는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestNewImageOGCmd(t *testing.T) {
	dir := t.TempDir()
	cmd := newImageOGCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetArgs([]string{"hello", "Hello abloq", "--out", dir})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("image og: %v\noutput: %s", err, out.String())
	}
	if _, err := os.Stat(filepath.Join(dir, "hello.webp")); err != nil {
		t.Errorf("og image missing: %v", err)
	}
}
