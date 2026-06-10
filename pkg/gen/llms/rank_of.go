//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what 목록의 각 값에 선언 순서 인덱스를 매기는 랭크 맵 생성 — 언어·섹션 정렬 키의 입력
package llms

// rankOf maps each value to its declaration index for order-preserving sorts.
func rankOf(values []string) map[string]int {
	rank := make(map[string]int, len(values))
	for i, v := range values {
		rank[v] = i
	}
	return rank
}
