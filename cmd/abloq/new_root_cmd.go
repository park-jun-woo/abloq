//ff:func feature=cli type=command control=sequence
//ff:what abloq 루트 cobra 명령 생성 — 서브커맨드(validate/generate/check) 등록
package main

import "github.com/spf13/cobra"

// newRootCmd builds the abloq root command.
func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "abloq",
		Short:         "abloq — Agentic blog Quest framework CLI",
		Long:          "abloq drives an agent-operable blog from a single blog.yaml SSOT.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.AddCommand(newValidateCmd())
	cmd.AddCommand(newGenerateCmd())
	cmd.AddCommand(newCheckCmd())
	return cmd
}
