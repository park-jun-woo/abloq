//ff:func feature=cli type=command control=sequence topic=diagnostics
//ff:what runValidate의 에러 경로 — 파일 없음 IO 에러, text 모드 진단 출력, JSON 쓰기 실패를 검증
package main

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunValidateError(t *testing.T) {
	invalidDir := filepath.Join("..", "..", "pkg", "blogyaml", "testdata", "invalid-lang")

	t.Run("missing blog.yaml returns IO error", func(t *testing.T) {
		var out bytes.Buffer
		if err := runValidate(&out, t.TempDir(), false); err == nil {
			t.Fatalf("want IO error for missing blog.yaml, got nil\noutput: %s", out.String())
		}
	})
	t.Run("text mode prints diagnostics and fails", func(t *testing.T) {
		var out bytes.Buffer
		err := runValidate(&out, invalidDir, false)
		if err == nil {
			t.Fatalf("want error for invalid blog.yaml, got nil\noutput: %s", out.String())
		}
		if !strings.Contains(err.Error(), "1 issue(s) found") {
			t.Errorf("want error mentioning issue count, got %q", err.Error())
		}
		if !strings.Contains(out.String(), "[lang-bcp47]") {
			t.Errorf("want lang-bcp47 diagnostic line in output, got %q", out.String())
		}
	})
	t.Run("json write failure returns error", func(t *testing.T) {
		if err := runValidate(errWriter{}, invalidDir, true); err == nil {
			t.Errorf("want error from failing writer, got nil")
		}
	})
}
