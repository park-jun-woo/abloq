//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what DirSource.List가 없는 루트 디렉토리를 에러로 거부하는지 검증
package cflog

import "testing"

func TestDirListError(t *testing.T) {
	if _, err := (DirSource{Root: "testdata/does-not-exist"}).List("", ""); err == nil {
		t.Error("missing root accepted")
	}
}
