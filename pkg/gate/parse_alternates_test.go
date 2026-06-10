//ff:func feature=gate type=parser control=sequence
//ff:what parseAlternates가 rel=alternate 링크의 hreflang→href 맵을 추출하고 무관 링크를 무시하는지 검증
package gate

import "testing"

func TestParseAlternates(t *testing.T) {
	html := `<head>
<link rel="alternate" hreflang="ko" href="https://e.com/ko/a/">
<link rel="alternate" hreflang="en" href="https://e.com/en/a/">
<link rel="stylesheet" href="/main.css">
</head>`
	alts := parseAlternates(html)
	if len(alts) != 2 {
		t.Fatalf("want 2 alternates, got %v", alts)
	}
	if alts["ko"] != "https://e.com/ko/a/" || alts["en"] != "https://e.com/en/a/" {
		t.Errorf("alts = %v", alts)
	}
	if got := parseAlternates("<head></head>"); len(got) != 0 {
		t.Errorf("no links: want empty, got %v", got)
	}
}
