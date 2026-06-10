//ff:func feature=init type=command control=sequence
//ff:what "abloq init <dir>" cobra 명령 생성 — 제목/언어/섹션 플래그, --interactive 외에는 비대화형(에이전트 경로)
package main

import (
	"strings"

	"github.com/spf13/cobra"
)

// newInitCmd builds the init subcommand. Prompts are opt-in (--interactive);
// the default path takes everything from flags so agents can script it.
func newInitCmd() *cobra.Command {
	o := initOpts{}
	var langs, sections string
	cmd := &cobra.Command{
		Use:   "init <dir>",
		Short: "Scaffold a new abloq blog (blog.yaml + template + CLAUDE.md, gate-clean)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			o.Languages = strings.Split(langs, ",")
			o.Sections = strings.Split(sections, ",")
			return runInit(cmd.OutOrStdout(), cmd.InOrStdin(), args[0], o)
		},
	}
	cmd.Flags().StringVar(&o.Title, "title", "My Blog", "site title")
	cmd.Flags().StringVar(&o.BaseURL, "baseurl", "https://example.com", "absolute http(s) base URL")
	cmd.Flags().StringVar(&o.Author, "author", "Author", "author name")
	cmd.Flags().StringVar(&langs, "languages", "en", "comma-separated BCP-47 codes; first = default language")
	cmd.Flags().StringVar(&sections, "sections", "posts", "comma-separated content sections")
	cmd.Flags().BoolVar(&o.Interactive, "interactive", false, "prompt for answers instead of flags/defaults")
	return cmd
}
