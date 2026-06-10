//ff:func feature=claudemd type=generator control=sequence
//ff:what CLAUDE.md 명령어 섹션 — abloq CLI(게이트 실행법 포함)와 hugo 명령 레퍼런스
package claudemd

// commandsSection renders the command reference (fixed text — the CLI is the
// same for every abloq blog).
func commandsSection() string {
	return `## 명령어

` + "```bash" + `
abloq validate .             # blog.yaml 스키마 + 검증 룰 (exit 1 = 진단 있음)
abloq generate .             # 파생물 4종 재생성 (멱등)
abloq check .                # 파생물 드리프트 검사 — blog.yaml과 바이트 일치 확인
abloq gate .                 # 게이트 전체 룰 실행 (exit 1 = 위반)
abloq gate --offline .       # 네트워크 룰(citation-exists) 제외
abloq gate --rule <id> .     # 룰 1개만 (예: --rule honest-lastmod)
abloq gate --json .          # 진단을 JSON으로
abloq postbuild md .         # public/ 옆에 글별 .md 산출
abloq image convert <src>    # WebP 변환 (흰 배경 평탄화, 가로 1400px 제한)
abloq image og <slug> <제목> # 1200×630 OG 이미지 생성
abloq claudemd .             # 이 파일 재생성 (blog.yaml 변경 후)
hugo                         # 사이트 빌드 → public/ (게이트 전 빌드는 --minify 없이)
` + "```" + `

`
}
