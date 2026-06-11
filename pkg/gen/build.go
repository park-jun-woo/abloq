//ff:func feature=gen type=generator control=sequence
//ff:what 파생물을 고정 순서(hugo.toml/robots.txt/llms.txt/jsonld.json)로 생성 — 같은 입력이면 바이트 동일, llms_txt mode manual/off면 llms.txt 제외
//ff:why mode != auto의 옵트아웃 게이트는 이 한 곳 — check가 Build 목록을 기대 파일로 쓰므로 generate·check가 동시에 llms.txt에서 손을 뗀다 (BUG001)
package gen

import (
	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/gen/hugo"
	"github.com/park-jun-woo/abloq/pkg/gen/jsonld"
	"github.com/park-jun-woo/abloq/pkg/gen/llms"
	"github.com/park-jun-woo/abloq/pkg/gen/robots"
)

// Build renders every derived file for the blog rooted at dir.
// dir is only read for content/ (llms.txt); all other inputs come from b.
// With geo.llms_txt mode manual/off the llms.txt output is omitted entirely:
// generate never touches the hand-curated file and check never enforces it.
func Build(dir string, b *blogyaml.Blog) []Output {
	outs := []Output{
		{Path: "hugo.toml", Data: hugo.Render(b)},
		{Path: "static/robots.txt", Data: robots.Render(b)},
	}
	if mode := b.Geo.LlmsTxtMode(); mode != "manual" && mode != "off" {
		outs = append(outs, Output{Path: "static/llms.txt", Data: llms.Render(b, llms.Collect(dir, b))})
	}
	return append(outs, Output{Path: "data/jsonld.json", Data: jsonld.Render(b)})
}
