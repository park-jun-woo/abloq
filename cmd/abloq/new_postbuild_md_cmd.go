//ff:func feature=cli type=command control=sequence
//ff:what "abloq postbuild md [dir]" cobra 명령 생성 — 글마다 노이즈 제로 .md를 public/에 병행 산출
package main

import "github.com/spf13/cobra"

// newPostbuildMDCmd builds the postbuild md subcommand.
func newPostbuildMDCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "md [dir]",
		Short: "Serve a clean .md beside every built article (AI context format)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := "."
			if len(args) == 1 {
				dir = args[0]
			}
			return runPostbuildMD(cmd.OutOrStdout(), dir)
		},
	}
}
