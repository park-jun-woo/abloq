//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what claimTexts가 검출 주장 텍스트를 모으고 exclude 해시는 제외하며 nil exclude는 전건인지 검증
package common

import (
	"testing"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

func TestClaimTexts(t *testing.T) {
	c := consFixture(t, claimBaseMD, claimBaseMD)
	d := c.Target.Articles[0].Doc
	all := claimTexts(d, nil)
	if len(all) != 2 {
		t.Fatalf("want 2 claims, got %v", all)
	}
	excluded := claimTexts(d, map[string]bool{agate.HashText(all[0]): true})
	if len(excluded) != 1 || excluded[0] == all[0] {
		t.Errorf("exclusion failed: %v", excluded)
	}
}
