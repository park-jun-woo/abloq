//ff:func feature=claudemd type=generator control=sequence
//ff:what CLAUDE.md 디렉토리 규약 섹션 — SSOT/콘텐츠/레이아웃/생성물/큐/IaC 경로와 취급 규칙, llms.txt는 mode auto일 때만 생성물 목록에 포함
package claudemd

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// dirsSection renders the directory conventions. static/llms.txt is listed
// as a generated file only in llms_txt mode auto — in manual/off the file is
// hand-curated (or absent) and abloq keeps its hands off it.
func dirsSection(b *blogyaml.Blog) string {
	derived := "`hugo.toml` · `static/robots.txt` · `data/jsonld.json`"
	if b.Geo.LlmsTxtMode() == "auto" {
		derived = "`hugo.toml` · `static/robots.txt` · `static/llms.txt` · `data/jsonld.json`"
	}
	return `## 디렉토리 규약

- ` + "`blog.yaml`" + ` — SSOT. 파생물·게이트 파라미터 전부의 원천. 구조·임계값 변경은 여기서만.
- ` + "`content/{lang}/{section}/{slug}.md`" + ` — 글. 전 언어 동일 slug(영문 소문자-하이픈 파일명).
- ` + "`layouts/ assets/ static/`" + ` — 레이아웃 팩. 수정 자유, 게이트와 무관. 소셜 링크·홈 배너는 ` + "`layouts/partials/hooks/`" + `.
- ` + derived + ` — **생성물. 직접 수정 금지.** ` + "`abloq generate`" + `가 만들고 ` + "`abloq check`" + `가 드리프트를 잡는다.
- ` + "`config/_default/`" + ` — 인스턴스 전용 Hugo 설정 오버레이(선택, 사람 소유). Hugo가 생성된 hugo.toml 위에 병합한다. blog.yaml이 이미 내는 키(baseURL·title·defaultContentLanguage*·[sitemap]·languages.*.weight/contentDir)는 재선언 금지.
- ` + "`public/`" + ` — hugo 빌드 출력. 커밋하지 않는다.
- ` + "`quests/queue/`" + ` — 운용 백엔드가 떨어뜨린 작업 큐(yaml). 갱신·근거 보강·클러스터 퀘스트의 유일한 입력.
- ` + "`deploy/terraform/`" + ` — 배포 IaC(옵션). 적용은 사람이 결정한다.

`
}
