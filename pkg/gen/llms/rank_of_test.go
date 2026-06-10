//ff:func feature=gen type=generator control=sequence
//ff:what rankOf가 선언 순서 인덱스 맵을 만들고 미등록 값은 0랭크가 되는지 확인
package llms

import "testing"

func TestRankOf(t *testing.T) {
	rank := rankOf([]string{"ko", "en", "ja"})
	if rank["ko"] != 0 || rank["en"] != 1 || rank["ja"] != 2 {
		t.Errorf("rankOf = %v, want ko:0 en:1 ja:2", rank)
	}
	if rank["zz"] != 0 {
		t.Errorf("unlisted value rank = %d, want zero value 0", rank["zz"])
	}
}
