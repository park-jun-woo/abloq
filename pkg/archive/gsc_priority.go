//ff:func feature=archive type=client control=sequence
//ff:what GSC 쿼터 분할 우선순위 — 0=신규 글(date==lastmod), 1=갱신 글, 2=posts 인덱스에 없는 target
package archive

// gscPriority ranks one pending receipt for the quota split. New posts win:
// first-time indexing is where the Indexing API moves the needle most.
func gscPriority(p Pending) int {
	if p.Date == "" {
		return 2
	}
	if p.Date == p.Lastmod {
		return 0
	}
	return 1
}
