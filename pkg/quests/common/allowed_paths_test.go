//ff:func feature=quest type=generator control=sequence topic=queue
//ff:what AllowedPaths가 키마다 글 2형+사이드카 2형을 펼치고 큐 파일은 포함하지 않는지 검증
package common

import "testing"

func TestAllowedPaths(t *testing.T) {
	got := AllowedPaths([]string{"en/posts/a", "ko/posts/a"})
	for _, want := range []string{
		"content/en/posts/a.md", "content/en/posts/a/index.md",
		"content/en/posts/a.insight.yaml", "content/en/posts/a/insight.yaml",
		"content/ko/posts/a.md", "content/ko/posts/a/index.md",
	} {
		if !got[want] {
			t.Errorf("missing allowed path %s", want)
		}
	}
	if len(got) != 8 {
		t.Errorf("want 8 paths (4 per key), got %d: %v", len(got), got)
	}
	if got["quests/queue/refresh-en-posts-a.yaml"] {
		t.Error("the queue file must never be allowed (deletion is the post-gate signal)")
	}
}
