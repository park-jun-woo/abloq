//ff:type feature=cli type=output topic=diagnostics
//ff:what 항상 쓰기에 실패하는 io.Writer — 출력 에러 경로 테스트용
package main

import "errors"

// errWriter always fails, to exercise output error paths.
type errWriter struct{}

func (errWriter) Write([]byte) (int, error) { return 0, errors.New("write failed") }
