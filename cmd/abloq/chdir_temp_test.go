//ff:func feature=cli type=command control=sequence
//ff:what 테스트 헬퍼 — 작업 디렉토리를 dir로 옮기고 테스트 종료 시 원복 (image og의 cwd 상대 경로 검증용)
package main

import (
	"os"
	"testing"
)

func chdirTemp(t *testing.T, dir string) {
	t.Helper()
	old, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir %s: %v", dir, err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(old); err != nil {
			t.Fatalf("chdir back: %v", err)
		}
	})
}
