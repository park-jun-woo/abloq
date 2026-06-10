//ff:func feature=visibility type=generator control=sequence topic=report
//ff:what prevCell이 전월 데이터 부재면 n/a, 있으면 값 문자열을 내는지 검증
package report

import "testing"

func TestPrevCell(t *testing.T) {
	if got := prevCell(Report{PrevHasData: false}, 42); got != "n/a" {
		t.Errorf("first month must read n/a, got %q", got)
	}
	if got := prevCell(Report{PrevHasData: true}, 42); got != "42" {
		t.Errorf("want 42, got %q", got)
	}
}
