//ff:func feature=archive type=client control=sequence
//ff:what SA 자격 로드 — GSC_SA_JSON(인라인) 우선, 없으면 GSC_SA_JSON_PATH(파일), 둘 다 없으면 에러
package archive

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// loadServiceAccount reads the service-account credentials. Credentials live
// only in the backend environment — never in receipts, never with agents.
func loadServiceAccount() (*serviceAccount, error) {
	raw := []byte(os.Getenv("GSC_SA_JSON"))
	if len(raw) == 0 {
		path := os.Getenv("GSC_SA_JSON_PATH")
		if path == "" {
			return nil, errors.New("neither GSC_SA_JSON nor GSC_SA_JSON_PATH is set")
		}
		data, err := os.ReadFile(path)
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
