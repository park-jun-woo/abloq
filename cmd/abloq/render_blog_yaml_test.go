//ff:func feature=init type=generator control=sequence
//ff:what renderBlogYAML 산출물이 blogyaml.Load를 진단 0으로 통과하고 입력 값을 그대로 담는지 검증
package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestRenderBlogYAML(t *testing.T) {
	o := initOpts{Title: "T: subtitle", BaseURL: "https://t.example.com", Author: "A",
		Languages: []string{"ko", "en"}, Sections: []string{"opinion", "tech"}}
	dir := t.TempDir()
	path := filepath.Join(dir, "blog.yaml")
	if err := os.WriteFile(path, renderBlogYAML(o), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	b, diags, err := blogyaml.Load(path)
	if err != nil || len(diags) > 0 {
		t.Fatalf("generated blog.yaml must be valid: err %v, diags %v", err, diags)
	}
	if b.Site.Title != "T: subtitle" || b.Languages[0] != "ko" || len(b.Sections) != 2 {
		t.Errorf("round-trip mismatch: %+v", b)
	}
	if b.Structure.Headings["sources"]["ko"] != "출처" {
		t.Errorf("sources heading ko = %q, want 출처", b.Structure.Headings["sources"]["ko"])
	}
}
