//ff:func feature=blogyaml type=parser control=sequence
//ff:what image.og 블록 파싱 검증 — provider/model/overlay/prompt 디코드, variant optional 포인터(미지정 nil vs 명시 false/"") 구분, 미지 키 거부, 미선언 하위호환
package blogyaml

import (
	"strings"
	"testing"
)

func TestParseImageOG(t *testing.T) {
	src := []byte(`languages: [en]
sections: [tech]
image:
  og:
    provider: gemini
    model: base-model
    overlay: true
    prompt: |
      Site prompt for "{title}".
    variants:
      - name: minimal
        prompt: Flat shapes.
      - name: photo
        model: ""
        overlay: false
`)
	b, _, diags := Parse("blog.yaml", src)
	if len(diags) != 0 {
		t.Fatalf("want 0 diagnostics, got %v", diags)
	}
	og := b.Image.OG
	if og.Provider != "gemini" || og.Model != "base-model" || !og.Overlay {
		t.Errorf("og = %+v, want gemini/base-model/overlay", og)
	}
	if len(og.Variants) != 2 {
		t.Fatalf("variants = %d, want 2", len(og.Variants))
	}
	// minimal: prompt set, model/overlay absent (nil = inherit)
	if og.Variants[0].Prompt == nil || *og.Variants[0].Prompt != "Flat shapes." {
		t.Errorf("minimal prompt = %v", og.Variants[0].Prompt)
	}
	if og.Variants[0].Model != nil || og.Variants[0].Overlay != nil {
		t.Errorf("minimal model/overlay must be nil (inherit), got %+v", og.Variants[0])
	}
	// photo: explicit "" and false are real overrides, not absence
	if og.Variants[1].Model == nil || *og.Variants[1].Model != "" {
		t.Errorf("photo model = %v, want explicit empty string", og.Variants[1].Model)
	}
	if og.Variants[1].Overlay == nil || *og.Variants[1].Overlay {
		t.Errorf("photo overlay = %v, want explicit false", og.Variants[1].Overlay)
	}

	// unknown key inside image.og is rejected (strict parsing)
	bad := []byte("languages: [en]\nsections: [tech]\nimage:\n  og:\n    providr: gemini\n")
	_, _, diags = Parse("blog.yaml", bad)
	if len(diags) == 0 || !strings.Contains(diags[0].Message, "providr") {
		t.Errorf("unknown key: want unknown-key diagnostic, got %v", diags)
	}

	// absent image block keeps full backward compatibility: effective local
	plain := []byte("languages: [en]\nsections: [tech]\n")
	b, _, diags = Parse("blog.yaml", plain)
	if len(diags) != 0 {
		t.Fatalf("plain parse: %v", diags)
	}
	if b.Image.OGProvider() != "local" || len(b.Image.OG.Variants) != 0 {
		t.Errorf("absent image block: provider %q variants %d, want local/0", b.Image.OGProvider(), len(b.Image.OG.Variants))
	}
}
