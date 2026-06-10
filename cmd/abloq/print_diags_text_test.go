//ff:func feature=cli type=output control=sequence topic=diagnostics
//ff:what printDiagsText가 진단마다 "파일:라인 [룰ID] 메시지" 한 줄을 출력하는지 검증
package main

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestPrintDiagsText(t *testing.T) {
	var out bytes.Buffer
	printDiagsText(&out, []blogyaml.Diagnostic{
		{File: "blog.yaml", Line: 3, Rule: "lang-bcp47", Message: "bad code"},
		{File: "blog.yaml", Line: 9, Rule: "sections-empty", Message: "no sections"},
	})
	want := "blog.yaml:3 [lang-bcp47] bad code\nblog.yaml:9 [sections-empty] no sections\n"
	if got := out.String(); got != want {
		t.Errorf("want %q, got %q", want, got)
	}

	out.Reset()
	printDiagsText(&out, nil)
	if got := out.String(); got != "" {
		t.Errorf("want empty output for no diagnostics, got %q", got)
	}
}
