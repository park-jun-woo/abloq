//ff:func feature=init type=parser control=iteration dimension=1
//ff:what fs.FS 디렉토리를 재귀 순회해 파일 경로 목록을 결정적 순서(ReadDir 정렬)로 수집
package scaffold

import (
	"io/fs"
	"path"
)

// ListDir returns every file path under dir (recursive), in lexical order.
func ListDir(fsys fs.FS, dir string) ([]string, error) {
	entries, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, err
	}
	var paths []string
	for _, e := range entries {
		p := path.Join(dir, e.Name())
		if !e.IsDir() {
			paths = append(paths, p)
			continue
		}
		sub, err := ListDir(fsys, p)
		if err != nil {
			return nil, err
		}
		paths = append(paths, sub...)
	}
	return paths, nil
}
