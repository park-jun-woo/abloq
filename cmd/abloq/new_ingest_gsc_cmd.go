//ff:func feature=cli type=command control=sequence topic=gsc
//ff:what "abloq ingest gsc [--site <property>] [--days N] [--repo <dir>]" cobra 명령 생성 — Search Analytics 단발 조회 (DB 없이)
package main

import "github.com/spf13/cobra"

// newIngestGSCCmd builds the GSC ingest subcommand: a stateless one-shot
// Search Analytics readout of the last N closed days — no cursor, no
// database. Credentials ride in env only (GSC_SA_JSON / GSC_SA_JSON_PATH).
func newIngestGSCCmd() *cobra.Command {
	var site, repo string
	var days int
	cmd := &cobra.Command{
		Use:   "gsc [--site <property>] [--days <n>] [--repo <dir>]",
		Short: "Fetch recent Search Analytics rows from GSC (one-shot, no DB)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runIngestGSC(cmd.OutOrStdout(), site, repo, days)
		},
	}
	cmd.Flags().StringVar(&site, "site", "", "Search Console property (default: blog.yaml baseURL of --repo)")
	cmd.Flags().StringVar(&repo, "repo", ".", "blog repository root (baseURL fallback)")
	cmd.Flags().IntVar(&days, "days", 7, "closed days to fetch (today-margin backwards)")
	return cmd
}
