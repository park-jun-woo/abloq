//ff:func feature=sitesyaml type=parser control=sequence
//ff:what Parse가 strict 디코드(미지 키 진단)·active 기본값 주입·빈 파일·구문 오류 진단을 지키는지 검증
package sitesyaml

import "testing"

func TestParse(t *testing.T) {
	src := []byte("sites:\n  - name: a\n    repo_path: /blogs/a\n  - name: b\n    repo_path: /blogs/b\n    active: false\n")
	s, idx, diags := Parse("sites.yaml", src)
	if len(diags) != 0 {
		t.Fatalf("want 0 diagnostics, got %v", diags)
	}
	if len(s.Sites) != 2 {
		t.Fatalf("want 2 sites, got %d", len(s.Sites))
	}
	if !s.Sites[0].Active {
		t.Error("absent active key must default to true")
	}
	if s.Sites[1].Active {
		t.Error("explicit active: false must stay false")
	}
	if lineOf(idx, "sites[1].repo_path") != 5 {
		t.Errorf("line index missing: %d", lineOf(idx, "sites[1].repo_path"))
	}

	_, _, diags = Parse("sites.yaml", []byte("sites:\n  - nam: a\n"))
	if len(diags) == 0 || diags[0].Rule != "unknown-key" {
		t.Errorf("unknown key must be rejected (strict), got %v", diags)
	}

	_, _, diags = Parse("sites.yaml", []byte(""))
	if len(diags) != 1 || diags[0].Rule != "yaml-syntax" || diags[0].Message != "sites.yaml is empty" {
		t.Errorf("empty file diag = %v", diags)
	}

	_, _, diags = Parse("sites.yaml", []byte("sites: [\n"))
	if len(diags) == 0 || diags[0].Rule != "yaml-syntax" {
		t.Errorf("syntax error diag = %v", diags)
	}
}
