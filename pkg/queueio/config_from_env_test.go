//ff:func feature=queueio type=parser control=sequence
//ff:what ConfigFromEnv가 REPO_URL·WORKDIR 필수 검증과 author 기본값(abloqd-bot)을 적용하는지 검증
package queueio

import "testing"

func TestConfigFromEnv(t *testing.T) {
	t.Setenv("QUEUE_EXPORT_REPO_URL", "")
	t.Setenv("QUEUE_EXPORT_WORKDIR", "")
	if _, err := ConfigFromEnv(); err == nil {
		t.Error("missing REPO_URL must error")
	}
	t.Setenv("QUEUE_EXPORT_REPO_URL", "file:///tmp/q.git")
	if _, err := ConfigFromEnv(); err == nil {
		t.Error("missing WORKDIR must error")
	}
	t.Setenv("QUEUE_EXPORT_WORKDIR", "/tmp/q-work")
	cfg, err := ConfigFromEnv()
	if err != nil {
		t.Fatalf("ConfigFromEnv: %v", err)
	}
	if cfg.AuthorName != "abloqd-bot" || cfg.AuthorEmail != "abloqd-bot@abloq.local" {
		t.Errorf("author defaults missing: %+v", cfg)
	}
	t.Setenv("QUEUE_EXPORT_AUTHOR", "custom-bot")
	t.Setenv("QUEUE_EXPORT_AUTHOR_EMAIL", "bot@example.com")
	cfg, _ = ConfigFromEnv()
	if cfg.AuthorName != "custom-bot" || cfg.AuthorEmail != "bot@example.com" {
		t.Errorf("author overrides ignored: %+v", cfg)
	}
}
