//ff:type feature=quest type=schema
//ff:what Item.Payload — 인스턴스 루트(절대), 루트 기준 원문/번역 글 경로, 원문·대상 언어와 Key 부품(lang/section/slug)
package translation

// Payload is the translation quest item's persisted payload: the blog
// instance root (absolute, where blog.yaml lives), the root-relative origin
// (default-language) and target translation article paths, plus the languages
// and the key parts (lang/section/slug) the item key was derived from.
type Payload struct {
	Root       string `json:"root"`
	Origin     string `json:"origin"`
	Article    string `json:"article"`
	OriginLang string `json:"origin_lang"`
	Lang       string `json:"lang"`
	Section    string `json:"section"`
	Slug       string `json:"slug"`
}
