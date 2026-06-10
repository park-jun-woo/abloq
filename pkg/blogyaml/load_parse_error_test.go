//ff:func feature=blogyaml type=parser control=sequence topic=diagnostics
//ff:what Load가 파싱 단계에서 실패한 파일(unknown key)을 검증 없이 파싱 진단으로 반환하는지 검증
package blogyaml

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadParseError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "blog.yaml")
	src := []byte("languages: [ko]\nsections: [tech]\nbogus_key: 1\n")
	if err := os.WriteFile(path, src, 0o644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	b, diags, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if b != nil {
		t.Errorf("want nil Blog on parse failure, got %+v", b)
	}
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %d: %v", len(diags), diags)
	}
	if diags[0].Rule != "unknown-key" {
		t.Errorf("want rule unknown-key, got %q", diags[0].Rule)
	}
}
