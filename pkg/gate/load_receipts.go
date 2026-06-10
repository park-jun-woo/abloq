//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what 영수증 파일(.abloq/citation-receipts.json) 로드 — 없거나 깨졌으면 빈 맵(전수 재검증)
package gate

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// loadReceipts reads the citation-verification receipt cache under dir.
// A missing or corrupt file simply means every URL gets re-verified.
func loadReceipts(dir string) map[string]receipt {
	m := map[string]receipt{}
	data, err := os.ReadFile(filepath.Join(dir, ".abloq", "citation-receipts.json"))
	if err != nil {
		return m
	}
	if json.Unmarshal(data, &m) != nil || m == nil {
		return map[string]receipt{}
	}
	return m
}
