//ff:func feature=queueio type=parser control=sequence
//ff:what PublishConfigFromEnv가 QUEUE_EXPORT_*를 공유하되 workdir에 -reports 접미사를 붙이는지 검증
package queueio

import "testing"

func TestPublishConfigFromEnv(t *testing.T) {
	t.Setenv("QUEUE_EXPORT_REPO_URL", "file:///tmp/q.git")
	t.Setenv("QUEUE_EXPORT_WORKDIR", "/tmp/q-work")
	cfg, err := PublishConfigFromEnv()
	if err != nil {
		t.Fatalf("PublishConfigFromEnv: %v", err)
	}
	if cfg.Workdir != "/tmp/q-work-reports" {
		t.Errorf("want workdir /tmp/q-work-reports, got %q", cfg.Workdir)
	}
	if cfg.RepoURL != "file:///tmp/q.git" {
		t.Errorf("repo url must be shared: %q", cfg.RepoURL)
	}
	t.Setenv("QUEUE_EXPORT_REPO_URL", "")
	if _, err := PublishConfigFromEnv(); err == nil {
		t.Error("missing REPO_URL must error")
	}
}
