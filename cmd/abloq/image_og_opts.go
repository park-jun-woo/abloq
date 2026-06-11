//ff:type feature=cli type=command
//ff:what image og 실행 옵션 묶음 — 인자(slug/title)·텍스트 옵션·provider 2축 플래그·플래그 명시 여부(해석 순서용)
//ff:why 해석 순서(플래그 > blog.yaml image.og > local)는 "플래그가 명시되었는가"를 알아야 한다 — cobra Changed를 *Set 필드로 운반 (BUG002)
package main

// imageOGOpts carries one `abloq image og` invocation. The *Set fields record
// whether the corresponding flag was explicitly given, so blog.yaml values
// only apply when the flag is absent.
type imageOGOpts struct {
	Slug, Title string
	// Summary is resolved internally from the article's front matter (not a
	// flag — the CLI surface stays (slug, title)); buildOGRuns injects it into
	// the {summary} prompt slot. Empty on the local/no-blog.yaml paths.
	Summary     string
	Brand       string
	FontPath    string
	OutDir      string
	Provider    string
	Model       string
	Overlay     bool
	VariantList string
	AllVariants bool
	Count       int
	ProviderSet bool
	ModelSet    bool
	OverlaySet  bool
}
