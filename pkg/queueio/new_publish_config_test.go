//ff:func feature=queueio type=parser control=sequence
//ff:what NewPublishConfig가 NewConfig 검증을 공유하되 workdir에 -reports 접미사를 붙이는지 검증
package queueio

import "testing"

func TestNewPublishConfig(t *testing.T) {
	cfg, err := NewPublishConfig("file:///tmp/q.git", "/tmp/q-work", "", "")
	if err != nil {
		t.Fatalf("NewPublishConfig: %v", err)
	}
	if cfg.Workdir != "/tmp/q-work-reports" {
		t.Errorf("want workdir /tmp/q-work-reports, got %q", cfg.Workdir)
	}
	if cfg.RepoURL != "file:///tmp/q.git" {
		t.Errorf("repo url must be shared: %q", cfg.RepoURL)
	}
	if cfg.AuthorName != "abloqd-bot" {
		t.Errorf("author default missing: %+v", cfg)
	}
	if _, err := NewPublishConfig("", "/tmp/q-work", "", ""); err == nil {
		t.Error("missing repo URL must error")
	}
}
