//ff:func feature=cli type=command control=sequence
//ff:what "abloq insight match <insight.yaml> <article>" cobra 명령 생성 — 인자 2개 고정
package main

import "github.com/spf13/cobra"

// newInsightMatchCmd builds the insight match subcommand.
func newInsightMatchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "match <insight.yaml> <article>",
		Short: "Screen insight claims against an article body (REVIEW aid, default language only)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInsightMatch(cmd.OutOrStdout(), args[0], args[1])
		},
	}
}
