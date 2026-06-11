//ff:func feature=cli type=command control=sequence
//ff:what 임시 인사이트 픽스처(플랫 글 + 규약 위치 사이드카 insight.yaml)를 만들어 insight match 테스트에 제공
package main

import (
	"os"
	"path/filepath"
	"testing"
)

// writeInsightFixture creates content/en/tech/post.md plus its conventional
// sidecar insight.yaml in a temp dir and returns (insightPath, articlePath).
func writeInsightFixture(t *testing.T, insightYAML string) (string, string) {
	t.Helper()
	dir := filepath.Join(t.TempDir(), "content", "en", "tech")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	articlePath := filepath.Join(dir, "post.md")
	article := "---\ntitle: \"P\"\n---\n\nThe ratchet never moves backward.\n"
	if err := os.WriteFile(articlePath, []byte(article), 0o644); err != nil {
		t.Fatal(err)
	}
	insightPath := filepath.Join(dir, "post.insight.yaml")
	if err := os.WriteFile(insightPath, []byte(insightYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	return insightPath, articlePath
}
