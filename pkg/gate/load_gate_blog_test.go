//ff:func feature=gate type=parser control=sequence
//ff:what 게이트 테스트 공통 blog.yaml(ko/en, 구조 7키)을 로드 — 진단이 있으면 즉시 실패
package gate

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func loadGateBlog(t *testing.T) *blogyaml.Blog {
	t.Helper()
	b, diags, err := blogyaml.Load(filepath.Join("testdata", "blog.yaml"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(diags) != 0 {
		t.Fatalf("fixture blog.yaml has diagnostics: %v", diags)
	}
	return b
}
