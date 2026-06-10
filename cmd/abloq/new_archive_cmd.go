//ff:func feature=cli type=command control=sequence
//ff:what "abloq archive <url>" cobra 명령 생성 — Wayback/IndexNow/GSC 3종 제출, 백엔드 아카이버와 동일 pkg/archive 공유
package main

import "github.com/spf13/cobra"

// newArchiveCmd builds the archive subcommand. It runs the same submission
// code path as the abloqd processor (pkg/archive) without needing the
// backend — credentials and base URLs come from the environment.
func newArchiveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "archive <url>",
		Short: "Submit a URL to Wayback, IndexNow and the GSC Indexing API (env credentials)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runArchive(cmd.OutOrStdout(), args[0])
		},
	}
}
