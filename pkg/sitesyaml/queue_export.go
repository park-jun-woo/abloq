//ff:type feature=sitesyaml type=schema
//ff:what 사이트별 큐 export 설정 — 블로그 저장소 origin URL과 커밋 author 2종 (QUEUE_EXPORT_* env의 사이트 행 대응)
package sitesyaml

// QueueExport carries the per-site queue exporter settings. An empty RepoURL
// means the site does not export its queue (the backend keeps the current
// "repo URL unset = 500" behaviour per site).
type QueueExport struct {
	RepoURL     string `yaml:"repo_url" json:"repo_url"`
	Author      string `yaml:"author" json:"author"`
	AuthorEmail string `yaml:"author_email" json:"author_email"`
}
