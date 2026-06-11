//ff:type feature=quest type=schema
//ff:what 번역 퀘스트 Definition — reins gate.Definition(Seed/Render/Prepare/Rules) 구현체, 상태 없는 마커
package translation

// Definition is the translation quest's reins gate.Definition implementation.
// It is stateless: every method derives its inputs from the item payload and
// the instance on disk, so a disposable agent can resume from the session.
type Definition struct{}
