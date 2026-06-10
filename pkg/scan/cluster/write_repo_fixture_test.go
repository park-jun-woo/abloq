//ff:func feature=scan type=parser control=iteration dimension=1 topic=cluster
//ff:what 임시 저장소 픽스처 — 정상 hub·링크부족 thin·고아태그 orphan·고립 island·taxonomy 위반 offtax + draft + en 번역, Scan 통합 테스트 공용
//ff:why backend/fixtures/cluster-blog와 같은 그래프다 — 단위 테스트의 기대값이 그대로 scenario-cluster.hurl의 정확 assert가 된다
package cluster

import (
	"os"
	"path/filepath"
	"testing"
)

// writeRepoFixture builds the cluster fixture repository: hub satisfies
// every rule; thin lacks outlinks; orphan carries the lonely tag alone;
// island has no inlinks and shares the off-taxonomy rogue tag with offtax
// (so rogue is taxonomy drift, not an orphan). A draft and an en
// translation prove the graph only sees published default-language posts.
func writeRepoFixture(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	posts := map[string]string{
		"content/ko/tech/hub.md": "---\ntitle: Hub\ndate: 2026-01-05\ntags: [geo, abloq]\n---\n\n" +
			"허브 글이다. [thin](/tech/thin/)과 [offtax](/tech/offtax/)를 잇고 [외부](https://example.org/x)도 인용한다.\n",
		"content/ko/tech/thin.md": "---\ntitle: Thin\ndate: 2026-01-04\ntags: [geo, abloq]\n---\n\n" +
			"링크가 모자란 글이다. [hub](/tech/hub/)만 잇는다.\n",
		"content/ko/tech/orphan.md": "---\ntitle: Orphan\ndate: 2026-01-03\ntags: [geo, lonely]\n---\n\n" +
			"고아 태그 글이다. [hub](/tech/hub/)와 [thin](/tech/thin/)을 잇는다.\n",
		"content/ko/tech/island.md": "---\ntitle: Island\ndate: 2026-01-02\ntags: [abloq, rogue]\n---\n\n" +
			"아무도 가리키지 않는 글이다. [hub](/tech/hub/)와 [offtax](/tech/offtax/)를 잇는다.\n",
		"content/ko/tech/offtax.md": "---\ntitle: Offtax\ndate: 2026-01-01\ntags: [geo, rogue]\n---\n\n" +
			"taxonomy 밖 태그 글이다. [hub](/tech/hub/)와 [orphan](/tech/orphan/)을 잇는다.\n",
		"content/ko/tech/draft.md": "---\ntitle: Draft\ndate: 2026-01-06\ndraft: true\ntags: [rogue]\n---\n\n" +
			"초안은 그래프 밖이다.\n",
		"content/en/tech/hub.md": "---\ntitle: Hub EN\ndate: 2026-01-05\ntags: [geo]\n---\n\n" +
			"Translation, out of the graph. [thin](/en/tech/thin/).\n",
	}
	for rel, body := range posts {
		path := filepath.Join(dir, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("MkdirAll: %v", err)
		}
		if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
			t.Fatalf("write %s: %v", rel, err)
		}
	}
	return dir
}
