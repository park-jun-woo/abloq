//ff:func feature=cli type=command control=sequence topic=citation
//ff:what citations 샘플링 실행 본체 — 질의 파일 + blog.yaml(도메인·citation_budget)로 env 키 엔진들을 1회전 실행해 출력 (DB 없이)
//ff:why 입력은 인자·출력은 stdout(ingest crawl 선례): 적재는 백엔드 소관. budget 0이면 blog.yaml이 샘플링을 끈 것 — no-op도 그대로 보여준다 (Phase013)
package main

import (
	"io"

	"github.com/park-jun-woo/abloq/pkg/visibility/citation"
)

// runSampleCitations runs one citation sampling round without the backend:
// queries from the --queries file, domain and budget from the blog.yaml of
// repo, engines from the API-key environment. Results go to out only.
func runSampleCitations(out io.Writer, queriesPath, repo string) error {
	queries, err := loadQueriesFile(queriesPath)
	if err != nil {
		return err
	}
	b, err := loadValidBlog(out, repo)
	if err != nil {
		return err
	}
	engines := citation.EnginesFromEnv()
	samples := citation.Run(engines, queries, b.Geo.CitationBudget, b.Site.BaseURL, citation.IntervalFromEnv())
	printCitationSamples(out, len(engines), b.Geo.CitationBudget, samples)
	return nil
}
