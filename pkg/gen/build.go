//ff:func feature=gen type=generator control=sequence
//ff:what 파생물 4종(hugo.toml/robots.txt/llms.txt/jsonld.json)을 고정 순서로 생성 — 같은 입력이면 바이트 동일
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
func Build(dir string, b *blogyaml.Blog) []Output {
	return []Output{
		{Path: "hugo.toml", Data: hugo.Render(b)},
		{Path: "static/robots.txt", Data: robots.Render(b)},
		{Path: "static/llms.txt", Data: llms.Render(b, llms.Collect(dir, b))},
		{Path: "data/jsonld.json", Data: jsonld.Render(b)},
	}
}
