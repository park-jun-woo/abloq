//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what 문자열 슬라이스 일치 검증 헬퍼 — 길이와 각 원소를 비교
package gate

import "testing"

func checkStrings(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
