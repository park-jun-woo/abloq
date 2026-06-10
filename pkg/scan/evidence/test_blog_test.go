//ff:func feature=scan type=parser control=sequence topic=evidence
//ff:what 근거 스캐너 테스트 공통 blog.yaml(ko/tech, 출처 헤딩)을 임시 디렉토리에 쓰고 로드 — 진단이 있으면 즉시 실패
package evidence

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

const testBlogYAML = "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
	"languages: [ko]\nsections: [tech]\n" +
	"structure:\n  order: [body, sources]\n  headings:\n    sources: { ko: \"출처\" }\n"

func testBlog(t *testing.T) *blogyaml.Blog {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(testBlogYAML), 0o644); err != nil {
		t.Fatalf("write blog.yaml: %v", err)
	}
	b, diags, err := blogyaml.Load(filepath.Join(dir, "blog.yaml"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(diags) != 0 {
		t.Fatalf("fixture blog.yaml has diagnostics: %v", diags)
	}
	return b
}
