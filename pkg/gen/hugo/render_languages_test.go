//ff:func feature=gen type=generator control=sequence
//ff:what renderLanguages가 선언 순서대로 weight를 매기며 [languages.코드] 블록을 내는지 검증
package hugo

import "testing"

func TestRenderLanguages(t *testing.T) {
	got := renderLanguages([]string{"ko", "en"})
	want := "\n[languages.ko]\nlanguageCode = \"ko\"\ncontentDir = \"content/ko\"\nweight = 1\n" +
		"\n[languages.en]\nlanguageCode = \"en\"\ncontentDir = \"content/en\"\nweight = 2\n"
	if got != want {
		t.Errorf("renderLanguages = %q, want %q", got, want)
	}
}
