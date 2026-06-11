//ff:func feature=queueio type=parser control=sequence
//ff:what 필드 직접 주입 Config 생성 — env 비의존, repoURL·workdir 필수 검증, author 기본 abloqd-bot (사이트 행 값 주입 경로)
//ff:why 멀티사이트: 백엔드가 sites 행 값으로 사이트별 export 설정을 만들 수 있어야 한다 — env 직독(ConfigFromEnv)은 단일 사이트 하위호환으로 남는다 (Phase020)
package queueio

import "errors"

// NewConfig builds an exporter Config from caller-given fields (site row
// values in the multi-site backend) without reading the environment.
// RepoURL and Workdir are required; the commit author defaults to abloqd-bot.
func NewConfig(repoURL, workdir, authorName, authorEmail string) (Config, error) {
	if repoURL == "" {
		return Config{}, errors.New("queue export repo URL is not set")
	}
	if workdir == "" {
		return Config{}, errors.New("queue export workdir is not set")
	}
	cfg := Config{
		RepoURL:     repoURL,
		Workdir:     workdir,
		AuthorName:  authorName,
		AuthorEmail: authorEmail,
	}
	if cfg.AuthorName == "" {
		cfg.AuthorName = "abloqd-bot"
	}
	if cfg.AuthorEmail == "" {
		cfg.AuthorEmail = "abloqd-bot@abloq.local"
	}
	return cfg, nil
}
