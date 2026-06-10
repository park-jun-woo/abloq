//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what env에 API 키가 있는 엔진만 조립 — perplexity/openai/anthropic 고정 순서, 베이스 URL·모델은 env 오버라이드
//ff:why API 키는 env로만(설계 §3.3 — 백엔드 전용, 에이전트·영수증에 없다). 키가 없는 엔진은 조용히 빠진다 — 옵트인 (Phase013)
package citation

// EnginesFromEnv assembles the engines whose API key env is set, in the
// fixed order perplexity, openai, anthropic (sample rows stay comparable
// across runs). Base URLs are env-overridable so the Hurl stub can
// intercept every engine on one port.
func EnginesFromEnv() []Engine {
	var engines []Engine
	if key := envOr("PERPLEXITY_API_KEY", ""); key != "" {
		base := envOr("PERPLEXITY_BASE_URL", "https://api.perplexity.ai")
		model := envOr("PERPLEXITY_MODEL", "sonar")
		engines = append(engines, Engine{Name: "perplexity", Ask: func(q string) ([]string, error) {
			return askPerplexity(base, key, model, q)
		}})
	}
	if key := envOr("OPENAI_API_KEY", ""); key != "" {
		base := envOr("OPENAI_BASE_URL", "https://api.openai.com")
		model := envOr("OPENAI_MODEL", "gpt-4.1")
		engines = append(engines, Engine{Name: "openai", Ask: func(q string) ([]string, error) {
			return askOpenAI(base, key, model, q)
		}})
	}
	if key := envOr("ANTHROPIC_API_KEY", ""); key != "" {
		base := envOr("ANTHROPIC_BASE_URL", "https://api.anthropic.com")
		model := envOr("ANTHROPIC_MODEL", "claude-sonnet-4-5")
		engines = append(engines, Engine{Name: "anthropic", Ask: func(q string) ([]string, error) {
			return askAnthropic(base, key, model, q)
		}})
	}
	return engines
}
