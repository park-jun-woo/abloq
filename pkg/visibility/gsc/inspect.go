//ff:func feature=visibility type=client control=iteration dimension=1 topic=gsc
//ff:what URL Inspection 선별 조회 — 대상 URL들을 1건씩 inspectOne으로 조회해 요약 목록으로 합산 (옵트인 전용)
//ff:why 쿼터가 작다(일 2000·분 600) — 호출자는 lastmod 최근 N일 글로 좁히고 1회전 상한을 둔다. 읽는 이 없는 출력에 쿼터를 태우지 않도록 기본 off (Phase013)
package gsc

// Inspect runs the URL Inspection API over the given URLs and returns one
// verdict summary per URL. Callers cap and filter the list — this function
// burns quota exactly len(urls) times.
func Inspect(base, token, site string, urls []string) ([]Inspection, error) {
	endpoint := base + "/v1/urlInspection/index:inspect"
	var out []Inspection
	for _, u := range urls {
		ins, err := inspectOne(endpoint, token, site, u)
		if err != nil {
			return nil, err
		}
		out = append(out, ins)
	}
	return out, nil
}
