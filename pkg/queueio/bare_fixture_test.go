//ff:func feature=queueio type=client control=sequence
//ff:what 테스트 픽스처 — file:// bare 저장소(초기 커밋 포함)와 exporter Config를 임시 디렉토리에 구성
package queueio

import (
	"os"
	"path/filepath"
	"testing"
)

// bareFixture builds a local bare origin with one initial commit and returns
// an exporter Config pointing a fresh workdir at it. file:// keeps every
// git test offline.
func bareFixture(t *testing.T) Config {
	t.Helper()
	root := t.TempDir()
	bare := filepath.Join(root, "origin.git")
	seed := filepath.Join(root, "seed")
	mustGit(t, "", "init", "--bare", "-b", "main", bare)
	mustGit(t, "", "init", "-b", "main", seed)
	if err := os.WriteFile(filepath.Join(seed, "README.md"), []byte("seed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	mustGit(t, seed, "add", ".")
	mustGit(t, seed, "-c", "user.name=seed", "-c", "user.email=seed@test", "commit", "-m", "seed")
	mustGit(t, seed, "push", "file://"+bare, "main")
	return Config{
		RepoURL:     "file://" + bare,
		Workdir:     filepath.Join(root, "work"),
		AuthorName:  "abloqd-bot",
		AuthorEmail: "abloqd-bot@abloq.local",
	}
}
