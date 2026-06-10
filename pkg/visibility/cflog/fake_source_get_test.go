//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what 가짜 소스의 Get — 설정된 에러(없으면 기본 에러)를 반환, 본문은 제공하지 않음
package cflog

import (
	"errors"
	"io"
)

func (f fakeSource) Get(key string) (io.ReadCloser, error) {
	if f.getErr != nil {
		return nil, f.getErr
	}
	return nil, errors.New("fakeSource: no body")
}
