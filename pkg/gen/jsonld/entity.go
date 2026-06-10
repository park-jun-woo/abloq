//ff:type feature=gen type=schema
//ff:what JSON-LD 엔티티 1개 — @type/name/url, 저자(Person) 표현
package jsonld

// Entity is one schema.org entity reference consumed by the JSON-LD partial.
type Entity struct {
	Type string `json:"@type"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
