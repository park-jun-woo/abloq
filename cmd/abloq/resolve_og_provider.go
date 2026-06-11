//ff:func feature=cli type=command control=selection
//ff:what provider명→인스턴스 해석 — gemini는 env 키로 구성(부재 시 명확한 에러), 그 외는 미지원 진단. 실효 모델명을 echo용으로 동봉
//ff:why 인스턴스 주입은 cmd 계층 책임 — pkg/img는 (variant, Provider) 쌍만 받아 네트워크 0 불변을 지킨다 (BUG002)
package main

import (
	"fmt"

	"github.com/park-jun-woo/abloq/pkg/img"
	"github.com/park-jun-woo/abloq/pkg/img/ogprovider"
)

// resolveOGProvider builds the provider instance for one variant and returns
// it with the effective model id (for plan/outcome echo). "local" never comes
// through here — it bypasses Provider entirely.
func resolveOGProvider(name, model string) (img.Provider, string, error) {
	switch name {
	case "gemini":
		g, err := ogprovider.NewGemini(model)
		if err != nil {
			return nil, "", err
		}
		return g, g.Model, nil
	}
	return nil, "", fmt.Errorf("unknown OG provider %q (supported: local, gemini)", name)
}
