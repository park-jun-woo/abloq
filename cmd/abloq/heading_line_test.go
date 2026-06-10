//ff:func feature=init type=generator control=sequence
//ff:what headingLine이 언어 선언 순서대로 현지화 헤딩 쌍을 렌더하는지 검증
package main

import "testing"

func TestHeadingLine(t *testing.T) {
	got := headingLine([]string{"ko", "en", "xx"})
	want := `ko: "출처", en: "Sources", xx: "Sources"`
	if got != want {
		t.Errorf("headingLine = %q, want %q", got, want)
	}
}
