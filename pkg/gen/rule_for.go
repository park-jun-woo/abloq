//ff:func feature=gen type=rule control=selection topic=drift
//ff:what 파생물 경로를 게이트 룰ID로 매핑 — robots-policy-match/llmstxt-sync/hugo-config-sync/jsonld-sync
package gen

// ruleFor names the gate rule violated when a derived file drifts.
func ruleFor(path string) string {
	switch path {
	case "hugo.toml":
		return "hugo-config-sync"
	case "static/robots.txt":
		return "robots-policy-match"
	case "static/llms.txt":
		return "llmstxt-sync"
	case "data/jsonld.json":
		return "jsonld-sync"
	default:
		return "derived-sync"
	}
}
