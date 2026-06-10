//ff:func feature=blogyaml type=parser control=sequence
//ff:what 골든 테스트 — parkjunwoo.com 12언어 예제가 진단 0건으로 통과하는지 검증
package blogyaml

import (
	"path/filepath"
	"testing"
)

func TestLoadValid(t *testing.T) {
	b, diags, err := Load(filepath.Join("testdata", "valid", "blog.yaml"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(diags) != 0 {
		t.Fatalf("want 0 diagnostics, got %d: %v", len(diags), diags)
	}
	if len(b.Languages) != 12 {
		t.Errorf("want 12 languages, got %d", len(b.Languages))
	}
	if b.Languages[0] != "ko" {
		t.Errorf("want default language ko, got %q", b.Languages[0])
	}
	if b.Site.BaseURL != "https://parkjunwoo.com" {
		t.Errorf("unexpected baseURL %q", b.Site.BaseURL)
	}
	if len(b.Structure.Order) != 7 {
		t.Errorf("want 7 structure.order entries, got %d", len(b.Structure.Order))
	}
	if b.Geo.Crawlers["bytespider"] != "block" {
		t.Errorf("want bytespider block, got %q", b.Geo.Crawlers["bytespider"])
	}
}
