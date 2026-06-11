//ff:func feature=cli type=command control=sequence
//ff:what "abloq quest" 부모 명령 생성 — reins NewQuestCmd("writing"/"translation"/"refresh"/"evidence"/"cluster")를 마운트 (scan/next/submit/status/export/rules)
package main

import (
	rcli "github.com/park-jun-woo/reins/pkg/cli"
	"github.com/spf13/cobra"

	qcluster "github.com/park-jun-woo/abloq/pkg/quests/cluster"
	qevidence "github.com/park-jun-woo/abloq/pkg/quests/evidence"
	qrefresh "github.com/park-jun-woo/abloq/pkg/quests/refresh"
	"github.com/park-jun-woo/abloq/pkg/quests/translation"
	"github.com/park-jun-woo/abloq/pkg/quests/writing"
)

// newQuestCmd builds the quest command group. Each quest is a reins consumer:
// reins supplies the ratchet/commands, the Definition supplies the domain.
// refresh/evidence/cluster are the Phase018 queue consumers — their scan
// reads quests/queue/ instead of taking spec/article arguments.
func newQuestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quest",
		Short: "Agent quests (reins-gated): writing, translation, refresh, evidence, cluster",
	}
	cmd.AddCommand(rcli.NewQuestCmd("writing", writing.Definition{}, rcli.Options{}))
	cmd.AddCommand(rcli.NewQuestCmd("translation", translation.Definition{}, rcli.Options{}))
	cmd.AddCommand(rcli.NewQuestCmd("refresh", qrefresh.Definition{}, rcli.Options{}))
	cmd.AddCommand(rcli.NewQuestCmd("evidence", qevidence.Definition{}, rcli.Options{}))
	cmd.AddCommand(rcli.NewQuestCmd("cluster", qcluster.Definition{}, rcli.Options{}))
	return cmd
}
