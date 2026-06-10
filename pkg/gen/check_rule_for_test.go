//ff:func feature=gen type=rule control=sequence topic=drift
//ff:what ruleFor 케이스 하나를 실행해 룰ID 매핑 결과를 검증
package gen

import "testing"

func checkRuleFor(t *testing.T, path, want string) {
	t.Helper()
	if got := ruleFor(path); got != want {
		t.Errorf("ruleFor(%q) = %q, want %q", path, got, want)
	}
}
