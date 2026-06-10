//ff:func feature=cli type=command control=sequence
//ff:what "abloq image" 부모 명령 생성 — 이미지 도구 서브커맨드(og/convert) 등록
package main

import "github.com/spf13/cobra"

// newImageCmd builds the image command group.
func newImageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image",
		Short: "Image tools: OG card generation and WebP conversion (pure Go)",
	}
	cmd.AddCommand(newImageOGCmd())
	cmd.AddCommand(newImageConvertCmd())
	return cmd
}
