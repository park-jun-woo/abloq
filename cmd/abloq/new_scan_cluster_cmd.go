//ff:func feature=cli type=command control=sequence
//ff:what "abloq scan cluster [dir]" cobra 명령 생성 — 태그·내부링크 그래프 위반 검출 + 연결 후보 제안 큐 적재
package main

import "github.com/spf13/cobra"

// newScanClusterCmd builds the cluster scan subcommand.
func newScanClusterCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "cluster [dir]",
		Short: "Detect cluster violations (tags, internal links) and write quests/queue/ files with link candidates",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := "."
			if len(args) == 1 {
				dir = args[0]
			}
			return runScanCluster(cmd.OutOrStdout(), dir)
		},
	}
}
