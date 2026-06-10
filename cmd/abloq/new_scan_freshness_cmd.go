//ff:func feature=cli type=command control=sequence
//ff:what "abloq scan freshness [dir]" cobra 명령 생성 — freshness_days 초과 글 검출, quests/queue/에 큐 파일 직접 기록
package main

import "github.com/spf13/cobra"

// newScanFreshnessCmd builds the freshness scan subcommand.
func newScanFreshnessCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "freshness [dir]",
		Short: "Detect stale articles (geo.freshness_days) and write quests/queue/ files",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := "."
			if len(args) == 1 {
				dir = args[0]
			}
			return runScanFreshness(cmd.OutOrStdout(), dir)
		},
	}
}
