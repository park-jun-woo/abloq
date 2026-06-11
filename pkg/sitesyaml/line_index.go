//ff:type feature=sitesyaml type=parser
//ff:what YAML 키 경로("sites[0].repo_path") → 소스 라인 번호 맵, 진단 라인 표기의 근거
package sitesyaml

// lineIndex maps a dotted key path in sites.yaml to its 1-based source line.
type lineIndex map[string]int
