//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what envOr가 설정된 env 값을 반환하고 미설정·빈 값이면 기본값인지 검증
package citation

import "testing"

func TestEnvOrCitation(t *testing.T) {
	t.Setenv("ABLOQ_TEST_ENV_OR", "set-value")
	if got := envOr("ABLOQ_TEST_ENV_OR", "default"); got != "set-value" {
		t.Errorf("set = %q", got)
	}
	t.Setenv("ABLOQ_TEST_ENV_OR", "")
	if got := envOr("ABLOQ_TEST_ENV_OR", "default"); got != "default" {
		t.Errorf("empty = %q, want default", got)
	}
}
