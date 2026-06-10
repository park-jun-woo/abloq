//ff:func feature=archive type=client control=sequence
//ff:what GSC 실행당 제출 상한 — GSC_DAILY_QUOTA env, 기본 200 (Indexing API 일일 쿼터)
package archive

import (
	"os"
	"strconv"
)

// gscQuota reads the per-run submission cap. Invalid or non-positive values
// fall back to the Indexing API default daily quota of 200.
func gscQuota() int {
	v, err := strconv.Atoi(os.Getenv("GSC_DAILY_QUOTA"))
	if err != nil || v <= 0 {
		return 200
	}
	return v
}
