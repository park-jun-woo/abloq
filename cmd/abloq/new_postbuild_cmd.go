//ff:func feature=cli type=command control=sequence
//ff:what "abloq postbuild" 부모 명령 생성 — 빌드 후처리 서브커맨드(md) 등록
package main

import "github.com/spf13/cobra"

// newPostbuildCmd builds the postbuild command group.
func newPostbuildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "postbuild",
		Short: "Post-build steps for the generated site (run after hugo)",
	}
	cmd.AddCommand(newPostbuildMDCmd())
	return cmd
}
