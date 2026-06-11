//ff:type feature=quest type=schema
//ff:what Item.Payload — 인스턴스 루트(절대)와 루트 기준 insight/대상 글 경로, Key 부품(lang/section/slug)
package writing

// Payload is the writing quest item's persisted payload: the blog instance
// root (absolute, where blog.yaml lives) plus the root-relative insight spec
// and target article paths the key (lang/section/slug) was derived from.
type Payload struct {
	Root    string `json:"root"`
	Insight string `json:"insight"`
	Article string `json:"article"`
	Lang    string `json:"lang"`
	Section string `json:"section"`
	Slug    string `json:"slug"`
}
