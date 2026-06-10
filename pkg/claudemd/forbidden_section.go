//ff:func feature=claudemd type=generator control=sequence
//ff:what CLAUDE.md 금지 사항 섹션 — 치즈 방어 원칙: 게이트가 잡는 모든 우회를 시도 자체부터 금지
package claudemd

// forbiddenSection renders the cheese-defense rules. Every entry mirrors a
// deterministic gate — bypassing is detected, so don't attempt it.
func forbiddenSection() string {
	return `## 금지 사항 (치즈 방어)

게이트는 결과를 잡지만, 시도 자체가 금지다. 게이트를 "통과시키는" 게 아니라 일을 끝내라.

- 생성물(hugo.toml/robots.txt/llms.txt/jsonld.json)을 손으로 고치지 않는다 — blog.yaml을 고치고 ` + "`abloq generate`" + `.
- 본문 실변경 없이 lastmod만 갱신하지 않는다 — honest-lastmod가 토큰 diff로 차단한다.
- 출처를 날조하지 않는다 — citation-exists가 URL 실재(HTTP 200 + 제목 일치)를 검증한다.
- 수치 주장은 같은 문단에 출처 링크 필수 — 출처를 못 찾으면 주장을 빼라, 지어내지 마라.
- 번역·갱신에서 본문 라인을 누락·재배열하지 않는다 — body-lossless·section-preserved가 git HEAD와 비교한다.
- ` + "`quests/queue/`" + `에 없는 글의 주장·구조를 임의로 바꾸지 않는다.
- 게이트 실패 상태로 커밋·배포하지 않는다.
- 게이트를 통과시키려고 blog.yaml 임계값을 완화하지 않는다 — 임계값 변경은 사람의 결정이다.
`
}
