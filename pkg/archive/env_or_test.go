//ff:func feature=archive type=client control=sequence
//ff:what envOr가 설정된 env 값을 쓰고, 비어 있으면 기본값으로 떨어지는지 검증
package archive

import "testing"

func TestEnvOr(t *testing.T) {
	t.Setenv("ABLOQ_TEST_ENV_OR", "override")
	if got := envOr("ABLOQ_TEST_ENV_OR", "default"); got != "override" {
		t.Errorf("envOr set = %q, want override", got)
	}
	t.Setenv("ABLOQ_TEST_ENV_OR", "")
	if got := envOr("ABLOQ_TEST_ENV_OR", "default"); got != "default" {
		t.Errorf("envOr empty = %q, want default", got)
	}
}
