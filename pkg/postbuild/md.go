//ff:func feature=postbuild type=generator control=iteration dimension=1
//ff:what .md 병행 서빙 본체 — content/ 전 글을 빌드 산출물(public/) 옆에 노이즈 제로 .md로 기록, 생성 수 반환
package postbuild

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// MD writes a clean .md beside every built article: dir/content/... sources
// become dir/public/{lang}/{section}/{slug}.md (AI context format; a
// root-served default language drops its language directory). It is
// idempotent — same inputs rewrite the same bytes.
func MD(dir string, b *blogyaml.Blog) (int, error) {
	contentDir := filepath.Join(dir, "content")
	publicDir := filepath.Join(dir, "public")
	posts, err := CollectPosts(contentDir)
	if err != nil {
		return 0, err
	}
	for i, src := range posts {
		data, err := os.ReadFile(src)
		if err != nil {
			return i, err
		}
		dest := DestPath(contentDir, publicDir, src, b)
		if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
			return i, err
		}
		if err := os.WriteFile(dest, RenderMD(data), 0o644); err != nil {
			return i, err
		}
	}
	return len(posts), nil
}
