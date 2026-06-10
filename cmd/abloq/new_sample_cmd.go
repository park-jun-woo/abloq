//ff:func feature=cli type=command control=sequence topic=citation
//ff:what "abloq sample" 부모 명령 생성 — 샘플링 서브커맨드(citations) 등록
package main

import "github.com/spf13/cobra"

// newSampleCmd builds the sample command group. Sampling records trends
// only — its output feeds no gate or verdict (§6.3 non-determinism
// quarantine).
func newSampleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sample",
		Short: "Visibility sampling: probe AI engines for own-domain citations",
	}
	cmd.AddCommand(newSampleCitationsCmd())
	return cmd
}
