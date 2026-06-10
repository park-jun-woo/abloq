//ff:func feature=gate type=parser control=iteration dimension=1
//ff:what fmKeyLine이 키의 파일 라인 번호(펜스 보정 +2)를 반환하고 부재 시 1을 반환하는지 검증
package gate

import "testing"

func TestFMKeyLine(t *testing.T) {
	fm := "title: x\ndate: 2026-01-01\nlastmod: 2026-01-02"
	cases := []struct {
		name, key string
		want      int
	}{
		{"first key", "title", 2},
		{"third key", "lastmod", 4},
		{"absent", "tags", 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := fmKeyLine(fm, tc.key); got != tc.want {
				t.Errorf("fmKeyLine(%s) = %d, want %d", tc.key, got, tc.want)
			}
		})
	}
}
