//ff:func feature=gate type=rule control=sequence
//ff:what structure.order에 특수 토큰(image/attribution 등)이 선언됐는지 판정 — 미선언 룰은 스킵
package gate

import (
	"slices"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// orderHas reports whether structure.order declares the given entry.
func orderHas(b *blogyaml.Blog, key string) bool {
	return slices.Contains(b.Structure.Order, key)
}
