//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what env 값 조회 — 비어 있으면 기본값 (엔진 베이스 URL·모델 오버라이드의 공통 입구)
package citation

import "os"

// envOr returns the environment value of key, or def when unset/empty.
func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
