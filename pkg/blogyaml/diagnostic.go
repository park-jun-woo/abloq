//ff:type feature=blogyaml type=schema topic=diagnostics
//ff:what 검증 진단 1건 — 파일/라인/룰ID/메시지, 텍스트·JSON 출력의 공통 모델
package blogyaml

// Diagnostic is one validation finding, rendered as "file:line [rule] message".
type Diagnostic struct {
	File    string `json:"file"`
	Line    int    `json:"line"`
	Rule    string `json:"rule"`
	Message string `json:"message"`
}
