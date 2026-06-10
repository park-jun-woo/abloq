//ff:func feature=gate type=rule control=sequence topic=evidence
//ff:what [citation-exists] 영수증 캐시 — 24h 내 ok 영수증은 재검증 생략, 만료 영수증은 재검증(RETRY), retry는 캐시 미기록
package gate

import (
	"testing"
	"time"
)

func TestRuleCitationExistsCache(t *testing.T) {
	b := loadGateBlog(t)
	url := "http://127.0.0.1:1/dead" // refused if ever fetched
	dir := t.TempDir()
	a := artFromContent(t, b, "Claim per [x]("+url+").\n")
	if err := saveReceipts(dir, map[string]receipt{url: {CheckedAt: time.Now(), Verdict: "ok"}}); err != nil {
		t.Fatal(err)
	}
	tgt := NewTarget(dir, b, []*Article{a})
	checkDiags(t, ruleCitationExists(tgt), 0, "", "")
	stale := map[string]receipt{url: {CheckedAt: time.Now().Add(-25 * time.Hour), Verdict: "ok"}}
	if err := saveReceipts(dir, stale); err != nil {
		t.Fatal(err)
	}
	checkDiags(t, ruleCitationExists(tgt), 1, "citation-exists", "RETRY")
	if got := loadReceipts(dir)[url].Verdict; got != "ok" {
		t.Errorf("a retry verdict must not overwrite the receipt, got %q", got)
	}
}
