//ff:func feature=gate type=frame control=iteration dimension=1 topic=baseline
//ff:what 작업 트리에서 변경된 글에만 git HEAD 원본을 파싱해 부착 — 미변경 글은 현재본을 원본으로 공유
package gate

// attachBaselines sets each article's Base. Unchanged and untracked files
// share Doc as their own baseline (comparison rules pass trivially); changed
// tracked files get the parsed HEAD snapshot; files added since HEAD and
// non-git dirs keep Base nil (skip — a new article has no original).
func attachBaselines(dir string, hi headingIndex, arts []*Article) {
	changed, ok := gitChangedSet(dir)
	if !ok {
		return
	}
	for _, a := range arts {
		if !changed[a.Path] {
			a.Base = a.Doc
			continue
		}
		raw, found := gitShow(dir, a.Path)
		if found {
			a.Base = parseDoc(hi, a.Lang, string(raw))
		}
	}
}
