//ff:func feature=scan type=parser control=sequence topic=evidence
//ff:what citeURLs가 cite 목록의 URL만 순서대로 사상하는지 검증
package evidence

import "testing"

func TestCiteURLs(t *testing.T) {
	urls := citeURLs([]cite{
		{Lang: "ko", Section: "tech", Slug: "a", URL: "https://a.example/1"},
		{Lang: "ko", Section: "tech", Slug: "b", URL: "https://b.example/2"},
	})
	if len(urls) != 2 || urls[0] != "https://a.example/1" || urls[1] != "https://b.example/2" {
		t.Errorf("citeURLs = %v", urls)
	}
}
