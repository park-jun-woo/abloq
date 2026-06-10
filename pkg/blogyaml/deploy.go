//ff:type feature=blogyaml type=schema
//ff:what blog.yaml deploy 섹션 — provider/terraform/indexnow 배포 선언
package blogyaml

// Deploy declares how the built site is shipped.
type Deploy struct {
	Provider  string `yaml:"provider" json:"provider"`
	Terraform bool   `yaml:"terraform" json:"terraform"`
	IndexNow  bool   `yaml:"indexnow" json:"indexnow"`
}
