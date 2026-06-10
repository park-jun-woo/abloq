//ff:func feature=blogyaml type=schema control=sequence
//ff:what URLLang이 기본 언어 루트 서빙(false)·서브디렉토리(true)·비기본 언어를 올바른 세그먼트로 푸는지 검증
package blogyaml

import "testing"

func TestURLLang(t *testing.T) {
	b := &Blog{Languages: []string{"en", "ko"}}
	b.Site.DefaultLangInSubdir = true
	if got := b.URLLang("en"); got != "en" {
		t.Errorf("subdir=true default lang: want en, got %q", got)
	}
	b.Site.DefaultLangInSubdir = false
	if got := b.URLLang("en"); got != "" {
		t.Errorf("subdir=false default lang: want empty, got %q", got)
	}
	if got := b.URLLang("ko"); got != "ko" {
		t.Errorf("subdir=false other lang: want ko, got %q", got)
	}
	empty := &Blog{}
	if got := empty.URLLang("en"); got != "en" {
		t.Errorf("no languages: want en, got %q", got)
	}
}
