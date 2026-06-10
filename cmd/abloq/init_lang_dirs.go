//ff:func feature=init type=command control=iteration dimension=1
//ff:what 언어 1개의 섹션 디렉토리(content/{lang}/{section}/)를 생성 — 빈 디렉토리 보존용 .gitkeep 포함
package main

import (
	"os"
	"path/filepath"
)

// initLangDirs creates the section directories for one language.
func initLangDirs(dir, lang string, sections []string) error {
	for _, section := range sections {
		p := filepath.Join(dir, "content", lang, section)
		if err := os.MkdirAll(p, 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(p, ".gitkeep"), nil, 0o644); err != nil {
			return err
		}
	}
	return nil
}
