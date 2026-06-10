//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what 섹션 목록을 blog.yaml 선언 순서 그대로 hugo.toml [permalinks] 블록("/섹션/:slug/")으로 렌더
package hugo

import (
	"fmt"
	"strings"
)

// renderPermalinks emits one "/section/:slug/" permalink per declared section.
func renderPermalinks(sections []string) string {
	var sb strings.Builder
	sb.WriteString("\n[permalinks]\n")
	for _, section := range sections {
		fmt.Fprintf(&sb, "%s = \"/%s/:slug/\"\n", section, section)
	}
	return sb.String()
}
