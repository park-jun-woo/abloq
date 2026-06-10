//ff:func feature=gate type=rule control=sequence topic=evidence
//ff:what [citation-exists] Offline 타깃은 네트워크 검증 전체를 스킵하고 영수증도 만들지 않는지 검증
package gate

import "testing"

func TestRuleCitationExistsOffline(t *testing.T) {
	b := loadGateBlog(t)
	a := artFromContent(t, b, "Claim per [x](http://127.0.0.1:1/dead).\n")
	dir := t.TempDir()
	tgt := NewTarget(dir, b, []*Article{a})
	tgt.Offline = true
	if got := ruleCitationExists(tgt); len(got) != 0 {
		t.Errorf("offline target must skip citation-exists, got %v", got)
	}
	if got := loadReceipts(dir); len(got) != 0 {
		t.Errorf("offline run must not write receipts, got %v", got)
	}
}
