//ff:func feature=cli type=command control=sequence
//ff:what "abloq image og <slug> <title>" cobra 명령 생성 — 1200×630 OG WebP, --brand/--font/--out 플래그
package main

import "github.com/spf13/cobra"

// newImageOGCmd builds the image og subcommand.
func newImageOGCmd() *cobra.Command {
	var brand, fontPath, outDir string
	cmd := &cobra.Command{
		Use:   "og <slug> <title>",
		Short: "Generate a 1200x630 OG image at static/images/{slug}.webp",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runImageOG(cmd.OutOrStdout(), args[0], args[1], brand, fontPath, outDir)
		},
	}
	cmd.Flags().StringVar(&brand, "brand", "", "brand line under the title (accent color)")
	cmd.Flags().StringVar(&fontPath, "font", "", "TTF/OTF path (default: embedded Go Bold, latin only)")
	cmd.Flags().StringVar(&outDir, "out", "static/images", "output directory")
	return cmd
}
