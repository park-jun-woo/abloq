//ff:func feature=image type=client control=sequence
//ff:what env 값 조회 — 비어 있으면 기본값 (API 키 fallback·베이스 URL 오버라이드의 공통 입구)
package ogprovider

import "os"

// envOr returns the environment value of key, or def when unset/empty.
func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
