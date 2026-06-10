//ff:func feature=cli type=command control=sequence
//ff:what "abloq claudemd [dir]" cobra 명령 생성 — blog.yaml에서 CLAUDE.md 운영 매뉴얼 재생성
package main

import "github.com/spf13/cobra"

// newClaudeMDCmd builds the claudemd subcommand.
func newClaudeMDCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "claudemd [dir]",
		Short: "Regenerate CLAUDE.md (agent ops manual) from blog.yaml",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := "."
			if len(args) == 1 {
				dir = args[0]
			}
			return runClaudeMD(cmd.OutOrStdout(), dir)
		},
	}
}
