//ff:func feature=scan type=generator control=sequence topic=evidence
//ff:what evidenceItem이 kind=evidence 좌표·검출 건수 priority·전 언어 keys·JSON 문자열 payload(빈 쪽 키 생략)를 채우는지 검증
package evidence

import "testing"

func TestEvidenceItem(t *testing.T) {
	b := testBlog(t)
	a := testArticle(t, b, "---\ntitle: T\n---\n\nBody.\n")
	claims := []ClaimRef{{Hash: "abcd", Loc: "content/ko/tech/fixture.md:5", Text: "x 40% 증가"}}
	it := evidenceItem(a, claims, []string{"https://gone.example/x"}, []string{"ko", "en"})
	if it.Kind != "evidence" || it.Lang != "ko" || it.Section != "tech" || it.Slug != "fixture" {
		t.Errorf("item coordinates: %+v", it)
	}
	if it.Priority != 2 {
		t.Errorf("priority = %d, want finding count 2", it.Priority)
	}
	if it.Payload["claims"] != `[{"hash":"abcd","loc":"content/ko/tech/fixture.md:5","text":"x 40% 증가"}]` {
		t.Errorf("claims payload = %s", it.Payload["claims"])
	}
	if it.Payload["rot_urls"] != `["https://gone.example/x"]` {
		t.Errorf("rot payload = %s", it.Payload["rot_urls"])
	}
	if len(it.Keys) != 2 || it.Keys[0] != "ko/tech/fixture" || it.Keys[1] != "en/tech/fixture" {
		t.Errorf("keys must cover every declared language: %v", it.Keys)
	}
	onlyClaims := evidenceItem(a, claims, nil, []string{"ko"})
	if _, ok := onlyClaims.Payload["rot_urls"]; ok {
		t.Error("empty rot list must omit the key")
	}
	onlyRot := evidenceItem(a, nil, []string{"https://gone.example/x"}, []string{"ko"})
	if _, ok := onlyRot.Payload["claims"]; ok {
		t.Error("empty claims list must omit the key")
	}
}
