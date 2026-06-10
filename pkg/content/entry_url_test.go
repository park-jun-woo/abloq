//ff:func feature=content type=generator control=sequence
//ff:what entryURL이 baseURL 끝 슬래시를 정리하고 루트 서빙 기본 언어는 언어 세그먼트를 생략하는지 검증
package content

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestEntryURL(t *testing.T) {
	b := &blogyaml.Blog{Languages: []string{"ko", "en"}}
	b.Site.BaseURL = "https://x.example.com/"
	b.Site.DefaultLangInSubdir = true
	if got := entryURL(b, "ko", "tech", "hello"); got != "https://x.example.com/ko/tech/hello/" {
		t.Errorf("subdir default lang = %q", got)
	}
	b.Site.DefaultLangInSubdir = false
	if got := entryURL(b, "ko", "tech", "hello"); got != "https://x.example.com/tech/hello/" {
		t.Errorf("root-served default lang = %q", got)
	}
	if got := entryURL(b, "en", "tech", "hello"); got != "https://x.example.com/en/tech/hello/" {
		t.Errorf("non-default lang = %q", got)
	}
}
