//ff:func feature=gate type=output control=sequence topic=evidence
//ff:what 영수증 맵을 .abloq/citation-receipts.json에 기록 — 디렉토리 자동 생성, 들여쓴 JSON
package gate

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// saveReceipts writes the citation-verification receipt cache under dir.
func saveReceipts(dir string, m map[string]receipt) error {
	path := filepath.Join(dir, ".abloq", "citation-receipts.json")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, _ := json.MarshalIndent(m, "", "  ") // a map[string]receipt cannot fail to marshal
	return os.WriteFile(path, data, 0o644)
}
