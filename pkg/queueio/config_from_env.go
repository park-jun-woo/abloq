//ff:func feature=queueio type=parser control=sequence
//ff:what QUEUE_EXPORT_* env → Config — REPO_URL·WORKDIR 필수(env 이름 그대로 에러), 기본값 주입은 NewConfig 위임
package queueio

import (
	"errors"
	"os"
)

// ConfigFromEnv reads the exporter configuration from the environment.
// QUEUE_EXPORT_REPO_URL and QUEUE_EXPORT_WORKDIR are required (the error
// names the missing variable); defaults come from NewConfig. Deploy-key SSH
// wiring stays outside — git honours GIT_SSH_COMMAND from the process
// environment.
func ConfigFromEnv() (Config, error) {
	repoURL := os.Getenv("QUEUE_EXPORT_REPO_URL")
	if repoURL == "" {
		return Config{}, errors.New("QUEUE_EXPORT_REPO_URL is not set")
	}
	workdir := os.Getenv("QUEUE_EXPORT_WORKDIR")
	if workdir == "" {
		return Config{}, errors.New("QUEUE_EXPORT_WORKDIR is not set")
	}
	return NewConfig(repoURL, workdir, os.Getenv("QUEUE_EXPORT_AUTHOR"), os.Getenv("QUEUE_EXPORT_AUTHOR_EMAIL"))
}
