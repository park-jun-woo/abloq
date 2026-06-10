//ff:func feature=cli type=command control=sequence topic=drift
//ff:what "abloq check [dir]" cobra 명령 생성 — 파생물 드리프트 검증, dir 기본값은 현재 디렉토리
package main

import "github.com/spf13/cobra"

// newCheckCmd builds the check subcommand.
func newCheckCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "check [dir]",
		Short: "Check derived files against a fresh regeneration (exit 1 on drift)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := "."
			if len(args) == 1 {
				dir = args[0]
			}
			return runCheck(cmd.OutOrStdout(), dir)
		},
	}
}
