//ff:func feature=cli type=command control=sequence
//ff:what "abloq image convert <src>" cobra 명령 생성 — WebP 변환, --slug/--max-width/--out 플래그
package main

import "github.com/spf13/cobra"

// newImageConvertCmd builds the image convert subcommand.
func newImageConvertCmd() *cobra.Command {
	var slug, outDir string
	var maxWidth int
	cmd := &cobra.Command{
		Use:   "convert <src>",
		Short: "Convert an image to WebP at static/images/{slug}.webp (white-flattened, resized)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runImageConvert(cmd.OutOrStdout(), args[0], slug, outDir, maxWidth)
		},
	}
	cmd.Flags().StringVar(&slug, "slug", "", "output name (default: source filename)")
	cmd.Flags().StringVar(&outDir, "out", "static/images", "output directory")
	cmd.Flags().IntVar(&maxWidth, "max-width", 1400, "downscale wider images to this width (0 = keep)")
	return cmd
}
