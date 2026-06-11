//ff:func feature=cli type=command control=sequence
//ff:what abloq 루트 cobra 명령 생성 — 서브커맨드(validate/generate/check/gate/init/postbuild/image/claudemd/archive/scan/ingest/sample/report/insight/quest) 등록
package main

import "github.com/spf13/cobra"

// newRootCmd builds the abloq root command.
func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "abloq",
		Version:       "0.1.0",
		Short:         "abloq — Agentic blog Quest framework CLI",
		Long:          "abloq drives an agent-operable blog from a single blog.yaml SSOT.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.AddCommand(newValidateCmd())
	cmd.AddCommand(newGenerateCmd())
	cmd.AddCommand(newCheckCmd())
	cmd.AddCommand(newGateCmd())
	cmd.AddCommand(newInitCmd())
	cmd.AddCommand(newPostbuildCmd())
	cmd.AddCommand(newImageCmd())
	cmd.AddCommand(newClaudeMDCmd())
	cmd.AddCommand(newArchiveCmd())
	cmd.AddCommand(newScanCmd())
	cmd.AddCommand(newIngestCmd())
	cmd.AddCommand(newSampleCmd())
	cmd.AddCommand(newReportCmd())
	cmd.AddCommand(newInsightCmd())
	cmd.AddCommand(newQuestCmd())
	return cmd
}
