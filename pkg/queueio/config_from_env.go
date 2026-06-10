//ff:func feature=queueio type=parser control=sequence
//ff:what QUEUE_EXPORT_* env → Config — REPO_URL·WORKDIR 필수, author 기본 abloqd-bot
package queueio

import (
	"errors"
	"os"
)

// ConfigFromEnv reads the exporter configuration from the environment.
// QUEUE_EXPORT_REPO_URL and QUEUE_EXPORT_WORKDIR are required; the commit
// author defaults to abloqd-bot. Deploy-key SSH wiring stays outside —
// git honours GIT_SSH_COMMAND from the process environment.
func ConfigFromEnv() (Config, error) {
	cfg := Config{
		RepoURL:     os.Getenv("QUEUE_EXPORT_REPO_URL"),
		Workdir:     os.Getenv("QUEUE_EXPORT_WORKDIR"),
		AuthorName:  os.Getenv("QUEUE_EXPORT_AUTHOR"),
		AuthorEmail: os.Getenv("QUEUE_EXPORT_AUTHOR_EMAIL"),
	}
	if cfg.RepoURL == "" {
		return Config{}, errors.New("QUEUE_EXPORT_REPO_URL is not set")
	}
	if cfg.Workdir == "" {
		return Config{}, errors.New("QUEUE_EXPORT_WORKDIR is not set")
	}
	if cfg.AuthorName == "" {
		cfg.AuthorName = "abloqd-bot"
	}
	if cfg.AuthorEmail == "" {
		cfg.AuthorEmail = "abloqd-bot@abloq.local"
	}
	return cfg, nil
}
