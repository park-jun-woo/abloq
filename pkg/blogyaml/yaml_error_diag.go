//ff:func feature=blogyaml type=parser control=sequence topic=diagnostics
//ff:what yaml 에러 메시지 1건에서 라인 번호를 뽑고 unknown-key/yaml-syntax 룰ID로 분류
package blogyaml

import (
	"regexp"
	"strconv"
	"strings"
)

var yamlErrLineRe = regexp.MustCompile(`line (\d+):`)

// yamlErrorDiag classifies one yaml.v3 error message into a Diagnostic.
func yamlErrorDiag(filename, msg string) Diagnostic {
	line := 1
	if m := yamlErrLineRe.FindStringSubmatch(msg); m != nil {
		line, _ = strconv.Atoi(m[1])
	}
	rule := "yaml-syntax"
	if strings.Contains(msg, "not found in type") {
		rule = "unknown-key"
	}
	msg = strings.TrimPrefix(msg, "yaml: ")
	return Diagnostic{File: filename, Line: line, Rule: rule, Message: msg}
}
