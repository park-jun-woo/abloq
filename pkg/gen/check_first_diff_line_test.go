//ff:func feature=gen type=rule control=sequence topic=drift
//ff:what firstDiffLine 케이스 하나를 실행해 라인 번호와 기대/실제 줄 반환값을 검증
package gen

import "testing"

func checkFirstDiffLine(t *testing.T, want, got string, wantLine int, wantW, wantG string) {
	t.Helper()
	line, w, g := firstDiffLine([]byte(want), []byte(got))
	if line != wantLine || w != wantW || g != wantG {
		t.Errorf("firstDiffLine(%q, %q) = (%d, %q, %q), want (%d, %q, %q)",
			want, got, line, w, g, wantLine, wantW, wantG)
	}
}
