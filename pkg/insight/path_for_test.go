//ff:func feature=insight type=parser control=sequence
//ff:what 저장 규약 검증 — 번들은 디렉토리의 insight.yaml, 플랫은 {어간}.insight.yaml 사이드카
package insight

import (
	"path/filepath"
	"testing"
)

func TestPathFor(t *testing.T) {
	bundle := filepath.Join("content", "en", "tech", "post-a", "index.md")
	if got := PathFor(bundle); got != filepath.Join("content", "en", "tech", "post-a", "insight.yaml") {
		t.Errorf("want bundle insight.yaml next to index.md, got %q", got)
	}
	flat := filepath.Join("content", "en", "tech", "ratchet-pattern.md")
	if got := PathFor(flat); got != filepath.Join("content", "en", "tech", "ratchet-pattern.insight.yaml") {
		t.Errorf("want flat sidecar {stem}.insight.yaml, got %q", got)
	}
}
