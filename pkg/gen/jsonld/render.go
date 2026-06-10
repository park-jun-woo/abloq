//ff:func feature=gen type=generator control=sequence
//ff:what blog.yaml에서 data/jsonld.json 바이트를 렌더 — geo.jsonld 타입 목록 + 저자 Person 엔티티, 구조체 필드 순서로 멱등
package jsonld

import (
	"encoding/json"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// Render produces data/jsonld.json deterministically from blog.yaml.
// Params holds only strings, so marshalling cannot fail.
func Render(b *blogyaml.Blog) []byte {
	params := Params{
		Types: b.Geo.JSONLD,
		Author: Entity{
			Type: "Person",
			Name: b.Site.Author,
			URL:  b.Site.BaseURL,
		},
		Publisher: b.Site.Title,
	}
	data, _ := json.MarshalIndent(params, "", "  ")
	return append(data, '\n')
}
