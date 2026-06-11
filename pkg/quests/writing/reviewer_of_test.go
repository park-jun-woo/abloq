//ff:func feature=quest type=parser control=sequence
//ff:what reviewerOf 검증 — reviewer: 라인 값 추출, 빈 값·부재는 빈 문자열
package writing

import "testing"

func TestReviewerOf(t *testing.T) {
	if got := reviewerOf("# REVIEW\nreviewer: agent-ctx-7\n"); got != "agent-ctx-7" {
		t.Errorf("got %q, want agent-ctx-7", got)
	}
	if got := reviewerOf("reviewer:\n- c1: addressed\n"); got != "" {
		t.Errorf("empty value: got %q, want empty", got)
	}
	if got := reviewerOf("no header here\n"); got != "" {
		t.Errorf("absent: got %q, want empty", got)
	}
}
