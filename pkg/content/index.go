//ff:func feature=content type=parser control=iteration dimension=1
//ff:what 블로그 저장소 전수 인덱싱 — blog.yaml 로드 후 선언 언어·섹션의 발행 글 front matter를 인덱스 항목으로 수집
//ff:why blog.yaml이 단일 진실: 선언되지 않은 디렉토리는 인덱스에 없다 — abloqd posts 테이블과 CLI가 같은 파서를 공유해 두 모드의 인덱스가 갈라질 수 없다 (설계서 §3.3)
package content

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// IndexRepo parses every published article's front matter under
// root/content/ for the languages and sections declared in root/blog.yaml.
// Order is deterministic: declared language order, declared section order,
// directory-name order.
func IndexRepo(root string) ([]Entry, error) {
	b, diags, err := blogyaml.Load(filepath.Join(root, "blog.yaml"))
	if err != nil {
		return nil, err
	}
	if len(diags) > 0 {
		return nil, fmt.Errorf("blog.yaml invalid: %s", diags[0].String())
	}
	entries := []Entry{}
	for _, lang := range b.Languages {
		entries = append(entries, indexLang(root, b, lang)...)
	}
	return entries, nil
}
