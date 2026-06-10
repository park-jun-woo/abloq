//ff:func feature=content type=parser control=sequence
//ff:what IndexRepo가 픽스처 저장소에서 발행 글만(draft·_index·비마크다운·index.md 없는 번들 제외) 인덱싱하고 필드(slug 오버라이드·lastmod 폴백·URL·지표)를 채우는지 검증
package content

import (
	"path/filepath"
	"testing"
)

func TestIndexRepo(t *testing.T) {
	entries, err := IndexRepo(filepath.Join("testdata", "fixture"))
	if err != nil {
		t.Fatalf("IndexRepo: %v", err)
	}
	if len(entries) != 3 {
		t.Fatalf("want 3 entries, got %d: %+v", len(entries), entries)
	}
	a := entries[0]
	if a.Slug != "post-a" || a.Lang != "ko" || a.Section != "tech" {
		t.Errorf("entries[0] identity = %+v", a)
	}
	if a.Title != "포스트 A" || a.Date != "2026-06-01" || a.Lastmod != "2026-06-05" {
		t.Errorf("post-a front matter = %+v", a)
	}
	if a.URL != "https://fixture.example.com/tech/post-a/" {
		t.Errorf("post-a URL = %q (root-served default lang)", a.URL)
	}
	if a.WordCount == 0 || a.InternalLinks != 1 || a.SourceCount != 2 {
		t.Errorf("post-a metrics = %+v", a)
	}
	if len(a.Tags) != 2 || a.Tags[0] != "geo" {
		t.Errorf("post-a tags = %v", a.Tags)
	}
	b := entries[1]
	if b.Slug != "post-b" || b.Title != "post-b" {
		t.Errorf("post-b title must fall back to slug: %+v", b)
	}
	if b.Lastmod != "2026-06-03" {
		t.Errorf("post-b lastmod must fall back to date: %+v", b)
	}
	if b.Tags == nil || len(b.Tags) != 0 {
		t.Errorf("post-b tags must be an empty non-nil slice: %#v", b.Tags)
	}
	c := entries[2]
	if c.Slug != "custom-en" || c.Lang != "en" {
		t.Errorf("front matter slug must override the file stem: %+v", c)
	}
	if c.URL != "https://fixture.example.com/en/tech/custom-en/" {
		t.Errorf("custom-en URL = %q", c.URL)
	}
	if c.InternalLinks != 1 || c.SourceCount != 1 {
		t.Errorf("custom-en metrics = %+v", c)
	}
}
