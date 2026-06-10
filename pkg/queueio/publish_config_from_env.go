//ff:func feature=queueio type=parser control=sequence
//ff:what 발행용 Config — QUEUE_EXPORT_* 공유, workdir만 "-reports" 접미사로 분리 (exporter 클론과 동시 사용 충돌 방지)
package queueio

// PublishConfigFromEnv reads the same QUEUE_EXPORT_* environment as the
// queue exporter but separates the work clone (Workdir + "-reports") by
// default: the publisher and the exporter must never race inside one
// checkout. Repo URL, author identity and the deploy-key SSH wiring are
// shared — the report is a publication copy into the same blog repository.
func PublishConfigFromEnv() (Config, error) {
	cfg, err := ConfigFromEnv()
	if err != nil {
		return Config{}, err
	}
	cfg.Workdir += "-reports"
	return cfg, nil
}
