//ff:func feature=bots type=dict control=sequence topic=crawl
//ff:what 미지 봇 후보 휴리스틱 토큰 사전 — bot/crawler/spider 류 + 비브라우저 클라이언트 패턴 (소문자, 부분일치용)
//ff:why analyze-stats.py BOTPAT의 일반 토큰부를 승계 — 일반 브라우저 UA에는 안 나타나는 토큰만 담아 unknown_bots 적재가 사전 갱신 입력으로 쓸 만하게 유지한다 (Phase012)
package bots

// candidateTokens returns the lowercase substrings that mark a User-Agent as
// a bot candidate: generic crawler tokens plus non-browser client patterns.
func candidateTokens() []string {
	return []string{
		"bot", "crawl", "spider", "slurp", "scan", "preview",
		"archive.org", "petal", "ahrefs", "semrush", "yandex",
		"facebookexternalhit", "embedly", "dataforseo", "barkrowler",
		"expanse", "censys", "masscan", "zgrab",
		"python-requests", "curl", "wget", "go-http", "http-client",
		"okhttp", "java/", "libwww", "node-fetch",
		"headless", "monitor", "uptime",
	}
}
