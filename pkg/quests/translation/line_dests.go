//ff:func feature=quest type=parser control=iteration dimension=1 topic=lossless
//ff:what 라인 1개의 인라인 링크 목적지 추출 — `[..](dest)` 매치 중 이미지(`![`) 선행 매치는 제외
package translation

import "regexp"

var reLinkDest = regexp.MustCompile(`\[[^\]]*\]\(([^)\s]+)[^)]*\)`)

// lineDests extracts the non-image inline link destinations of one prose
// line (a match preceded by '!' is an image, owned by the image check).
func lineDests(raw string) []string {
	var dests []string
	for _, m := range reLinkDest.FindAllStringSubmatchIndex(raw, -1) {
		if m[0] > 0 && raw[m[0]-1] == '!' {
			continue
		}
		dests = append(dests, raw[m[2]:m[3]])
	}
	return dests
}
