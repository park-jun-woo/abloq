//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what BuildURLMap이 blog.yaml 없는 저장소를 에러로 거부하는지 검증
package cflog

import "testing"

func TestBuildURLMapError(t *testing.T) {
	_, b := writeRepoFixture(t)
	if _, err := BuildURLMap(t.TempDir(), b); err == nil {
		t.Error("repo without blog.yaml accepted")
	}
}
