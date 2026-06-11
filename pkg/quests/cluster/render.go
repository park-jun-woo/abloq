//ff:func feature=quest type=generator control=sequence topic=queue
//ff:what Render — 큐레이션 프롬프트 조립: 대상 경로·payload 위반·후보 원문·제출 형식 헤더 + 공통 프로토콜 + tasks.md + context.md (read-only)
package cluster

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
	"github.com/park-jun-woo/abloq/quests"
	clusterdocs "github.com/park-jun-woo/abloq/quests/cluster"
)

// Render composes the cluster prompt `next` shows: the seeded paths, the
// frozen queue payload's violations and link candidates verbatim, the submit
// instructions, the shared consumption protocol, the task tree and the
// curation conventions. It never mutates the session (read-only).
func (Definition) Render(_ *quest.Session, it *quest.Item) (string, error) {
	var p common.QueuePayload
	if err := it.DecodePayload(&p); err != nil {
		return "", err
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "# cluster quest — %s\n\n", it.Key)
	fmt.Fprintf(&sb, "- instance root: %s\n", p.Root)
	fmt.Fprintf(&sb, "- target article: %s — curate it in the working tree (do NOT commit before submit)\n", p.Article)
	fmt.Fprintf(&sb, "- violations (resolve every kind on this article):\n  %s\n", p.Queue["violations"])
	if cands := p.Queue["candidates"]; cands != "" {
		fmt.Fprintf(&sb, "- link candidates (the only other articles you may edit):\n  %s\n", cands)
	}
	fmt.Fprintf(&sb, "- language companions (keys): %s\n", strings.Join(p.Keys, ", "))
	fmt.Fprintf(&sb, "- submit: abloq quest cluster submit --key %s --in <submission.json>\n", it.Key)
	fmt.Fprintf(&sb, "  submission.json: {\"article\": %q}\n", p.Article)
	sb.WriteString("\n---\n\n")
	sb.WriteString(quests.ProtocolMD)
	sb.WriteString("\n---\n\n")
	sb.WriteString(clusterdocs.TasksMD)
	sb.WriteString("\n---\n\n")
	sb.WriteString(clusterdocs.ContextMD)
	return sb.String(), nil
}
