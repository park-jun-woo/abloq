//ff:func feature=blogyaml type=parser control=sequence
//ff:what 키 경로 두 조각을 "."으로 연결 (prefix가 비면 key만)
package blogyaml

// joinPath joins a key path prefix and a key with a dot separator.
func joinPath(prefix, key string) string {
	if prefix == "" {
		return key
	}
	return prefix + "." + key
}
