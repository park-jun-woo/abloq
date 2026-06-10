//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what htmlTitles 케이스 — <title>, og:title(속성 순서 양방향), 둘 다, 없음
package gate

import "testing"

func TestHTMLTitles(t *testing.T) {
	cases := []struct {
		name, html string
		want       []string
	}{
		{"title only", `<html><head><title>Hello World</title></head></html>`, []string{"Hello World"}},
		{"og title property first", `<head><meta property="og:title" content="OG Hello"></head>`, []string{"OG Hello"}},
		{"og title content first", `<head><meta content="OG Hello" property="og:title"></head>`, []string{"OG Hello"}},
		{"both", `<head><title>T</title><meta property="og:title" content="O"></head>`, []string{"T", "O"}},
		{"none", `<head><meta name="description" content="x"></head>`, nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkStrings(t, htmlTitles(tc.html), tc.want)
		})
	}
}
