//ff:func feature=queueio type=parser control=sequence
//ff:what 필드 직접 주입 발행용 Config 생성 — NewConfig 공유, workdir만 "-reports" 접미사로 분리 (exporter 클론과 동시 사용 충돌 방지)
package queueio

// NewPublishConfig builds the publisher Config from caller-given fields the
// same way NewConfig does, but separates the work clone (Workdir +
// "-reports"): the publisher and the exporter must never race inside one
// checkout.
func NewPublishConfig(repoURL, workdir, authorName, authorEmail string) (Config, error) {
	cfg, err := NewConfig(repoURL, workdir, authorName, authorEmail)
	if err != nil {
		return Config{}, err
	}
	cfg.Workdir += "-reports"
	return cfg, nil
}
