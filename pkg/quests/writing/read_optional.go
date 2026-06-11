//ff:func feature=quest type=parser control=sequence
//ff:what 루트 기준 상대 경로 파일을 읽어 본문 반환 — 경로 공백·파일 부재는 빈 문자열 (부재 판정은 review-record 룰 몫)
package writing

import (
	"os"
	"path/filepath"
)

// readOptional reads a root-relative artifact file, returning "" when the
// path is empty or the file is absent: absence is a gate finding (the
// review-record rule), not a Prepare error.
func readOptional(root, rel string) string {
	if rel == "" {
		return ""
	}
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		return ""
	}
	return string(data)
}
