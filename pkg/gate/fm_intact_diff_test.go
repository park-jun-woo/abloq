//ff:func feature=gate type=rule control=iteration dimension=1 topic=baseline
//ff:what fmIntactDiff가 lastmod 변경·우측 공백·말미 빈 줄을 허용하고 그 외 변경을 검출하는지 검증
package gate

import "testing"

func TestFMIntactDiff(t *testing.T) {
	cases := []struct {
		name, orig, neu string
		wantOK          bool
	}{
		{"identical", "title: x\ndate: 1", "title: x\ndate: 1", true},
		{"lastmod changed", "title: x\nlastmod: 1", "title: x\nlastmod: 2", true},
		{"trailing space tolerated", "title: x", "title: x  ", true},
		{"trailing blank tolerated", "title: x", "title: x\n\n", true},
		{"value changed", "title: x", "title: y", false},
		{"line added", "title: x", "title: x\ndraft: true", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, ok := fmIntactDiff(tc.orig, tc.neu); ok != tc.wantOK {
				t.Errorf("fmIntactDiff ok = %v, want %v", ok, tc.wantOK)
			}
		})
	}
}
