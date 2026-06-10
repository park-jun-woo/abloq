//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what [crawlers-policy] geo.crawlers의 각 정책 값이 allow|block enum인지 검증
package blogyaml

import "fmt"

// ruleCrawlersPolicy requires every crawler policy value to be "allow" or "block".
func ruleCrawlersPolicy(filename string, b *Blog, idx lineIndex) []Diagnostic {
	var diags []Diagnostic
	for _, name := range sortedKeys(b.Geo.Crawlers) {
		policy := b.Geo.Crawlers[name]
		if policy != "allow" && policy != "block" {
			diags = append(diags, Diagnostic{
				File: filename, Line: lineOf(idx, "geo.crawlers."+name), Rule: "crawlers-policy",
				Message: fmt.Sprintf("geo.crawlers.%s %q must be one of: allow, block", name, policy),
			})
		}
	}
	return diags
}
