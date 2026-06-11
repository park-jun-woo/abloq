//ff:type feature=sitesyaml type=schema
//ff:what sites.yaml 스키마 v1 최상위 — abloqd 인스턴스가 운용하는 사이트 목록 선언 (sites 리스트)
package sitesyaml

// Sites is the root of sites.yaml schema v1. Unknown keys are rejected
// (strict parsing). The list is the instance-level SSOT of every blog site
// one abloqd serves — the backend upserts it into the sites table at boot.
type Sites struct {
	Sites []Site `yaml:"sites" json:"sites"`
}
