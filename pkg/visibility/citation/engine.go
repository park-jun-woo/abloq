//ff:type feature=visibility type=schema topic=citation
//ff:what 샘플링 엔진 1종 — 이름과 Ask(질의→인용 URL 목록) 함수 필드, 구현은 ask_* 단발 HTTPS JSON POST
//ff:why 엔진을 인터페이스 뒤에 둔다(설계 §6.3): 러너는 이름과 질의→인용만 안다 — API 형태 차이는 ask_* 함수에 격리, 신규 의존성 0 (Phase013)
package citation

// Engine is one sampling engine: a stable name (the citation_samples.engine
// value) and an Ask function returning the structured citation URLs of the
// engine's answer. Keys and base URLs are bound by EnginesFromEnv.
type Engine struct {
	Name string
	Ask  func(query string) ([]string, error)
}
