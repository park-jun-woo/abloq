//ff:func feature=cli type=command control=iteration dimension=1
//ff:what --variant 콤마 목록을 선언된 안으로 해석 — 각 이름을 상속 병합 완료 스펙으로, 미선언 이름은 에러
package main

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// ogVariantsFromList resolves a comma-separated --variant list against the
// blog.yaml declarations, preserving the requested order.
func ogVariantsFromList(og blogyaml.ImageOG, list string) ([]blogyaml.OGVariantSpec, error) {
	var specs []blogyaml.OGVariantSpec
	for _, name := range strings.Split(list, ",") {
		name = strings.TrimSpace(name)
		v, ok := og.Variant(name)
		if !ok {
			return nil, fmt.Errorf("--variant %q is not declared in blog.yaml image.og.variants", name)
		}
		specs = append(specs, v)
	}
	return specs, nil
}
