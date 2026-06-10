//ff:func feature=gate type=parser control=sequence topic=baseline
//ff:what git show HEAD:./경로로 글의 HEAD 스냅샷을 읽음 — HEAD에 없으면(신규 글) false
package gate

import "os/exec"

// gitShow reads the HEAD snapshot of a dir-relative path.
func gitShow(dir, rel string) ([]byte, bool) {
	out, err := exec.Command("git", "-C", dir, "show", "HEAD:./"+rel).Output()
	if err != nil {
		return nil, false
	}
	return out, true
}
