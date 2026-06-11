//ff:func feature=quest type=frame control=sequence
//ff:what hugo 바이너리 경로 조회 — PATH 탐색(테스트 주입 가능), Prepare 사전 점검과 hugo-build 룰이 공유
//ff:why hugo 부재는 SKIP이 아니라 Prepare 에러(제출 중단·try 미소진) — SKIP 진단은 reins 계약상 침묵 통과 = PATH 조작 치즈 구멍 (Phase017 계획)
package translation

import "os/exec"

// hugoLook is the PATH lookup, swappable in tests.
var hugoLook = exec.LookPath

// hugoPath resolves the hugo binary both Prepare (presence pre-check, no try
// burned) and the hugo-build rule use.
func hugoPath() (string, error) {
	return hugoLook("hugo")
}
