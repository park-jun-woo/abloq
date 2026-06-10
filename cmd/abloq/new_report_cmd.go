//ff:func feature=cli type=command control=sequence topic=report
//ff:what "abloq report" 부모 명령 생성 — 리포트 서브커맨드(monthly) 등록
package main

import "github.com/spf13/cobra"

// newReportCmd builds the report command group. The visibility report is
// also producible without the backend (design §3.3) — within what a
// stateless run can honestly measure.
func newReportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "Visibility reports: crawl x index x citation joined per article",
	}
	cmd.AddCommand(newReportMonthlyCmd())
	return cmd
}
