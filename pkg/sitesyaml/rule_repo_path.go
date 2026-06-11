//ff:func feature=sitesyaml type=rule control=iteration dimension=1
//ff:what [repo-path] repo_path 필수 + 절대경로 검증 — 마운트 컨벤션과 무관하게 코드는 절대경로만 신뢰한다
package sitesyaml

import (
	"fmt"
	"path/filepath"
)

// ruleRepoPath requires every repo_path and keeps it absolute: the backend
// trusts only the declared absolute path, never a mount convention.
func ruleRepoPath(filename string, s *Sites, idx lineIndex) []Diagnostic {
	var diags []Diagnostic
	for i, site := range s.Sites {
		if site.RepoPath == "" {
			diags = append(diags, Diagnostic{
				File: filename, Line: lineOfSite(idx, i, "repo_path"), Rule: "repo-path",
				Message: fmt.Sprintf("sites[%d].repo_path is required", i),
			})
			continue
		}
		if !filepath.IsAbs(site.RepoPath) {
			diags = append(diags, Diagnostic{
				File: filename, Line: lineOfSite(idx, i, "repo_path"), Rule: "repo-path",
				Message: fmt.Sprintf("sites[%d].repo_path %q must be an absolute path", i, site.RepoPath),
			})
		}
	}
	return diags
}
