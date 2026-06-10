//ff:func feature=archive type=client control=iteration dimension=1
//ff:what pending 영수증들의 target URL 목록 추출
package archive

// targetList collects the target URLs of the pending receipts in order.
func targetList(pending []Pending) []string {
	urls := make([]string, 0, len(pending))
	for _, p := range pending {
		urls = append(urls, p.Target)
	}
	return urls
}
