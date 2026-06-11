//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what 테스트 헬퍼 — git 커밋된 인스턴스(blog.yaml·전 룰 통과 글·kind=refresh 큐 파일) 생성
package refresh

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/queueio"
)

const instanceBlogYAML = `site:
  baseURL: https://example.com
  title: Refresh Fixture
  author: Tester

languages: [en]
sections: [posts]

structure:
  order: [body, sources]
  headings:
    sources: { en: "Sources" }
`

const baseArticleMD = `---
title: "Test Article"
date: 2026-06-01
lastmod: 2026-06-02
tags: [test]
---

![main](cover.png)

*Image: by Tester*

This stale body sentence still describes the situation as of early 2025 in vendor terms.

Throughput grew 40% in 2025 per the vendor study. [Vendor study](https://example.org/spec)

## Sources

- [Vendor study](https://example.org/spec)
`

func writeInstance(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, root, "blog.yaml", instanceBlogYAML)
	writeFile(t, root, "content/en/posts/a.md", baseArticleMD)
	it := queueio.Item{Kind: "refresh", Slug: "a", Lang: "en", Section: "posts",
		Priority: 7, Keys: []string{"en/posts/a"},
		Payload: map[string]string{"freshness_days": "90", "lastmod": "2026-06-02"}}
	if err := queueio.WriteDir(filepath.Join(root, "quests", "queue"), []queueio.Item{it}); err != nil {
		t.Fatal(err)
	}
	mustGit(t, root, "init", "-q", "-b", "main")
	mustGit(t, root, "add", "-A")
	mustGit(t, root, "-c", "user.name=t", "-c", "user.email=t@test", "commit", "-q", "-m", "fixture")
	return root
}
