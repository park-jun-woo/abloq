//ff:func feature=scan type=parser control=iteration dimension=1 topic=cluster
//ff:what resolveTarget이 사이트 절대경로·baseURL·프래그먼트를 해석하고 외부·번역·비글 경로를 거부하는지 검증
package cluster

import "testing"

func TestResolveTarget(t *testing.T) {
	b := testBlog()
	cases := []struct {
		name   string
		target string
		want   string
		ok     bool
	}{
		{"site absolute", "/tech/hub/", "tech/hub", true},
		{"no trailing slash", "/tech/hub", "tech/hub", true},
		{"baseURL absolute", "https://t.example.com/tech/hub/", "tech/hub", true},
		{"fragment stripped", "/tech/hub/#sources", "tech/hub", true},
		{"query stripped", "/tech/hub/?ref=x", "tech/hub", true},
		{"external URL", "https://example.org/tech/hub/", "", false},
		{"relative", "img/photo.png", "", false},
		{"translation lang segment", "/en/tech/hub/", "", false},
		{"unknown section", "/tags/geo/", "", false},
		{"too deep", "/tech/hub/extra/", "", false},
		{"section page", "/tech/", "", false},
		{"root", "/", "", false},
	}
	for _, tc := range cases {
		got, ok := resolveTarget(b, "ko", tc.target)
		if got != tc.want || ok != tc.ok {
			t.Errorf("%s: resolveTarget(%q) = (%q, %v), want (%q, %v)", tc.name, tc.target, got, ok, tc.want, tc.ok)
		}
	}
	// Subdir-served default language requires its language segment.
	sub := testBlog()
	sub.Site.DefaultLangInSubdir = true
	if _, ok := resolveTarget(sub, "ko", "/tech/hub/"); ok {
		t.Error("subdir mode must reject a language-less path")
	}
	if got, ok := resolveTarget(sub, "ko", "/ko/tech/hub/"); !ok || got != "tech/hub" {
		t.Errorf("subdir mode: got (%q, %v)", got, ok)
	}
}
