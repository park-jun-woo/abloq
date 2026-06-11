//ff:func feature=insight type=rule control=iteration dimension=1
//ff:what 인사이트 명세를 글 본문과 대조 — front matter 제외 본문을 NFC+케이스 폴딩 후 claim별 anchors 부분문자열 스크리닝, 섹션 실위치 포함
//ff:why 출현 = 대응 보장이 아니다 — 결과는 REVIEW 보조 자료로만 쓴다 (Phase015, 형태소 분석 등 비결정 요소 배제)
package insight

// Match screens every claim of ins against the article body (front matter
// excluded, NFC + case folding, substring). articlePath determines
// Result.Section from the file system location.
func Match(ins *Insight, articlePath string, article []byte) Result {
	foldedBody := fold(stripFrontMatter(article))
	res := Result{Section: sectionOf(articlePath)}
	for _, c := range ins.Claims {
		if claimFound(foldedBody, c) {
			res.Found = append(res.Found, c.ID)
			continue
		}
		res.Missing = append(res.Missing, c)
	}
	return res
}
