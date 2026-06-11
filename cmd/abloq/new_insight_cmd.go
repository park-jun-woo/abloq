//ff:func feature=cli type=command control=sequence
//ff:what "abloq insight" 부모 명령 생성 — match 서브커맨드 등록
package main

import "github.com/spf13/cobra"

// newInsightCmd builds the insight command group.
func newInsightCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insight",
		Short: "Insight spec tools: screen insight.yaml claims against an article",
	}
	cmd.AddCommand(newInsightMatchCmd())
	return cmd
}
