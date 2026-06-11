//ff:func feature=archive type=client control=sequence
//ff:what SA 자격 로드 — 호출자 인자(인라인 JSON > 파일 경로)가 우선, 둘 다 비면 GSC_SA_JSON(인라인) > GSC_SA_JSON_PATH(파일) env fallback, 전부 없으면 에러
//ff:why 사이트 행 값 > 전역 env 재정의 — 사이트별 sa_json_path가 전역 인라인 GSC_SA_JSON에 가려지면 격리 누수다 (Phase020)
package archive

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// loadServiceAccount reads the service-account credentials. A caller-given
// value (site row: inline JSON wins over the file path) takes precedence;
// only when both arguments are empty does it fall back to the global
// GSC_SA_JSON / GSC_SA_JSON_PATH environment. Credentials live only in the
// backend environment — never in receipts, never with agents.
func loadServiceAccount(inline, path string) (*serviceAccount, error) {
	raw := []byte(inline)
	if len(raw) == 0 && path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		raw = data
	}
	if len(raw) == 0 {
		raw = []byte(os.Getenv("GSC_SA_JSON"))
	}
	if len(raw) == 0 {
		envPath := os.Getenv("GSC_SA_JSON_PATH")
		if envPath == "" {
			return nil, errors.New("neither GSC_SA_JSON nor GSC_SA_JSON_PATH is set")
		}
		data, err := os.ReadFile(envPath)
		if err != nil {
			return nil, err
		}
		raw = data
	}
	var sa serviceAccount
	if err := json.Unmarshal(raw, &sa); err != nil {
		return nil, fmt.Errorf("service account JSON: %w", err)
	}
	if sa.ClientEmail == "" || sa.PrivateKey == "" {
		return nil, errors.New("service account JSON lacks client_email or private_key")
	}
	return &sa, nil
}
