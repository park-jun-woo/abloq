//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what 테스트 헬퍼 — git 커밋된 인스턴스(고립+링크부족 thin·후보 hub·무관 extra·kind=cluster 큐 파일) 생성
package cluster

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/queueio"
)

const instanceBlogYAML = `site:
  baseURL: https://example.com
  title: Cluster Fixture
  author: Tester
  default_lang_in_subdir: false

languages: [en]
sections: [posts]

geo:
  min_internal_links: 1
`

const thinArticleMD = `---
title: "Thin"
date: 2026-06-01
lastmod: 2026-06-02
tags: [geo]
---

An isolated article with no internal links yet.
`

const hubArticleMD = `---
title: "Hub"
date: 2026-06-03
lastmod: 2026-06-04
tags: [geo]
---

The hub links [extra](/posts/extra/) for context.
`

const extraArticleMD = `---
title: "Extra"
date: 2026-06-05
lastmod: 2026-06-06
tags: [geo]
---

Extra links back to the [hub](/posts/hub/).
`

func writeInstance(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, root, "blog.yaml", instanceBlogYAML)
	writeFile(t, root, "content/en/posts/thin.md", thinArticleMD)
	writeFile(t, root, "content/en/posts/hub.md", hubArticleMD)
	writeFile(t, root, "content/en/posts/extra.md", extraArticleMD)
	it := queueio.Item{Kind: "cluster", Slug: "thin", Lang: "en", Section: "posts",
		Priority: 3, Keys: []string{"en/posts/thin"},
		Payload: map[string]string{
			"violations": `[{"rule":"min-internal-links","detail":"outbound internal links 0 below min 1"},` +
				`{"rule":"no-isolated-post","detail":"no inbound internal links"}]`,
			"candidates": `[{"section":"posts","slug":"hub","shared_tags":1,"directions":["out","in"]}]`,
		}}
	if err := queueio.WriteDir(filepath.Join(root, "quests", "queue"), []queueio.Item{it}); err != nil {
		t.Fatal(err)
	}
	mustGit(t, root, "init", "-q", "-b", "main")
	mustGit(t, root, "add", "-A")
	mustGit(t, root, "-c", "user.name=t", "-c", "user.email=t@test", "commit", "-q", "-m", "fixture")
	return root
}
