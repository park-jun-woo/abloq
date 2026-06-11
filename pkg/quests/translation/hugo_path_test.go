//ff:func feature=quest type=frame control=sequence
//ff:what hugoPath 검증 — 주입된 lookup으로 경로/에러 전달 확인
package translation

import (
	"fmt"
	"testing"
)

func TestHugoPath(t *testing.T) {
	orig := hugoLook
	defer func() { hugoLook = orig }()
	hugoLook = func(name string) (string, error) { return "/fake/" + name, nil }
	if p, err := hugoPath(); err != nil || p != "/fake/hugo" {
		t.Errorf("path = %q err=%v", p, err)
	}
	hugoLook = func(string) (string, error) { return "", fmt.Errorf("absent") }
	if _, err := hugoPath(); err == nil {
		t.Error("absent: want error")
	}
}
