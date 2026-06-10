//ff:func feature=cli type=command control=sequence
//ff:what validate 명령 케이스 하나를 실행해 에러 여부와 출력 내용을 검증
package main

import (
	"bytes"
	"strings"
	"testing"
)

func checkValidateCmdCase(t *testing.T, args []string, wantErr bool, wantOut string) {
	t.Helper()
	cmd := newValidateCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs(args)
	err := cmd.Execute()
	if wantErr && err == nil {
		t.Fatalf("want error, got nil\noutput: %s", out.String())
	}
	if !wantErr && err != nil {
		t.Fatalf("Execute: %v\noutput: %s", err, out.String())
	}
	if wantOut != "" && !strings.Contains(out.String(), wantOut) {
		t.Errorf("want output containing %q, got %q", wantOut, out.String())
	}
}
