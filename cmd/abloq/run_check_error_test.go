//ff:func feature=cli type=command control=sequence topic=drift
//ff:what runCheck가 blog.yaml 누락 디렉토리에서 IO 에러를 반환하는지 검증
package main

import (
	"bytes"
	"testing"
)

func TestRunCheckError(t *testing.T) {
	var out bytes.Buffer
	if err := runCheck(&out, t.TempDir()); err == nil {
		t.Errorf("want error for dir without blog.yaml, got nil")
	}
}
