//ff:func feature=cli type=command control=sequence topic=crawl
//ff:what "abloq ingest" 부모 명령 생성 — 수집 서브커맨드(crawl) 등록
package main

import "github.com/spf13/cobra"

// newIngestCmd builds the ingest command group. Like the scanners, every
// backend ingest module is also runnable here without the backend (design
// §3.3 — the backend only adds schedule and state).
func newIngestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingest",
		Short: "Visibility ingests: aggregate external signals without the backend",
	}
	cmd.AddCommand(newIngestCrawlCmd())
	return cmd
}
