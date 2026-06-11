//ff:func feature=archive type=client control=sequence
//ff:what loadServiceAccount가 호출자 인라인>경로 우선과 env(GSC_SA_JSON>GSC_SA_JSON_PATH) fallback, 사이트 경로의 전역 인라인 우선, 필드 누락·미설정 거부를 지키는지 검증
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
	sa, err := loadServiceAccount("", "")
	if err != nil || sa.ClientEmail == "" || sa.PrivateKey == "" {
		t.Fatalf("env inline load: %+v, %v", sa, err)
	}

	path := filepath.Join(t.TempDir(), "sa.json")
	if err := os.WriteFile(path, []byte(saJSON), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	t.Setenv("GSC_SA_JSON", "")
	t.Setenv("GSC_SA_JSON_PATH", path)
	if _, err := loadServiceAccount("", ""); err != nil {
		t.Errorf("env path load: %v", err)
	}

	t.Setenv("GSC_SA_JSON", "not json — a caller path must win over this")
	t.Setenv("GSC_SA_JSON_PATH", "")
	if _, err := loadServiceAccount("", path); err != nil {
		t.Errorf("caller path must win over env inline: %v", err)
	}

	if _, err := loadServiceAccount(saJSON, filepath.Join(t.TempDir(), "missing.json")); err != nil {
		t.Errorf("caller inline must win over caller path: %v", err)
	}

	t.Setenv("GSC_SA_JSON", "")
	if _, err := loadServiceAccount("", filepath.Join(t.TempDir(), "missing.json")); err == nil {
		t.Error("missing caller file must fail")
	}

	t.Setenv("GSC_SA_JSON_PATH", filepath.Join(t.TempDir(), "missing.json"))
	if _, err := loadServiceAccount("", ""); err == nil {
		t.Error("missing env file must fail")
	}

	t.Setenv("GSC_SA_JSON", `{"client_email":"x@y"}`)
	t.Setenv("GSC_SA_JSON_PATH", "")
	if _, err := loadServiceAccount("", ""); err == nil {
		t.Error("missing private_key must fail")
	}

	t.Setenv("GSC_SA_JSON", "not json")
	if _, err := loadServiceAccount("", ""); err == nil {
		t.Error("invalid JSON must fail")
	}

	t.Setenv("GSC_SA_JSON", "")
	t.Setenv("GSC_SA_JSON_PATH", "")
	if _, err := loadServiceAccount("", ""); err == nil {
		t.Error("unset credentials must fail")
	}
}
