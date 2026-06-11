//ff:func feature=insight type=parser control=sequence
//ff:what strict 파싱 검증 — 정상 디코드, unknown key 에러, 빈 파일 진단
package insight

import "testing"

func TestParse(t *testing.T) {
	ins, diags := Parse("insight.yaml", []byte("topic: t\nsection: tech\nclaims:\n  - id: a\n    text: x\n    kind: claim\n    anchors: [\"k\"]\n"))
	if len(diags) != 0 {
		t.Fatalf("want no diagnostics for valid yaml, got %v", diags)
	}
	if ins.Topic != "t" || ins.Section != "tech" || len(ins.Claims) != 1 || ins.Claims[0].ID != "a" {
		t.Errorf("want decoded insight, got %+v", ins)
	}
	if ins.Claims[0].RequiresSource {
		t.Errorf("want requires_source default false, got true")
	}
	if _, diags := Parse("insight.yaml", []byte("topic: t\nbogus: 1\n")); len(diags) == 0 {
		t.Errorf("want unknown-key diagnostic for strict parse, got none")
	}
	_, diags = Parse("insight.yaml", nil)
	if len(diags) != 1 || diags[0].Rule != "yaml-syntax" {
		t.Errorf("want yaml-syntax diagnostic for empty file, got %v", diags)
	}
}
