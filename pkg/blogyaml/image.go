//ff:type feature=blogyaml type=schema
//ff:what blog.yaml image 섹션 — OG 이미지 생성 선언(og)만, 미선언 시 전부 기존 local 동작 (완전 하위호환)
package blogyaml

// Image declares the site's image tooling policy. Only the OG block exists
// today; an absent image key keeps every zero value, which means the local
// deterministic card — full backward compatibility (BUG002).
type Image struct {
	OG ImageOG `yaml:"og" json:"og"`
}
