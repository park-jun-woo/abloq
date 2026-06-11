//ff:func feature=cli type=output control=sequence
//ff:what 매칭 결과 출력 검증 — 전부 출현이면 요약만, 미출현 있으면 id: text 목록 추가
package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/insight"
)

func TestPrintInsightMatch(t *testing.T) {
	var out bytes.Buffer
	printInsightMatch(&out, insight.Result{Found: []string{"a", "b"}}, 2)
	if got := out.String(); got != "anchored claims: 2/2\n" {
		t.Errorf("want summary only when nothing missing, got %q", got)
	}
	out.Reset()
	printInsightMatch(&out, insight.Result{Found: []string{"a"}, Missing: []insight.Claim{{ID: "m", Text: "miss"}}}, 2)
	got := out.String()
	if !strings.Contains(got, "anchored claims: 1/2") || !strings.Contains(got, "  m: miss") {
		t.Errorf("want summary and missing list, got %q", got)
	}
}
