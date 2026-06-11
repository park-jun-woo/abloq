//ff:type feature=blogyaml type=schema
//ff:what geo.llms_txt.pinned 엔트리 1개 — 목록 선두 고정 링크(title/url 필수, desc/group 선택)
package blogyaml

// LlmsPinned is one hand-curated llms.txt entry rendered before the
// collected posts (master index links etc.). URL is an absolute URL or a
// "/"-rooted path. Group optionally names the heading the entry leads;
// when empty the entry sits at the very top of the list without a heading.
type LlmsPinned struct {
	Title string `yaml:"title" json:"title"`
	URL   string `yaml:"url" json:"url"`
	Desc  string `yaml:"desc" json:"desc"`
	Group string `yaml:"group" json:"group"`
}
