//ff:func feature=init type=parser control=sequence
//ff:what ListDir가 중첩 디렉토리의 파일 경로를 정렬 순서로 모으고 누락 디렉토리에 에러를 내는지 검증
package scaffold

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestListDir(t *testing.T) {
	fsys := fstest.MapFS{
		"b.txt":         {Data: []byte("b")},
		"a/one.txt":     {Data: []byte("1")},
		"a/sub/two.txt": {Data: []byte("2")},
	}
	got, err := ListDir(fsys, ".")
	if err != nil {
		t.Fatalf("ListDir: %v", err)
	}
	want := []string{"a/one.txt", "a/sub/two.txt", "b.txt"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ListDir = %v, want %v", got, want)
	}
	if _, err := ListDir(fsys, "missing"); err == nil {
		t.Error("ListDir(missing) expected error, got nil")
	}
	if _, err := ListDir(failSubFS{fsys}, "."); err == nil {
		t.Error("ListDir with failing subdirectory expected error, got nil")
	}
}
