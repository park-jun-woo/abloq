//ff:func feature=sitesyaml type=schema control=sequence topic=diagnostics
//ff:what Diagnostic을 "파일:라인 [룰ID] 메시지" 한 줄로 포맷
package sitesyaml

import "fmt"

// String renders the diagnostic in the canonical "file:line [rule] message" form.
func (d Diagnostic) String() string {
	return fmt.Sprintf("%s:%d [%s] %s", d.File, d.Line, d.Rule, d.Message)
}
