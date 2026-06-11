//ff:func feature=quest type=parser control=sequence
//ff:what imagePaths 검증 — 본문 이미지 경로 전수 추출, 일반 링크와 펜스 안 이미지 구문은 미포함
package translation

import (
	"fmt"
	"testing"
)

func TestImagePaths(t *testing.T) {
	md := "---\ntitle: x\n---\n\n![a](/images/a.png)\n\nsee [doc](/posts/doc/) and ![b](/images/b.webp)\n\n```md\n![c](/images/c.png)\n```\n"
	got := fmt.Sprint(imagePaths(docOf(t, "en", md)))
	if got != "[/images/a.png /images/b.webp]" {
		t.Errorf("paths = %s", got)
	}
}
