//ff:func feature=quest type=generator control=sequence
//ff:what Render — 번역 프롬프트 조립: 대상 경로·언어쌍·제출 형식 헤더 + 헤딩 맵(원문→대상) + 원문 전문 + tasks.md + context.md (read-only)
package translation

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	translationdocs "github.com/park-jun-woo/abloq/quests/translation"
)

// Render composes the translation prompt `next` shows: the seeded paths and
// submit instructions, the per-language section heading map from blog.yaml,
// the verbatim origin article, the task tree and the translation conventions.
// It never mutates the session (read-only).
func (Definition) Render(_ *quest.Session, it *quest.Item) (string, error) {
	var p Payload
	if err := it.DecodePayload(&p); err != nil {
		return "", err
	}
	b, diags, err := blogyaml.Load(filepath.Join(p.Root, "blog.yaml"))
	if err != nil {
		return "", err
	}
	if len(diags) > 0 {
		return "", fmt.Errorf("blog.yaml: %s", diags[0].String())
	}
	raw, err := os.ReadFile(filepath.Join(p.Root, filepath.FromSlash(p.Origin)))
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "# translation quest — %s\n\n", it.Key)
	fmt.Fprintf(&sb, "- instance root: %s\n", p.Root)
	fmt.Fprintf(&sb, "- origin article (%s): %s\n", p.OriginLang, p.Origin)
	fmt.Fprintf(&sb, "- target article (%s): %s — write the translation here\n", p.Lang, p.Article)
	fmt.Fprintf(&sb, "- front matter: copy the origin's date and lastmod verbatim; keep the slug\n")
	fmt.Fprintf(&sb, "- internal article links: rewrite to the /%s/... prefix\n", p.Lang)
	fmt.Fprintf(&sb, "- submit: abloq quest translation submit --key %s --in <submission.json>\n", it.Key)
	fmt.Fprintf(&sb, "  submission.json: {\"article\": %q}\n", p.Article)
	fmt.Fprintf(&sb, "\n## section headings (%s → %s)\n\n", p.OriginLang, p.Lang)
	sb.WriteString(strings.Join(headingMapLines(b, p.OriginLang, p.Lang), "\n"))
	sb.WriteString("\n\n## origin article (verbatim — never modify the origin file)\n\n")
	sb.WriteString("----- BEGIN ORIGIN -----\n")
	sb.Write(raw)
	sb.WriteString("----- END ORIGIN -----\n\n---\n\n")
	sb.WriteString(translationdocs.TasksMD)
	sb.WriteString("\n---\n\n")
	sb.WriteString(translationdocs.ContextMD)
	return sb.String(), nil
}
