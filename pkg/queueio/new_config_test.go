//ff:func feature=queueio type=parser control=sequence
//ff:what NewConfig가 env를 읽지 않고 인자만으로 Config를 만들며 repoURL·workdir 필수와 author 기본값을 지키는지 검증
package queueio

import "testing"

func TestNewConfig(t *testing.T) {
	// the env must be irrelevant: only the arguments may decide the Config
	t.Setenv("QUEUE_EXPORT_REPO_URL", "file:///tmp/env-must-lose.git")
	t.Setenv("QUEUE_EXPORT_WORKDIR", "/tmp/env-must-lose")
	t.Setenv("QUEUE_EXPORT_AUTHOR", "env-bot")
	t.Setenv("QUEUE_EXPORT_AUTHOR_EMAIL", "env@example.com")

	if _, err := NewConfig("", "/tmp/q-work", "", ""); err == nil {
		t.Error("missing repo URL must error")
	}
	if _, err := NewConfig("file:///tmp/q.git", "", "", ""); err == nil {
		t.Error("missing workdir must error")
	}

	cfg, err := NewConfig("file:///tmp/q.git", "/tmp/q-work", "", "")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	if cfg.RepoURL != "file:///tmp/q.git" || cfg.Workdir != "/tmp/q-work" {
		t.Errorf("fields not taken from arguments: %+v", cfg)
	}
	if cfg.AuthorName != "abloqd-bot" || cfg.AuthorEmail != "abloqd-bot@abloq.local" {
		t.Errorf("author defaults missing: %+v", cfg)
	}

	cfg, err = NewConfig("file:///tmp/q.git", "/tmp/q-work", "site-bot", "site@example.com")
	if err != nil {
		t.Fatalf("NewConfig with author: %v", err)
	}
	if cfg.AuthorName != "site-bot" || cfg.AuthorEmail != "site@example.com" {
		t.Errorf("author arguments ignored: %+v", cfg)
	}
}
