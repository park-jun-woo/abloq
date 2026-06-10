//ff:func feature=queueio type=generator control=sequence
//ff:what Filename이 kind-lang-section-slug.yaml 정방향 파일명을 내는지 검증 (다국어 동일 slug 비충돌 포함)
package queueio

import "testing"

func TestFilename(t *testing.T) {
	ko := Item{Kind: "refresh", Lang: "ko", Section: "tech", Slug: "post-a"}
	en := Item{Kind: "refresh", Lang: "en", Section: "tech", Slug: "post-a"}
	if got := Filename(ko); got != "refresh-ko-tech-post-a.yaml" {
		t.Errorf("want refresh-ko-tech-post-a.yaml, got %s", got)
	}
	if Filename(ko) == Filename(en) {
		t.Error("translations of the same slug must not collide")
	}
}
