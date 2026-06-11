//ff:func feature=image type=parser control=sequence
//ff:what parseGeminiImage의 빈 candidates·깨진 base64·깨진 JSON 에러 경로 검증
package ogprovider

import "testing"

func TestParseGeminiImage(t *testing.T) {
	if _, err := parseGeminiImage([]byte(`{"candidates":[]}`)); err == nil {
		t.Error("empty candidates expected error, got nil")
	}
	if _, err := parseGeminiImage([]byte(`{"candidates":[{"content":{"parts":[{"inlineData":{"data":"%%%"}}]}}]}`)); err == nil {
		t.Error("broken base64 expected error, got nil")
	}
	if _, err := parseGeminiImage([]byte(`not json`)); err == nil {
		t.Error("broken json expected error, got nil")
	}
}
