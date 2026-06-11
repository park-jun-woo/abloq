//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what 테스트 헬퍼 — git 커밋된 인스턴스(무출처 주장+rot 인용 글·kind=evidence 큐 파일·인용 영수증 캐시) 생성
package evidence

import (
	"path/filepath"
	"testing"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
	"github.com/park-jun-woo/abloq/pkg/queueio"
)

const instanceBlogYAML = `site:
  baseURL: https://example.com
  title: Evidence Fixture
  author: Tester

languages: [en]
sections: [posts]

structure:
  order: [body, sources]
  headings:
    sources: { en: "Sources" }
`

// unsourcedClaim is the queued claim line — no source link in its paragraph.
const unsourcedClaim = "Throughput grew 40% after the migration."

// rotURL is the queued confirmed-rot citation.
const rotURL = "https://gone.example/dead-study"

const baseArticleMD = `---
title: "Test Article"
date: 2026-06-01
lastmod: 2026-06-02
tags: [test]
---

![main](cover.png)

*Image: by Tester*

` + unsourcedClaim + `

Latency dropped 120ms in the same run. [Run log](https://example.org/runlog)

Background per the [dead study](` + rotURL + `).

## Sources

- [Run log](https://example.org/runlog)
`

func writeInstance(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, root, "blog.yaml", instanceBlogYAML)
	writeFile(t, root, "content/en/posts/a.md", baseArticleMD)
	// Receipt cache: every fixture URL pre-verified ok, so citation-exists
	// never touches the network in tests.
	writeFile(t, root, ".abloq/citation-receipts.json",
		`{"https://example.org/spec":{"checked_at":"2200-01-01T00:00:00Z","verdict":"ok"},`+
			`"https://example.org/runlog":{"checked_at":"2200-01-01T00:00:00Z","verdict":"ok"},`+
			`"https://example.org/live-study":{"checked_at":"2200-01-01T00:00:00Z","verdict":"ok"}}`)
	it := queueio.Item{Kind: "evidence", Slug: "a", Lang: "en", Section: "posts",
		Priority: 2, Keys: []string{"en/posts/a"},
		Payload: map[string]string{
			"claims":   `[{"hash":"` + agate.HashText(unsourcedClaim) + `","loc":"content/en/posts/a.md:11","text":"` + unsourcedClaim + `"}]`,
			"rot_urls": `["` + rotURL + `"]`,
		}}
	if err := queueio.WriteDir(filepath.Join(root, "quests", "queue"), []queueio.Item{it}); err != nil {
		t.Fatal(err)
	}
	mustGit(t, root, "init", "-q", "-b", "main")
	mustGit(t, root, "add", "-A")
	mustGit(t, root, "-c", "user.name=t", "-c", "user.email=t@test", "commit", "-q", "-m", "fixture")
	return root
}
