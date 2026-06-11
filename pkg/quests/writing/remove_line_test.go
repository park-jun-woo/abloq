//ff:func feature=quest type=parser control=iteration dimension=1
//ff:what 테스트 헬퍼 — 부분 문자열을 포함한 라인을 제거한 사본 반환 (픽스처 변형용)
package writing

import "strings"

func removeLine(text, substr string) string {
	var out []string
	for _, ln := range strings.Split(text, "\n") {
		if strings.Contains(ln, substr) {
			continue
		}
		out = append(out, ln)
	}
	return strings.Join(out, "\n")
}
