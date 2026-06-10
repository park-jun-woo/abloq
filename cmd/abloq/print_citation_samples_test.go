//ff:func feature=cli type=output control=sequence topic=citation
//ff:what printCitationSamples가 샘플 행과 엔진 수·budget·샘플 합계 한 줄을 출력하는지 검증
package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/visibility/citation"
)

func TestPrintCitationSamples(t *testing.T) {
	var out bytes.Buffer
	printCitationSamples(&out, 2, 3, []citation.Sample{
		{CitationQueriesID: 7, Engine: "openai", Cited: true, Evidence: `{"matched":["https://t.example.com/a/"]}`, ExtractorVersion: "v1"},
	})
	got := out.String()
	if !strings.Contains(got, "openai") || !strings.Contains(got, `{"matched":["https://t.example.com/a/"]}`) {
		t.Errorf("sample row missing:\n%s", got)
	}
	if !strings.Contains(got, "sample: 2 engine(s), budget 3, 1 sample(s) [extractor v1]") {
		t.Errorf("summary missing:\n%s", got)
	}
}
