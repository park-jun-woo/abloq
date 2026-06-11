//ff:func feature=quest type=parser control=iteration dimension=1
//ff:what 디렉토리에서 위로 blog.yaml을 탐색해 인스턴스 루트(절대 경로) 반환 — 파일시스템 루트까지 없으면 에러
package writing

import (
	"fmt"
	"os"
	"path/filepath"
)

// findRoot walks up from dir to the nearest directory containing blog.yaml —
// the blog instance root every payload path is relative to.
func findRoot(dir string) (string, error) {
	for d := dir; ; d = filepath.Dir(d) {
		if _, err := os.Stat(filepath.Join(d, "blog.yaml")); err == nil {
			return d, nil
		}
		if d == filepath.Dir(d) {
			return "", fmt.Errorf("blog.yaml not found in any ancestor of %s", dir)
		}
	}
}
