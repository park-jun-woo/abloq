//ff:type feature=gate type=schema topic=evidence
//ff:what 인용 1건 — 마크다운 인라인 링크의 표기 텍스트, http(s) URL, 파일 라인 (Phase010 스캐너 공유 모델)
package gate

// Citation is one external citation in an article body: a markdown inline
// link to an http(s) URL. Label is the link text as written — citation-exists
// matches it against the cited page's title/og:title.
type Citation struct {
	Label string
	URL   string
	Line  int // 1-based file line number
}
