//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what Key 부품(lang/section/slug)으로 실재하는 글 경로 해석 — 플랫(.md) 우선, 번들(slug/index.md) 차선, 둘 다 없으면 에러 (퀘스트 공용)
package common

import (
	"fmt"
	"os"
	"path/filepath"
)

// ArticlePath resolves the root-relative path of an existing article from
// its key parts. Queue items point at committed articles, so a key whose
// file exists in neither the flat nor the bundle form is an error.
func ArticlePath(root, lang, section, slug string) (string, error) {
	flat := "content/" + lang + "/" + section + "/" + slug + ".md"
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(flat))); err == nil {
		return flat, nil
	}
	bundle := "content/" + lang + "/" + section + "/" + slug + "/index.md"
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(bundle))); err == nil {
		return bundle, nil
	}
	return "", fmt.Errorf("article for key %s/%s/%s not found (looked at %s and %s)", lang, section, slug, flat, bundle)
}
