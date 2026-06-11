//ff:func feature=cli type=command control=sequence
//ff:what "abloq quest" 부모 명령 생성 — reins NewQuestCmd("writing")를 마운트 (scan/next/submit/status/export/rules)
package main

import (
	rcli "github.com/park-jun-woo/reins/pkg/cli"
	"github.com/spf13/cobra"

	"github.com/park-jun-woo/abloq/pkg/quests/writing"
)

// newQuestCmd builds the quest command group. Each quest is a reins consumer:
// reins supplies the ratchet/commands, the Definition supplies the domain.
func newQuestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quest",
		Short: "Agent quests (reins-gated): writing",
	}
	cmd.AddCommand(rcli.NewQuestCmd("writing", writing.Definition{}, rcli.Options{}))
	return cmd
}
