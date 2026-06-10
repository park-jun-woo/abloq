//ff:func feature=gen type=generator control=sequence
//ff:what 골든 디렉토리의 blog.yaml을 Load해 검증 통과를 확인하고 Blog를 반환하는 테스트 헬퍼
package gen

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func loadGoldenBlog(t *testing.T) *blogyaml.Blog {
	t.Helper()
	b, diags, err := blogyaml.Load(filepath.Join("testdata", "golden", "blog.yaml"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(diags) > 0 {
		t.Fatalf("golden blog.yaml must validate: %v", diags)
	}
	return b
}
