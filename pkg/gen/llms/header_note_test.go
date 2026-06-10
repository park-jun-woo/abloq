//ff:func feature=gen type=generator control=sequence
//ff:what headerNote가 저자 유무에 따라 "저자 — baseURL" 또는 baseURL만 내는지 검증
package llms

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestHeaderNote(t *testing.T) {
	b := &blogyaml.Blog{Site: blogyaml.Site{BaseURL: "https://x.com", Author: "A"}}
	if got := headerNote(b); got != "A — https://x.com" {
		t.Errorf("headerNote = %q, want %q", got, "A — https://x.com")
	}
	b.Site.Author = ""
	if got := headerNote(b); got != "https://x.com" {
		t.Errorf("headerNote without author = %q, want %q", got, "https://x.com")
	}
}
