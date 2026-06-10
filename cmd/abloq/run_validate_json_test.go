//ff:func feature=cli type=output control=sequence topic=diagnostics
//ff:what abloq validate --json이 오류 예제에서 에러를 반환하고 유효한 JSON 진단 배열을 출력하는지 검증
package main

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestRunValidateJSON(t *testing.T) {
	var out bytes.Buffer
	dir := filepath.Join("..", "..", "pkg", "blogyaml", "testdata", "invalid-lang")
	err := runValidate(&out, dir, true)
	if err == nil {
		t.Fatalf("want error for invalid blog.yaml, got nil\noutput: %s", out.String())
	}
	var diags []blogyaml.Diagnostic
	if jerr := json.Unmarshal(out.Bytes(), &diags); jerr != nil {
		t.Fatalf("output is not valid JSON: %v\noutput: %s", jerr, out.String())
	}
	if len(diags) != 1 || diags[0].Rule != "lang-bcp47" {
		t.Errorf("want 1 lang-bcp47 diagnostic, got %v", diags)
	}
}
