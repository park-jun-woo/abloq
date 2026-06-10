//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what 가짜 소스의 List — 설정된 키 목록과 에러를 그대로 반환
package cflog

func (f fakeSource) List(prefix, afterKey string) ([]string, error) {
	return f.keys, f.listErr
}
