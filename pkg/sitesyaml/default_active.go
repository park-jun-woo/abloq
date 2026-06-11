//ff:func feature=sitesyaml type=parser control=iteration dimension=1
//ff:what active 키가 파일에 없는 사이트에 기본값 true를 주입 — 라인 인덱스로 키 부재를 판별
//ff:why 리스트 항목은 blogyaml식 선주입 디코드가 불가능 — 포인터 필드 없이 "미지정 = true, 명시 false = false"를 지키려고 노드 인덱스로 부재를 본다
package sitesyaml

import "fmt"

// defaultActive sets Active to true for every site whose entry has no
// explicit active key. An explicit "active: false" stays false.
func defaultActive(s *Sites, idx lineIndex) {
	for i := range s.Sites {
		if _, ok := idx[fmt.Sprintf("sites[%d].active", i)]; !ok {
			s.Sites[i].Active = true
		}
	}
}
