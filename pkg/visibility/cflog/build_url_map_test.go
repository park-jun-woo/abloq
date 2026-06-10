//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what BuildURLMap이 글마다 페이지 경로(슬래시·index.html)와 .md 경로를 URLLang 규칙으로 등록하는지 검증 — 루트 서빙 기본 언어는 세그먼트 생략
package cflog

import "testing"

func TestBuildURLMap(t *testing.T) {
	root, b := writeRepoFixture(t)
	urls, err := BuildURLMap(root, b)
	if err != nil {
		t.Fatalf("BuildURLMap: %v", err)
	}
	cases := []struct {
		uri  string
		want Article
	}{
		{"/tech/post-a/", Article{Lang: "ko", Section: "tech", Slug: "post-a"}},
		{"/tech/post-a/index.html", Article{Lang: "ko", Section: "tech", Slug: "post-a"}},
		{"/tech/post-a.md", Article{Lang: "ko", Section: "tech", Slug: "post-a", MD: true}},
		{"/en/tech/post-a/", Article{Lang: "en", Section: "tech", Slug: "post-a"}},
		{"/en/tech/post-a.md", Article{Lang: "en", Section: "tech", Slug: "post-a", MD: true}},
	}
	for _, c := range cases {
		got, ok := urls[c.uri]
		if !ok || got != c.want {
			t.Errorf("urls[%q] = (%+v, %v), want %+v", c.uri, got, ok, c.want)
		}
	}
	if _, ok := urls["/ko/tech/post-a/"]; ok {
		t.Errorf("root-served default language must not register a /ko/ path")
	}
	if len(urls) != 9 {
		t.Errorf("len(urls) = %d, want 9 (3 posts x 3 paths)", len(urls))
	}
}
