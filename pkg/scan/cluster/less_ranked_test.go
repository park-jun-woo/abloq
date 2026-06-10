//ff:func feature=scan type=rule control=iteration dimension=1 topic=cluster
//ff:what lessRanked가 교집합→동일 섹션→발행일 근접→키 사전순으로 비교하는지 검증
package cluster

import "testing"

func TestLessRanked(t *testing.T) {
	cases := []struct {
		name string
		a, b ranked
		want bool
	}{
		{"more shared tags first", ranked{cand: Candidate{SharedTags: 2}}, ranked{cand: Candidate{SharedTags: 1}}, true},
		{"same section first", ranked{sameSection: true}, ranked{sameSection: false}, true},
		{"closer date first", ranked{dateDist: 1}, ranked{dateDist: 2}, true},
		{"key lexicographic", ranked{key: "tech/a"}, ranked{key: "tech/b"}, true},
		{"key reverse", ranked{key: "tech/b"}, ranked{key: "tech/a"}, false},
	}
	for _, tc := range cases {
		if got := lessRanked(tc.a, tc.b); got != tc.want {
			t.Errorf("%s: lessRanked = %v, want %v", tc.name, got, tc.want)
		}
	}
}
