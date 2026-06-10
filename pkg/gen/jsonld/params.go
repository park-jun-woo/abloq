//ff:type feature=gen type=schema
//ff:what JSON-LD 파셜(layouts/partials/jsonld.html, Phase005)이 소비하는 파라미터 — 타입 목록/저자 엔티티/발행자
package jsonld

// Params is written to data/jsonld.json; the Hugo partial reads it as site data.
// Struct field order fixes the JSON key order, keeping the output deterministic.
type Params struct {
	Types     []string `json:"types"`
	Author    Entity   `json:"author"`
	Publisher string   `json:"publisher"`
}
