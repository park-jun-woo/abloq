//ff:func feature=quest type=parser control=sequence
//ff:what headingLevels 검증 — 자유 헤딩 포함 레벨 시퀀스 추출, 코드 펜스 안 # 라인은 헤딩이 아님
package translation

import (
	"fmt"
	"testing"
)

func TestHeadingLevels(t *testing.T) {
	md := "---\ntitle: x\n---\n\n## A\n\n### B\n\n```sh\n# not a heading\n```\n\n## C\n"
	got := fmt.Sprint(headingLevels(docOf(t, "en", md)))
	if got != "[2 3 2]" {
		t.Errorf("levels = %s, want [2 3 2]", got)
	}
}
