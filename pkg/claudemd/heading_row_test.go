//ff:func feature=claudemd type=generator control=sequence
//ff:what headingRow가 언어 순서대로 셀을 채우고 미정의 언어를 빈 칸으로 두는지 검증
package claudemd

import "testing"

func TestHeadingRow(t *testing.T) {
	got := headingRow("sources", []string{"ko", "en", "ja"}, map[string]string{"ko": "출처", "en": "Sources"})
	want := "| sources | 출처 | Sources |  |\n"
	if got != want {
		t.Errorf("headingRow = %q, want %q", got, want)
	}
}
