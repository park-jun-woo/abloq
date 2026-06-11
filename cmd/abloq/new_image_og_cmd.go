//ff:func feature=cli type=command control=sequence
//ff:what "abloq image og <slug> <title>" cobra 명령 생성 — 텍스트 옵션(--brand/--font/--out) + provider 2축(--provider/--model/--overlay × --variant/--all-variants/--count)
package main

import "github.com/spf13/cobra"

// newImageOGCmd builds the image og subcommand. The default (no flags) stays
// the deterministic local card; AI providers and multi-candidate sampling are
// pure opt-in.
func newImageOGCmd() *cobra.Command {
	var opts imageOGOpts
	cmd := &cobra.Command{
		Use:   "og <slug> <title>",
		Short: "Generate a 1200x630 OG image at static/images/{slug}.webp",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Slug, opts.Title = args[0], args[1]
			opts.ProviderSet = cmd.Flags().Changed("provider")
			opts.ModelSet = cmd.Flags().Changed("model")
			opts.OverlaySet = cmd.Flags().Changed("overlay")
			return runImageOG(cmd.OutOrStdout(), opts)
		},
	}
	cmd.Flags().StringVar(&opts.Brand, "brand", "", "brand line under the title (accent color)")
	cmd.Flags().StringVar(&opts.FontPath, "font", "", "TTF/OTF path (default: embedded Go Bold, latin only)")
	cmd.Flags().StringVar(&opts.OutDir, "out", "static/images", "output directory")
	cmd.Flags().StringVar(&opts.Provider, "provider", "", "OG provider: local | gemini (flag > blog.yaml image.og > local)")
	cmd.Flags().StringVar(&opts.Model, "model", "", "provider model id (empty = provider default)")
	cmd.Flags().BoolVar(&opts.Overlay, "overlay", false, "composite the deterministic title/brand over the AI background")
	cmd.Flags().StringVar(&opts.VariantList, "variant", "", "comma-separated variant names from blog.yaml image.og.variants (drafts to files/og/)")
	cmd.Flags().BoolVar(&opts.AllVariants, "all-variants", false, "generate every declared variant (drafts to files/og/)")
	cmd.Flags().IntVar(&opts.Count, "count", 1, "samples per variant (>1 drafts to files/og/)")
	return cmd
}
