//ff:func feature=queueio type=rule control=iteration dimension=1
//ff:what consumed 동기화 — exported 행마다 행→파일명 정방향 계산 후 파일 부재면 consumed (파일명 역파싱 금지)
package queueio

import (
	"os"
	"path/filepath"
)

// consumedIDs reports the exported rows whose queue file no longer exists in
// the fresh work clone — the agent deleted it in its consumption commit. The
// check is strictly forward (row → filename → stat): file names are not
// injective for hyphenated lang/section and must never be parsed back.
func consumedIDs(queueDir string, exported []Row) []int64 {
	ids := make([]int64, 0)
	for _, r := range exported {
		if _, err := os.Stat(filepath.Join(queueDir, Filename(r.Item))); os.IsNotExist(err) {
			ids = append(ids, r.ID)
		}
	}
	return ids
}
