//ff:type feature=visibility type=client topic=crawl
//ff:what ListObjectsV2 응답 XML의 디코드 대상 필드 — 키 목록, 절단 여부, 연속 토큰
package cflog

// s3ListResult is the subset of the ListObjectsV2 XML response the lister
// consumes.
type s3ListResult struct {
	IsTruncated           bool     `xml:"IsTruncated"`
	NextContinuationToken string   `xml:"NextContinuationToken"`
	Keys                  []string `xml:"Contents>Key"`
}
