//ff:func feature=cli type=command control=sequence
//ff:what "abloq gate [--rule <id>] [--json] [dir]" cobra 명령 생성 — 구조 게이트 룰셋 실행
package main

import "github.com/spf13/cobra"

// newGateCmd builds the gate subcommand.
func newGateCmd() *cobra.Command {
	var ruleID string
	var jsonOut bool
	cmd := &cobra.Command{
		Use:   "gate [dir]",
		Short: "Run the structure gate rules on the blog's articles (exit 1 on violations)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := "."
			if len(args) == 1 {
				dir = args[0]
			}
			return runGate(cmd.OutOrStdout(), dir, ruleID, jsonOut)
		},
	}
	cmd.Flags().StringVar(&ruleID, "rule", "", "run a single rule by id (default: all rules)")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "emit diagnostics as JSON")
	return cmd
}
