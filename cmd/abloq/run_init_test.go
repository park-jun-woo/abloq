//ff:func feature=init type=command control=iteration dimension=1
//ff:what runInit이 blog.yaml/템플릿/골격/파생물/CLAUDE.md를 만들어 validate+gate 통과 상태로 끝내고 재실행을 거부하는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunInit(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "blog")
	o := initOpts{Title: "T", BaseURL: "https://t.example.com", Author: "A",
		Languages: []string{"ko", "en"}, Sections: []string{"tech"}}
	var out bytes.Buffer
	if err := runInit(&out, strings.NewReader(""), dir, o); err != nil {
		t.Fatalf("runInit: %v\noutput: %s", err, out.String())
	}
	musts := []string{
		"blog.yaml", "CLAUDE.md", "README.md", ".gitignore",
		"hugo.toml", "static/robots.txt", "static/llms.txt", "data/jsonld.json",
		"layouts/_default/baseof.html", "layouts/partials/jsonld.html",
		"assets/css/main.css", "deploy/terraform/main.tf",
		"content/ko/tech/.gitkeep", "content/en/tech/.gitkeep", "quests/queue/.gitkeep",
	}
	for _, m := range musts {
		if _, err := os.Stat(filepath.Join(dir, m)); err != nil {
			t.Errorf("init must create %s: %v", m, err)
		}
	}
	if err := runInit(&out, strings.NewReader(""), dir, o); err == nil {
		t.Error("second runInit must refuse to overwrite blog.yaml")
	}
	o2 := o
	o2.Interactive = true // EOF on stdin -> all defaults kept (agent-safe fallback)
	dir2 := filepath.Join(t.TempDir(), "blog2")
	if err := runInit(&out, strings.NewReader(""), dir2, o2); err != nil {
		t.Errorf("interactive init with EOF input must fall back to defaults: %v", err)
	}
	blockedParent := filepath.Join(t.TempDir(), "file")
	if err := os.WriteFile(blockedParent, []byte(""), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := runInit(&out, strings.NewReader(""), filepath.Join(blockedParent, "blog"), o); err == nil {
		t.Error("init under a regular file must error")
	}
}
