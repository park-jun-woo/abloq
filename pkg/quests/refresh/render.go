//ff:func feature=quest type=generator control=sequence topic=queue
//ff:what Render — 갱신 프롬프트 조립: 대상 경로·발급 근거(lastmod·임계)·제출 형식 헤더 + 공통 프로토콜 + tasks.md + context.md (read-only)
package refresh

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
	"github.com/park-jun-woo/abloq/quests"
	refreshdocs "github.com/park-jun-woo/abloq/quests/refresh"
)

// Render composes the refresh prompt `next` shows: the seeded paths, the
// queue payload rationale (current lastmod + freshness threshold), the
// submit instructions, the shared consumption protocol, the task tree and
// the refresh conventions. It never mutates the session (read-only).
func (Definition) Render(_ *quest.Session, it *quest.Item) (string, error) {
	var p common.QueuePayload
	if err := it.DecodePayload(&p); err != nil {
		return "", err
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "# refresh quest — %s\n\n", it.Key)
	fmt.Fprintf(&sb, "- instance root: %s\n", p.Root)
	fmt.Fprintf(&sb, "- target article: %s — refresh it in the working tree (do NOT commit before submit)\n", p.Article)
	fmt.Fprintf(&sb, "- stale since: lastmod %s exceeded the %s-day freshness window\n",
		p.Queue["lastmod"], p.Queue["freshness_days"])
	fmt.Fprintf(&sb, "- language companions (keys): %s\n", strings.Join(p.Keys, ", "))
	fmt.Fprintf(&sb, "- submit: abloq quest refresh submit --key %s --in <submission.json>\n", it.Key)
	fmt.Fprintf(&sb, "  submission.json: {\"article\": %q}\n", p.Article)
	sb.WriteString("\n---\n\n")
	sb.WriteString(quests.ProtocolMD)
	sb.WriteString("\n---\n\n")
	sb.WriteString(refreshdocs.TasksMD)
	sb.WriteString("\n---\n\n")
	sb.WriteString(refreshdocs.ContextMD)
	return sb.String(), nil
}
