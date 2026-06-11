//ff:type feature=quest type=schema
//ff:what Seed 중간 산물 — 원문 1편의 인스턴스 루트, 루트 기준 경로, Key 부품, 선언 언어 목록, 원문 lastmod(파싱 성공 여부 포함)
package translation

import (
	"time"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// seedSrc is what seedOrigin derives from one origin article argument: the
// instance root, the root-relative origin path and its key parts, the loaded
// blog.yaml (language list first = default) and the origin's lastmod when it
// parses (hasLastmod=false ⇒ every translation is treated as stale).
type seedSrc struct {
	root       string
	origin     string
	originLang string
	section    string
	slug       string
	blog       *blogyaml.Blog
	lastmod    time.Time
	hasLastmod bool
}
