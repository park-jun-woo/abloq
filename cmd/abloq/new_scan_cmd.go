//ff:func feature=cli type=command control=sequence
//ff:what "abloq scan" 부모 명령 생성 — 스캐너 서브커맨드(freshness·evidence·cluster) 등록
package main

import "github.com/spf13/cobra"

// newScanCmd builds the scan command group. Every backend scanner module is
// also runnable here without the backend (design §3.3 — the backend only adds
// schedule and state).
func newScanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Content scanners: detect quest candidates and write the local queue",
	}
	cmd.AddCommand(newScanFreshnessCmd())
	cmd.AddCommand(newScanEvidenceCmd())
	cmd.AddCommand(newScanClusterCmd())
	return cmd
}
