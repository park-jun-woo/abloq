//ff:type feature=blogyaml type=schema
//ff:what blog.yaml site 섹션 — baseURL/title/author, hugo.toml 생성의 입력
package blogyaml

// Site declares the blog identity. baseURL must be absolute http(s) with no query/fragment.
type Site struct {
	BaseURL string `yaml:"baseURL" json:"baseURL"`
	Title   string `yaml:"title" json:"title"`
	Author  string `yaml:"author" json:"author"`
}
