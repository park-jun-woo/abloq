//ff:func feature=content type=parser control=sequence
//ff:what internalLinks가 절대경로·baseURL 시작 링크만 세고 외부 URL·상대경로 이미지를 제외하는지 검증
package content

import "testing"

func TestInternalLinks(t *testing.T) {
	body := "[a](/tech/x/) [b](https://me.example.com/tech/y/) [c](https://other.example.com/) ![img](cover.webp)"
	if got := internalLinks(body, "https://me.example.com/"); got != 2 {
		t.Errorf("internalLinks = %d, want 2", got)
	}
	if got := internalLinks(body, ""); got != 1 {
		t.Errorf("internalLinks without baseURL = %d, want 1 (site-absolute only)", got)
	}
	if got := internalLinks("no links here", "https://me.example.com"); got != 0 {
		t.Errorf("linkless body = %d, want 0", got)
	}
}
