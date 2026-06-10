//ff:func feature=cli type=command control=iteration dimension=1
//ff:what validate 명령이 dir 인자(기본값 현재 디렉토리)와 --json 플래그를 받아 실행되는지 검증
package main

import (
	"path/filepath"
	"testing"
)

func TestNewValidateCmd(t *testing.T) {
	validDir := filepath.Join("..", "..", "pkg", "blogyaml", "testdata", "valid")
	cases := []struct {
		name    string
		args    []string
		wantErr bool
		wantOut string
	}{
		{"explicit dir", []string{validDir}, false, "OK"},
		{"explicit dir json", []string{"--json", validDir}, false, "[]"},
		{"default dir has no blog.yaml", []string{}, true, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkValidateCmdCase(t, tc.args, tc.wantErr, tc.wantOut) })
	}
}
