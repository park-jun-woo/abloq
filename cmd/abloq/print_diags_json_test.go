//ff:func feature=cli type=output control=sequence topic=diagnostics
//ff:what printDiagsJSON이 nil 진단을 []로, 진단 목록을 JSON 배열로 출력하고 쓰기 실패를 에러로 반환하는지 검증
package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestPrintDiagsJSON(t *testing.T) {
	t.Run("nil diags emit empty array", func(t *testing.T) {
		var out bytes.Buffer
		if err := printDiagsJSON(&out, nil); err != nil {
			t.Fatalf("printDiagsJSON: %v", err)
		}
		if got := out.String(); got != "[]\n" {
			t.Errorf("want %q, got %q", "[]\n", got)
		}
	})
	t.Run("diags round-trip as JSON", func(t *testing.T) {
		var out bytes.Buffer
		in := []blogyaml.Diagnostic{{File: "blog.yaml", Line: 3, Rule: "lang-bcp47", Message: "bad"}}
		if err := printDiagsJSON(&out, in); err != nil {
			t.Fatalf("printDiagsJSON: %v", err)
		}
		var got []blogyaml.Diagnostic
		if err := json.Unmarshal(out.Bytes(), &got); err != nil {
			t.Fatalf("output is not valid JSON: %v\noutput: %s", err, out.String())
		}
		if len(got) != 1 || got[0] != in[0] {
			t.Errorf("want %v, got %v", in, got)
		}
	})
	t.Run("write failure returns error", func(t *testing.T) {
		if err := printDiagsJSON(errWriter{}, nil); err == nil {
			t.Errorf("want error from failing writer, got nil")
		}
	})
}
