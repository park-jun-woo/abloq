//ff:func feature=init type=command control=iteration dimension=1
//ff:what 블로그 골격 디렉토리 생성 — 언어별 콘텐츠 디렉토리(initLangDirs)와 quests/queue, .gitkeep 포함
package main

import (
	"os"
	"path/filepath"
)

// initContentDirs creates content/{lang}/{section}/ for every language plus
// the quests/queue/ work-queue directory.
func initContentDirs(dir string, langs, sections []string) error {
	for _, lang := range langs {
		if err := initLangDirs(dir, lang, sections); err != nil {
			return err
		}
	}
	queue := filepath.Join(dir, "quests", "queue")
	if err := os.MkdirAll(queue, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(queue, ".gitkeep"), nil, 0o644)
}
