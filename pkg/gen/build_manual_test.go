//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what llms_txt mode manual/off에서 Build 출력 목록에 static/llms.txt가 빠지는지 검증 — generate·check 동시 옵트아웃의 근거
package gen

import "testing"

func TestBuildManualOff(t *testing.T) {
	for _, mode := range []string{"manual", "off"} {
		checkBuildExcludesLlms(t, mode)
	}
}
