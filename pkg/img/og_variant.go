//ff:type feature=image type=schema
//ff:what 해석 완료된 OG 안(variant) 1개 — 이름/모델(echo용)/오버레이 여부/치환 끝난 프롬프트/주입된 Provider 인스턴스
//ff:why provider명→인스턴스 해석은 cmd 계층 책임 — img는 (variant, Provider) 쌍만 받아 실행하므로 ogprovider import가 없고 네트워크 0 불변 (BUG002)
package img

// OGVariant is one fully resolved generation candidate: inheritance merging
// and prompt substitution are done upstream (blogyaml + cmd), the Provider
// instance is injected by the cmd layer. Model is informational (echo) — the
// instance is already configured with it.
type OGVariant struct {
	Name     string
	Model    string
	Overlay  bool
	Prompt   string
	Provider Provider
}
