//ff:func feature=scan type=parser control=iteration dimension=1 topic=evidence
//ff:what 임시 저장소 픽스처 — 무출처 주장 2건 글 + citeURL 인용 글 + draft 글, Scan 통합 테스트 공용
package evidence

import (
	"os"
	"path/filepath"
	"testing"
)

// writeRepoFixture builds a blog repository with post-claims (two unsourced
// numeric claims, no citations), post-rot (one citation to citeURL, no
// claims) and post-draft (a claim that must never be scanned).
func writeRepoFixture(t *testing.T, citeURL string) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(testBlogYAML), 0o644); err != nil {
		t.Fatalf("write blog.yaml: %v", err)
	}
	postDir := filepath.Join(dir, "content", "ko", "tech")
	if err := os.MkdirAll(postDir, 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	posts := map[string]string{
		"post-claims.md": "---\ntitle: Claims\ndate: 2026-06-01\n---\n\n" +
			"처리량이 이전 분기 대비 40% 증가했다.\n평균 응답 시간은 120ms 단축됐다.\n",
		"post-rot.md": "---\ntitle: Rot\ndate: 2026-06-02\n---\n\n" +
			"본문이다. [참고](" + citeURL + ") 링크.\n",
		"post-draft.md": "---\ntitle: Draft\ndate: 2026-06-03\ndraft: true\n---\n\n" +
			"초안 지표가 70% 상승했다.\n",
	}
	for name, body := range posts {
		if err := os.WriteFile(filepath.Join(postDir, name), []byte(body), 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	return dir
}
