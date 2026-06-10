//ff:func feature=gate type=rule control=iteration dimension=1
//ff:what [front-matter-schema] 픽스처 케이스 — 완전한 FM PASS 2건, 필드 누락·날짜 오류·역전·블록 없음 FAIL 검증
package gate

import "testing"

func TestRuleFrontMatterSchema(t *testing.T) {
	cases := []struct {
		name, file  string
		wantDiags   int
		wantMsgPart string
	}{
		{"pass canonical", "articles/pass.md", 0, ""},
		{"pass minimal", "articles/pass-minimal.md", 0, ""},
		{"fail missing lastmod and tags", "articles/schema-missing.md", 2, "lastmod must be"},
		{"fail bad date", "articles/schema-bad-date.md", 1, "date must be"},
		{"fail lastmod precedes date", "articles/schema-lastmod-before.md", 1, "must not precede"},
		{"fail no front matter", "articles/no-fm.md", 1, "missing or malformed"},
		{"fail broken yaml", "articles/schema-broken-yaml.md", 1, "not valid YAML"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkArticleRule(t, ruleFrontMatterSchema, tc.file, "front-matter-schema", tc.wantDiags, tc.wantMsgPart)
		})
	}
}
