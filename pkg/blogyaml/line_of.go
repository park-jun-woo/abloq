//ff:func feature=blogyaml type=parser control=sequence topic=diagnostics
//ff:what 키 경로의 라인 번호를 조회, 키가 파일에 없으면(기본값 주입 등) 1을 반환
package blogyaml

// lineOf returns the source line of a key path, falling back to line 1 for absent keys.
func lineOf(idx lineIndex, path string) int {
	if line, ok := idx[path]; ok {
		return line
	}
	return 1
}
