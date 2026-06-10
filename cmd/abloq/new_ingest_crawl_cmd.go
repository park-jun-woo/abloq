//ff:func feature=cli type=command control=sequence topic=crawl
//ff:what "abloq ingest crawl --source <dir|s3://bucket/prefix> --repo <dir>" cobra 명령 생성 — CF 로그 단발 분석 (DB 없이)
package main

import "github.com/spf13/cobra"

// newIngestCrawlCmd builds the crawl ingest subcommand: a stateless one-shot
// aggregation of CloudFront logs — no cursor, no database, prints the
// aggregate plus the unfiltered raw bot counters (the analyze-stats.py
// comparison point).
func newIngestCrawlCmd() *cobra.Command {
	var source, repo string
	cmd := &cobra.Command{
		Use:   "crawl --source <dir|s3://bucket/prefix> [--repo <dir>]",
		Short: "Aggregate AI-bot crawl hits from CloudFront logs (one-shot, no DB)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runIngestCrawl(cmd.OutOrStdout(), source, repo)
		},
	}
	cmd.Flags().StringVar(&source, "source", "", "log source: a local directory or s3://bucket/prefix (AWS_* env credentials)")
	cmd.Flags().StringVar(&repo, "repo", ".", "blog repository root (URL reverse map)")
	_ = cmd.MarkFlagRequired("source")
	return cmd
}
