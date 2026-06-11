//ff:type feature=sitesyaml type=schema
//ff:what 사이트 1개 선언 — name 슬러그(필수)·repo_path 절대경로(필수)·queue_export/gsc/cf_log_source/indexnow_key(선택)·active(기본 true)
package sitesyaml

// Site declares one blog site of the instance. Name is the URL-safe slug
// used as the {site} path parameter; RepoPath is the absolute path of the
// mounted blog repository. The optional groups carry the per-site values
// that used to be instance-global environment variables.
type Site struct {
	Name        string      `yaml:"name" json:"name"`
	RepoPath    string      `yaml:"repo_path" json:"repo_path"`
	QueueExport QueueExport `yaml:"queue_export" json:"queue_export"`
	GSC         GSC         `yaml:"gsc" json:"gsc"`
	CFLogSource string      `yaml:"cf_log_source" json:"cf_log_source"`
	IndexNowKey string      `yaml:"indexnow_key" json:"indexnow_key"`
	Active      bool        `yaml:"active" json:"active"` // absent key defaults to true (injected after decode)
}
