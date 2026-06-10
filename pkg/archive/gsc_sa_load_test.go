//ff:func feature=archive type=client control=sequence
//ff:what loadServiceAccount가 GSC_SA_JSON 인라인 우선·GSC_SA_JSON_PATH 폴백·필드 누락과 미설정 거부를 지키는지 검증
package archive

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadServiceAccount(t *testing.T) {
	saJSON := saJSONFixture(t)

	t.Setenv("GSC_SA_JSON", saJSON)
	t.Setenv("GSC_SA_JSON_PATH", "")
	sa, err := loadServiceAccount()
	if err != nil || sa.ClientEmail == "" || sa.PrivateKey == "" {
		t.Fatalf("inline load: %+v, %v", sa, err)
	}

	path := filepath.Join(t.TempDir(), "sa.json")
	if err := os.WriteFile(path, []byte(saJSON), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	t.Setenv("GSC_SA_JSON", "")
	t.Setenv("GSC_SA_JSON_PATH", path)
	if _, err := loadServiceAccount(); err != nil {
		t.Errorf("path load: %v", err)
	}

	t.Setenv("GSC_SA_JSON_PATH", filepath.Join(t.TempDir(), "missing.json"))
	if _, err := loadServiceAccount(); err == nil {
		t.Error("missing file must fail")
	}

	t.Setenv("GSC_SA_JSON", `{"client_email":"x@y"}`)
	if _, err := loadServiceAccount(); err == nil {
		t.Error("missing private_key must fail")
	}

	t.Setenv("GSC_SA_JSON", "not json")
	if _, err := loadServiceAccount(); err == nil {
		t.Error("invalid JSON must fail")
	}

	t.Setenv("GSC_SA_JSON", "")
	t.Setenv("GSC_SA_JSON_PATH", "")
	if _, err := loadServiceAccount(); err == nil {
		t.Error("unset credentials must fail")
	}
}
