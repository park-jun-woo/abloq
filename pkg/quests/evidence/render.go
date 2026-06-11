//ff:func feature=quest type=generator control=sequence topic=queue
//ff:what Render — 보강 프롬프트 조립: 대상 경로·payload 검출 내역(claims·rot_urls 원문)·제출 형식 헤더 + 공통 프로토콜 + tasks.md + context.md (read-only)
package evidence

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/quests/common"
	"github.com/park-jun-woo/abloq/quests"
	evidencedocs "github.com/park-jun-woo/abloq/quests/evidence"
)

// Render composes the evidence prompt `next` shows: the seeded paths, the
// frozen queue payload's findings verbatim (claims JSON + rot URLs JSON),
// the submit instructions, the shared consumption protocol, the task tree
// and the sourcing conventions. It never mutates the session (read-only).
func (Definition) Render(_ *quest.Session, it *quest.Item) (string, error) {
	var p common.QueuePayload
	if err := it.DecodePayload(&p); err != nil {
		return "", err
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "# evidence quest — %s\n\n", it.Key)
	fmt.Fprintf(&sb, "- instance root: %s\n", p.Root)
	fmt.Fprintf(&sb, "- target article: %s — source it in the working tree (do NOT commit before submit)\n", p.Article)
	if claims := p.Queue["claims"]; claims != "" {
		fmt.Fprintf(&sb, "- unsourced claims (hash·loc·text — the change-authorization list):\n  %s\n", claims)
	}
	if rots := p.Queue["rot_urls"]; rots != "" {
		fmt.Fprintf(&sb, "- confirmed rot URLs (replace every one): %s\n", rots)
	}
	fmt.Fprintf(&sb, "- language companions (keys): %s\n", strings.Join(p.Keys, ", "))
	fmt.Fprintf(&sb, "- submit: abloq quest evidence submit --key %s --in <submission.json>\n", it.Key)
	fmt.Fprintf(&sb, "  submission.json: {\"article\": %q}\n", p.Article)
	sb.WriteString("\n---\n\n")
	sb.WriteString(quests.ProtocolMD)
	sb.WriteString("\n---\n\n")
	sb.WriteString(evidencedocs.TasksMD)
	sb.WriteString("\n---\n\n")
	sb.WriteString(evidencedocs.ContextMD)
	return sb.String(), nil
}
