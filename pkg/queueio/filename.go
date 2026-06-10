//ff:func feature=queueio type=generator control=sequence
//ff:what 행→파일명 정방향 계산 전용 — <kind>-<lang>-<section>-<slug>.yaml (중복 적재 키와 1:1, 역파싱 금지)
package queueio

// Filename derives the queue file name from an item. The mapping is
// forward-only: hyphens inside lang/section make it non-injective, so the
// name must never be parsed back into its parts — consumed sync recomputes
// it from the row and checks file existence instead.
func Filename(it Item) string {
	return it.Kind + "-" + it.Lang + "-" + it.Section + "-" + it.Slug + ".yaml"
}
