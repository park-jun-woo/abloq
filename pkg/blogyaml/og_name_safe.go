//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what variant name의 URL-safe 판정 — 소문자/숫자/하이픈/언더스코어만 허용 (드래프트 파일명에 그대로 들어간다)
package blogyaml

// ogNameSafe reports whether name is URL/filename-safe: lowercase letters,
// digits, '-' and '_' only. Empty names are unsafe.
func ogNameSafe(name string) bool {
	if name == "" {
		return false
	}
	for _, r := range name {
		if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' && r != '_' {
			return false
		}
	}
	return true
}
