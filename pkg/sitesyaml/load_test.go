//ff:func feature=sitesyaml type=parser control=sequence
//ff:what Load가 골든 sites.yaml을 진단 0건으로 통과시키고, 검증 실패는 진단으로, 파일 미존재는 IO 에러로 돌려주는지 검증
package sitesyaml

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	s, diags, err := Load(filepath.Join("testdata", "valid", "sites.yaml"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(diags) != 0 {
		t.Fatalf("want 0 diagnostics, got %v", diags)
	}
	if len(s.Sites) != 2 {
		t.Fatalf("want 2 sites, got %d", len(s.Sites))
	}
	first := s.Sites[0]
	if first.Name != "parkjunwoo" || first.RepoPath != "/blogs/parkjunwoo" || !first.Active {
		t.Errorf("first site = %+v", first)
	}
	if first.QueueExport.RepoURL == "" || first.GSC.SAJSONPath != "/secrets/gsc-sa.json" {
		t.Errorf("nested groups not decoded: %+v", first)
	}
	if first.CFLogSource == "" || first.IndexNowKey == "" {
		t.Errorf("optional scalars not decoded: %+v", first)
	}
	if s.Sites[1].Active {
		t.Error("second site declares active: false")
	}

	bad := filepath.Join(t.TempDir(), "sites.yaml")
	if err := os.WriteFile(bad, []byte("sites:\n  - name: UPPER\n    repo_path: relative\n"), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	_, diags, err = Load(bad)
	if err != nil || len(diags) == 0 {
		t.Errorf("invalid file must come back as diagnostics: %v, %v", diags, err)
	}

	unknown := filepath.Join(t.TempDir(), "sites.yaml")
	if err := os.WriteFile(unknown, []byte("sites:\n  - nam: a\n"), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	s, diags, err = Load(unknown)
	if err != nil || s != nil || len(diags) == 0 || diags[0].Rule != "unknown-key" {
		t.Errorf("parse failure must stop before validation: %v, %v, %v", s, diags, err)
	}

	if _, _, err := Load(filepath.Join(t.TempDir(), "missing.yaml")); err == nil {
		t.Error("missing file must be an IO error")
	}
}
