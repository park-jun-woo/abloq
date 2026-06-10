//ff:func feature=cli type=command control=sequence
//ff:what "abloq validate [dir]" cobra 명령 생성 — --json 플래그, dir 기본값은 현재 디렉토리
package main

import "github.com/spf13/cobra"

// newValidateCmd builds the validate subcommand.
func newValidateCmd() *cobra.Command {
	var jsonOut bool
	cmd := &cobra.Command{
		Use:   "validate [dir]",
		Short: "Validate blog.yaml in dir (default: current directory)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := "."
			if len(args) == 1 {
				dir = args[0]
			}
			return runValidate(cmd.OutOrStdout(), dir, jsonOut)
		},
	}
	cmd.Flags().BoolVar(&jsonOut, "json", false, "emit diagnostics as JSON")
	return cmd
}
