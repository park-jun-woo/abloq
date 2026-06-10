//ff:func feature=cli type=command control=sequence
//ff:what image convert 명령이 src 인자와 --slug/--out 플래그로 WebP를 생성하는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestNewImageConvertCmd(t *testing.T) {
	dir := t.TempDir()
	src := writePNGFixture(t, dir)
	cmd := newImageConvertCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetArgs([]string{src, "--slug", "pic", "--out", dir})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("image convert: %v\noutput: %s", err, out.String())
	}
	if _, err := os.Stat(filepath.Join(dir, "pic.webp")); err != nil {
		t.Errorf("pic.webp missing: %v", err)
	}
}
