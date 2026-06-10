//ff:func feature=cli type=command control=sequence topic=diagnostics
//ff:what runGate가 blog.yaml 누락 디렉토리에서 IO 에러를 반환하는지 검증
package main

import (
	"bytes"
	"testing"
)

func TestRunGateError(t *testing.T) {
	var out bytes.Buffer
	if err := runGate(&out, t.TempDir(), "", false, false); err == nil {
		t.Errorf("want error for dir without blog.yaml, got nil")
	}
}
