//ff:func feature=init type=command control=sequence
//ff:what promptOpts가 입력 답변으로 5개 필드를 덮어쓰고 빈 답변·EOF에서 기본값을 유지하는지 검증
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestPromptOpts(t *testing.T) {
	def := initOpts{Title: "D", BaseURL: "https://d.example.com", Author: "DA",
		Languages: []string{"en"}, Sections: []string{"posts"}}
	var out bytes.Buffer
	in := strings.NewReader("My Blog\n\nWriter\nko,en\n")
	got := promptOpts(&out, in, def) // 5th answer hits EOF -> default
	if got.Title != "My Blog" || got.Author != "Writer" {
		t.Errorf("answers not applied: %+v", got)
	}
	if got.BaseURL != "https://d.example.com" {
		t.Errorf("empty answer must keep default baseURL, got %q", got.BaseURL)
	}
	if len(got.Languages) != 2 || got.Languages[0] != "ko" {
		t.Errorf("languages = %v, want [ko en]", got.Languages)
	}
	if len(got.Sections) != 1 || got.Sections[0] != "posts" {
		t.Errorf("EOF must keep default sections, got %v", got.Sections)
	}
}
