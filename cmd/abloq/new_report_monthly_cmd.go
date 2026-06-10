//ff:func feature=cli type=command control=sequence topic=report
//ff:what "abloq report monthly --ym <YYYY-MM> [--source <logs>] [dir]" cobra 명령 생성 — 무상태 부분 리포트 (DB 없이)
package main

import "github.com/spf13/cobra"

// newReportMonthlyCmd builds the monthly report subcommand: a stateless
// partial report. The crawl layer comes from CloudFront logs (--source) and
// the repository alone; the citation layer is impossible without the stored
// sample time series and GSC stays out of scope — the output says so.
func newReportMonthlyCmd() *cobra.Command {
	var ym, source string
	cmd := &cobra.Command{
		Use:   "monthly --ym <YYYY-MM> [--source <dir|s3://bucket/prefix>] [dir]",
		Short: "Render a partial monthly visibility report from logs and the repository (no DB)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := "."
			if len(args) == 1 {
				dir = args[0]
			}
			return runReportMonthly(cmd.OutOrStdout(), dir, ym, source)
		},
	}
	cmd.Flags().StringVar(&ym, "ym", "", "report month YYYY-MM (default: the last closed month, UTC)")
	cmd.Flags().StringVar(&source, "source", "", "CloudFront log source: a local directory or s3://bucket/prefix (omit for an empty crawl layer)")
	return cmd
}
