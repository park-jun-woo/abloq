//ff:func feature=queueio type=generator control=sequence
//ff:what rowItems가 DB id를 제거한 Item 목록을 반환하는지 검증
package queueio

import "testing"

func TestRowItems(t *testing.T) {
	rows := []Row{{ID: 1, Item: Item{Slug: "a"}}, {ID: 2, Item: Item{Slug: "b"}}}
	items := rowItems(rows)
	if len(items) != 2 || items[0].Slug != "a" || items[1].Slug != "b" {
		t.Errorf("unexpected items: %+v", items)
	}
}
