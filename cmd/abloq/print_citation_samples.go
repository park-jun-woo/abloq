//ff:func feature=cli type=output control=iteration dimension=1 topic=citation
//ff:what 인용 샘플 출력 — 샘플 행(엔진 질의id cited 근거)과 엔진 수·budget·샘플 합계 한 줄
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/abloq/pkg/visibility/citation"
)

// printCitationSamples prints one sampling round: the per-(engine, query)
// records and a one-line summary (budget 0 prints an explicit no-op).
func printCitationSamples(out io.Writer, engines, budget int, samples []citation.Sample) {
	fmt.Fprintln(out, "citation samples (engine query_id cited evidence):")
	for _, s := range samples {
		fmt.Fprintf(out, "  %-12s %4d %-5t %s\n", s.Engine, s.CitationQueriesID, s.Cited, s.Evidence)
	}
	fmt.Fprintf(out, "sample: %d engine(s), budget %d, %d sample(s) [extractor %s]\n",
		engines, budget, len(samples), citation.ExtractorVersion)
}
