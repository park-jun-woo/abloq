//ff:type feature=blogyaml type=schema
//ff:what blog.yaml site 섹션 — baseURL/title/author/기본언어 서브디렉토리 여부, hugo.toml 생성의 입력
package blogyaml

// Site declares the blog identity. baseURL must be absolute http(s) with no query/fragment.
// DefaultLangInSubdir (default true) decides whether the default language is
// served under /{lang}/ like every other language, or at the site root.
type Site struct {
	BaseURL             string `yaml:"baseURL" json:"baseURL"`
	Title               string `yaml:"title" json:"title"`
	Author              string `yaml:"author" json:"author"`
	DefaultLangInSubdir bool   `yaml:"default_lang_in_subdir" json:"default_lang_in_subdir"`
}
