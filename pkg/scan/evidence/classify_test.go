//ff:func feature=scan type=rule control=iteration dimension=1 topic=evidence
//ff:what classify 케이스 — 2xx/3xx ok, 404/410 hard, 그 외 4xx·5xx soft
package evidence

import "testing"

func TestClassify(t *testing.T) {
	cases := []struct {
		code int
		want string
	}{
		{200, "ok"}, {204, "ok"}, {301, "ok"},
		{404, "hard"}, {410, "hard"},
		{403, "soft"}, {429, "soft"}, {500, "soft"}, {503, "soft"},
	}
	for _, tc := range cases {
		if got := classify(tc.code); got != tc.want {
			t.Errorf("classify(%d) = %q, want %q", tc.code, got, tc.want)
		}
	}
}
