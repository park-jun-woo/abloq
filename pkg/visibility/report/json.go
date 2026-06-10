//ff:func feature=visibility type=generator control=sequence topic=report
//ff:what 리포트 → 기계용 JSON — 구조체 필드 순서 고정 직렬화, 같은 리포트면 바이트 동일
package report

import "encoding/json"

// JSON renders the machine-readable report. Field order is fixed by the
// struct declarations, so the output is byte-deterministic.
func JSON(r Report) []byte {
	data, _ := json.Marshal(r)
	return data
}
