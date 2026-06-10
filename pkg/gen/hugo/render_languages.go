//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what 언어 목록을 blog.yaml 선언 순서대로 hugo.toml [languages.코드] 블록(languageCode/contentDir/weight)으로 렌더
package hugo

import (
	"fmt"
	"strings"
)

// renderLanguages emits one [languages.<code>] block per language;
// weight follows the declaration order (1-based) so hreflang alternates stay stable.
func renderLanguages(langs []string) string {
	var sb strings.Builder
	for i, lang := range langs {
		fmt.Fprintf(&sb, "\n[languages.%s]\n", lang)
		fmt.Fprintf(&sb, "languageCode = \"%s\"\n", lang)
		fmt.Fprintf(&sb, "contentDir = \"content/%s\"\n", lang)
		fmt.Fprintf(&sb, "weight = %d\n", i+1)
	}
	return sb.String()
}
