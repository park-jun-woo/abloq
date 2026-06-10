//ff:func feature=queueio type=client control=sequence
//ff:what PublishFile이 도달 불가 origin의 클론 실패를 에러로 내는지 검증
package queueio

import "testing"

func TestPublishFileCloneError(t *testing.T) {
	cfg := Config{RepoURL: "file:///nonexistent/origin.git", Workdir: t.TempDir() + "/w",
		AuthorName: "a", AuthorEmail: "a@t"}
	if _, err := PublishFile(cfg, "reports/x.md", []byte("x")); err == nil {
		t.Error("an unreachable origin must error")
	}
}
