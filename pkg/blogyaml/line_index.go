//ff:type feature=blogyaml type=parser
//ff:what YAML 키 경로("geo.freshness_days", "languages[2]") → 소스 라인 번호 맵, 진단 라인 표기의 근거
package blogyaml

// lineIndex maps a dotted key path in blog.yaml to its 1-based source line.
type lineIndex map[string]int
