//ff:func feature=cli type=command control=sequence topic=citation
//ff:what "abloq sample citations --queries <yaml|json> [--repo <dir>]" cobra 명령 생성 — 인용 샘플링 단발 실행 (DB 없이)
package main

import "github.com/spf13/cobra"

// newSampleCitationsCmd builds the citation sampling subcommand: queries
// come from the argument file (the DB query set is backend territory),
// engines from the API-key environment, domain and budget from blog.yaml.
func newSampleCitationsCmd() *cobra.Command {
	var queries, repo string
	cmd := &cobra.Command{
		Use:   "citations --queries <yaml|json> [--repo <dir>]",
		Short: "Run one citation sampling round over a query file (one-shot, no DB)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSampleCitations(cmd.OutOrStdout(), queries, repo)
		},
	}
	cmd.Flags().StringVar(&queries, "queries", "", "query file: a YAML/JSON list of {id, query_text}")
	cmd.Flags().StringVar(&repo, "repo", ".", "blog repository root (domain + citation_budget)")
	_ = cmd.MarkFlagRequired("queries")
	return cmd
}
