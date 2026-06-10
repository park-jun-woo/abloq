//ff:type feature=visibility type=client topic=crawl
//ff:what 실패 주입용 가짜 소스 — List 에러와 Get 에러를 골라 일으켜 수집기의 에러 경로를 검증
package cflog

// fakeSource fails List or Get on demand — the collector's error paths.
type fakeSource struct {
	keys    []string
	listErr error
	getErr  error
}
