//ff:func feature=visibility type=generator control=sequence topic=report
//ff:what Build가 nil 입력 슬라이스를 빈 슬라이스로 정규화(JSON []·null 아님)하고 빈 리포트도 렌더되는지 검증
package report

import "testing"

func TestBuildNilSlices(t *testing.T) {
	r := Build(Input{YM: "2026-04"})
	if r.Queue == nil || r.UnknownBots == nil {
		t.Error("nil inputs must normalize to empty slices (JSON [] not null)")
	}
	if string(JSON(r)) == "" || r.Rows == nil {
		t.Error("an empty report must still render")
	}
}
