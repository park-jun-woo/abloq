//ff:func feature=quest type=parser control=sequence
//ff:what proseLines 검증 — 코드 펜스 안 라인과 펜스 구분 라인 제외, 나머지 본문 라인 보존
package translation

import (
	"strings"
	"testing"
)

func TestProseLines(t *testing.T) {
	md := "---\ntitle: x\n---\n\nkeep\n\n```sh\ndrop\n```\n\nkeep too\n"
	joined := strings.Join(proseLines(docOf(t, "en", md)), "\n")
	if strings.Contains(joined, "drop") || strings.Contains(joined, "```") {
		t.Errorf("fence content leaked: %q", joined)
	}
	if !strings.Contains(joined, "keep") || !strings.Contains(joined, "keep too") {
		t.Errorf("prose lost: %q", joined)
	}
}
