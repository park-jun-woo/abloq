//ff:func feature=cli type=command control=sequence
//ff:what "abloq scan evidence [dir]" cobra 명령 생성 — 무출처 수치 주장 큐 적재 + 인용 link rot 1회 점검 보고
package main

import "github.com/spf13/cobra"

// newScanEvidenceCmd builds the evidence scan subcommand.
func newScanEvidenceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "evidence [dir]",
		Short: "Detect unsourced numeric claims (quests/queue/ files) and report citation link rot",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := "."
			if len(args) == 1 {
				dir = args[0]
			}
			return runScanEvidence(cmd.OutOrStdout(), dir)
		},
	}
}
