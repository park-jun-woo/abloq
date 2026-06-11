//ff:type feature=sitesyaml type=schema topic=gsc
//ff:what 사이트별 GSC 설정 — 속성 URL(http(s) 또는 sc-domain:)과 SA JSON 파일 경로 (GSC_SITE_URL·GSC_SA_JSON_PATH의 사이트 행 대응)
package sitesyaml

// GSC carries the per-site Search Console settings. SiteURL is the property
// identifier (a URL-prefix property or an "sc-domain:" property); SAJSONPath
// points at the mounted service-account JSON — the path is not a secret, the
// file content is.
type GSC struct {
	SiteURL    string `yaml:"site_url" json:"site_url"`
	SAJSONPath string `yaml:"sa_json_path" json:"sa_json_path"`
}
