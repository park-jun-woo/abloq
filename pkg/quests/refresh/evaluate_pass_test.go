//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what 게이트 전 룰 통과 검증 — lastmod 전진+의미 diff+주장 교체(동수)의 정상 갱신을 reins 레벨집계로 평가해 PASS
package refresh

import (
	"strings"
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestEvaluatePass(t *testing.T) {
	root := writeInstance(t)
	refreshed := baseArticleMD
	refreshed = strings.Replace(refreshed, "lastmod: 2026-06-02", "lastmod: 2026-06-09", 1)
	refreshed = strings.Replace(refreshed,
		"This stale body sentence still describes the situation as of early 2025 in vendor terms.",
		"The refreshed body now reflects the mid 2026 landscape with current vendor guidance and revised context.", 1)
	refreshed = strings.Replace(refreshed,
		"Throughput grew 40% in 2025 per the vendor study.",
		"Throughput grew 55% in 2026 per the vendor study.", 1)
	writeFile(t, root, "content/en/posts/a.md", refreshed)
	v := rgate.Evaluate(Definition{}.Rules(), subWith(t, root))
	if v.Outcome != quest.OutPass {
		t.Fatalf("Outcome = %s, facts = %+v — want PASS", v.Outcome, v.Facts)
	}
}
