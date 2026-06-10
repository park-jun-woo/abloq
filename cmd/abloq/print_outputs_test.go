//ff:func feature=cli type=output control=sequence
//ff:what printOutputs가 dir 기준 파생물 경로를 한 줄씩 출력하는지 검증
package main

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/gen"
)

func TestPrintOutputs(t *testing.T) {
	var out bytes.Buffer
	outs := []gen.Output{{Path: "hugo.toml"}, {Path: "static/robots.txt"}}
	printOutputs(&out, "blog", outs)
	want := filepath.Join("blog", "hugo.toml") + "\n" + filepath.Join("blog", "static/robots.txt") + "\n"
	if out.String() != want {
		t.Errorf("printOutputs = %q, want %q", out.String(), want)
	}
}
