//ff:func feature=blogyaml type=parser control=sequence
//ff:what Load가 존재하지 않는 파일에 대해 IO 에러를 반환하는지 검증
package blogyaml

import (
	"path/filepath"
	"testing"
)

func TestLoadMissingFile(t *testing.T) {
	b, diags, err := Load(filepath.Join(t.TempDir(), "no-such-blog.yaml"))
	if err == nil {
		t.Fatalf("want IO error for missing file, got nil (blog=%+v, diags=%v)", b, diags)
	}
	if b != nil {
		t.Errorf("want nil Blog on IO error, got %+v", b)
	}
	if diags != nil {
		t.Errorf("want nil diagnostics on IO error, got %v", diags)
	}
}
