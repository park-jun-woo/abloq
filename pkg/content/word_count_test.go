//ff:func feature=content type=parser control=sequence
//ff:what wordCount가 공백 구분 토큰 수를 세고 빈 본문은 0을 돌려주는지 검증
package content

import "testing"

func TestWordCount(t *testing.T) {
	if got := wordCount("one two  three\nfour"); got != 4 {
		t.Errorf("wordCount = %d, want 4", got)
	}
	if got := wordCount("  \n\t"); got != 0 {
		t.Errorf("blank body = %d, want 0", got)
	}
}
