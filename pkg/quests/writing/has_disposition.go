//ff:func feature=quest type=rule control=iteration dimension=1
//ff:what REVIEW 기록에 claim ID의 disposition 라인("- <id>: addressed|revised|excluded ...")이 있는지 검사
package writing

import "strings"

// hasDisposition reports whether the review record contains a disposition
// line for the claim ID: `- <id>: <disposition> ...` where the disposition
// token is addressed, revised or excluded.
func hasDisposition(review, id string) bool {
	for _, line := range strings.Split(review, "\n") {
		rest, ok := strings.CutPrefix(strings.TrimSpace(line), "- "+id+":")
		if !ok {
			continue
		}
		fields := strings.Fields(rest)
		if len(fields) == 0 {
			continue
		}
		if fields[0] == "addressed" || fields[0] == "revised" || fields[0] == "excluded" {
			return true
		}
	}
	return false
}
