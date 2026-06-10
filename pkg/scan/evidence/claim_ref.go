//ff:type feature=scan type=schema topic=evidence
//ff:what 무출처 수치 주장 1건의 큐 payload 표현 — 텍스트 해시(1차 키)·위치(파일:라인, 힌트)·주장 원문
//ff:why 라인 번호는 후속 커밋에 흔들린다 — 에이전트·게이트("큐에 없는 주장 변경 금지")는 해시로 주장을 특정하고 라인은 힌트로만 쓴다
package evidence

// ClaimRef is one unsourced numeric claim as it rides in the queue payload.
// Hash is the stable primary key (sha256 prefix of the claim text); Loc is a
// repository-relative path:line hint that may drift; Text is the claim line
// verbatim so the agent can locate it without re-running detection.
type ClaimRef struct {
	Hash string `json:"hash"`
	Loc  string `json:"loc"`
	Text string `json:"text"`
}
