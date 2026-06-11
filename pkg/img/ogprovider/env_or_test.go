//ff:func feature=image type=client control=sequence
//ff:what envOr 검증 — 설정된 env 값 우선, 미설정·빈 문자열은 기본값
package ogprovider

import "testing"

func TestEnvOr(t *testing.T) {
	const key = "ABLOQ_TEST_ENV_OR"

	t.Setenv(key, "from-env")
	if got := envOr(key, "fallback"); got != "from-env" {
		t.Errorf("set: %q, want from-env", got)
	}

	t.Setenv(key, "")
	if got := envOr(key, "fallback"); got != "fallback" {
		t.Errorf("empty: %q, want fallback", got)
	}

	if got := envOr("ABLOQ_TEST_ENV_OR_UNSET", "fallback"); got != "fallback" {
		t.Errorf("unset: %q, want fallback", got)
	}
}
