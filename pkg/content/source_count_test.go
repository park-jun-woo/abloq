//ff:func feature=content type=parser control=sequence
//ff:what sourceCount가 sources 헤딩 아래 리스트 항목(-, *, n.)만 세고 다음 헤딩에서 멈추며 빈 헤딩은 0인지 검증
package content

import "testing"

func TestSourceCount(t *testing.T) {
	body := "intro\n\n## 출처\n\n- one\n* two\n3. three\nprose line\n\n## 기타\n\n- not a source\n"
	if got := sourceCount(body, "출처"); got != 3 {
		t.Errorf("sourceCount = %d, want 3", got)
	}
	if got := sourceCount(body, ""); got != 0 {
		t.Errorf("empty heading = %d, want 0", got)
	}
	if got := sourceCount("no headings\n- list\n", "출처"); got != 0 {
		t.Errorf("missing section = %d, want 0", got)
	}
}
