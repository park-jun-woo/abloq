//ff:func feature=insight type=parser control=sequence
//ff:what Load 검증 — 정상 픽스처 로드, 파일 없음 IO 에러, 검증 에러·경고 분리 반환
package insight

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	ins, errs, warns, err := Load(filepath.Join("testdata", "parkjunwoo", "content", "en", "tech", "ratchet-pattern.insight.yaml"))
	if err != nil || len(errs) != 0 || len(warns) != 0 {
		t.Fatalf("want clean load, got errs=%v warns=%v err=%v", errs, warns, err)
	}
	if ins.Section != "tech" || len(ins.Claims) != 7 {
		t.Errorf("want section tech with 7 claims, got %+v", ins)
	}
	if _, _, _, err := Load(filepath.Join(t.TempDir(), "absent.insight.yaml")); err == nil {
		t.Errorf("want IO error for missing file, got nil")
	}
	bad := filepath.Join(t.TempDir(), "bad.insight.yaml")
	if err := os.WriteFile(bad, []byte("topic: t\nclaims:\n  - id: a\n    text: x\n    kind: bogus\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	ins, errs, warns, err = Load(bad)
	if err != nil || ins != nil || len(errs) == 0 || len(warns) == 0 {
		t.Errorf("want nil insight with kind error and anchors warning, got ins=%v errs=%v warns=%v err=%v", ins, errs, warns, err)
	}
	unparsable := filepath.Join(t.TempDir(), "unknown.insight.yaml")
	if err := os.WriteFile(unparsable, []byte("topic: t\nbogus: 1\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	ins, errs, warns, err = Load(unparsable)
	if err != nil || ins != nil || len(errs) == 0 || len(warns) != 0 {
		t.Errorf("want parse diagnostics for unknown key, got ins=%v errs=%v warns=%v err=%v", ins, errs, warns, err)
	}
}
