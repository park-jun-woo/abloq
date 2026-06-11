//ff:type feature=insight type=schema
//ff:what 주장 1건 — id/text(한 문장)/kind(claim|rebuttal|prediction|definition)/requires_source/anchors(동의어 허용 목록)
package insight

// Claim is one assertion of the insight spec. Anchors are the article-language
// vocabulary (synonyms allowed) used for deterministic body screening;
// Text stays in the author's language and may differ from the article's.
type Claim struct {
	ID             string   `yaml:"id" json:"id"`
	Text           string   `yaml:"text" json:"text"`
	Kind           string   `yaml:"kind" json:"kind"`
	RequiresSource bool     `yaml:"requires_source" json:"requires_source"`
	Anchors        []string `yaml:"anchors" json:"anchors"`
}
