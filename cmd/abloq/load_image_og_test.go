//ff:func feature=cli type=command control=sequence
//ff:what loadImageOG 검증 — blog.yaml 부재는 영값(local) 무에러, 유효 선언은 image 블록 반환, 무효 blog.yaml은 에러
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadImageOG(t *testing.T) {
	var out bytes.Buffer

	empty := t.TempDir()
	cfg, err := loadImageOG(&out, empty)
	if err != nil || cfg.OGProvider() != "local" {
		t.Errorf("missing blog.yaml: cfg %+v err %v, want zero/local", cfg, err)
	}

	dir := writeOGBlogFixture(t)
	cfg, err = loadImageOG(&out, dir)
	if err != nil || cfg.OG.Provider != "gemini" || len(cfg.OG.Variants) != 3 {
		t.Errorf("fixture: cfg %+v err %v, want gemini with 3 variants", cfg, err)
	}

	bad := t.TempDir()
	if err := os.WriteFile(filepath.Join(bad, "blog.yaml"), []byte("nope: 1\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if _, err := loadImageOG(&out, bad); err == nil {
		t.Error("invalid blog.yaml must error")
	}
}
