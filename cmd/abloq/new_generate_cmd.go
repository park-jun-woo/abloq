//ff:func feature=cli type=command control=sequence
//ff:what "abloq generate [dir]" cobra 명령 생성 — dir 기본값은 현재 디렉토리
package main

import "github.com/spf13/cobra"

// newGenerateCmd builds the generate subcommand.
func newGenerateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generate [dir]",
		Short: "Generate derived files (hugo.toml, robots.txt, llms.txt, jsonld.json) from blog.yaml",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := "."
			if len(args) == 1 {
				dir = args[0]
			}
			return runGenerate(cmd.OutOrStdout(), dir)
		},
	}
}
