//ff:func feature=gen type=rule control=iteration dimension=1 topic=drift
//ff:what firstDiffLine이 첫 불일치 라인 번호·양쪽 줄을 반환하고 동일 입력·길이 차이를 처리하는지 검증
package gen

import "testing"

func TestFirstDiffLine(t *testing.T) {
	cases := []struct {
		name      string
		want, got string
		wantLine  int
		wantW     string
		wantG     string
	}{
		{"equal", "a\nb\n", "a\nb\n", 0, "", ""},
		{"middle line differs", "a\nb\nc\n", "a\nx\nc\n", 2, "b", "x"},
		{"got truncated", "a\nb\n", "a\n", 2, "b", ""},
		{"got longer", "a\n", "a\nextra\n", 2, "", "extra"},
		{"first line differs", "a\n", "z\n", 1, "a", "z"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkFirstDiffLine(t, tc.want, tc.got, tc.wantLine, tc.wantW, tc.wantG)
		})
	}
}
