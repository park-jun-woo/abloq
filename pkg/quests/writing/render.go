//ff:func feature=quest type=generator control=sequence
//ff:what Render — 저작 프롬프트 조립: 대상 경로·제출 형식 헤더 + insight.yaml 원문 + tasks.md(T1~T4) + context.md (read-only)
package writing

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/reins/pkg/quest"

	writingdocs "github.com/park-jun-woo/abloq/quests/writing"
)

// Render composes the authoring prompt `next` shows: the seeded target paths
// and submit instructions, the raw insight spec, the embedded T1~T4 task tree
// and the authoring conventions. It never mutates the session (read-only).
func (Definition) Render(_ *quest.Session, it *quest.Item) (string, error) {
	var p Payload
	if err := it.DecodePayload(&p); err != nil {
		return "", err
	}
	raw, err := os.ReadFile(filepath.Join(p.Root, filepath.FromSlash(p.Insight)))
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "# writing quest — %s\n\n", it.Key)
	fmt.Fprintf(&sb, "- instance root: %s\n", p.Root)
	fmt.Fprintf(&sb, "- insight spec: %s\n", p.Insight)
	fmt.Fprintf(&sb, "- target article: %s\n", p.Article)
	fmt.Fprintf(&sb, "- worklog (convention): quests/writing/logs/%s.md\n", p.Slug)
	fmt.Fprintf(&sb, "- review record (convention): quests/writing/reviews/%s.md\n", p.Slug)
	fmt.Fprintf(&sb, "- submit: abloq quest writing submit --key %s --in <submission.json>\n", it.Key)
	sb.WriteString("\n## insight.yaml\n\n```yaml\n")
	sb.Write(raw)
	sb.WriteString("```\n\n---\n\n")
	sb.WriteString(writingdocs.TasksMD)
	sb.WriteString("\n---\n\n")
	sb.WriteString(writingdocs.ContextMD)
	return sb.String(), nil
}
