//ff:type feature=quest type=schema topic=queue
//ff:what 근거 보강 퀘스트 Definition — reins gate.Definition(Seed/Render/Prepare/Rules) 구현체, 상태 없는 마커
package evidence

// Definition is the evidence quest's reins gate.Definition implementation.
// It is stateless: every method derives its inputs from the item payload and
// the instance on disk, so a disposable agent can resume from the session.
type Definition struct{}
