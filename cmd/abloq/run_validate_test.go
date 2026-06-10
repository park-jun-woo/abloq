//ff:func feature=cli type=command control=sequence
//ff:what abloq validate가 12언어 골든 예제 디렉토리에서 OK로 통과하는지 검증
package main

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunValidate(t *testing.T) {
	var out bytes.Buffer
	dir := filepath.Join("..", "..", "pkg", "blogyaml", "testdata", "valid")
	if err := runValidate(&out, dir, false); err != nil {
		t.Fatalf("runValidate: %v\noutput: %s", err, out.String())
	}
	if !strings.Contains(out.String(), "OK") {
		t.Errorf("want OK in output, got %q", out.String())
	}
}
