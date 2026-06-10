//ff:func feature=gate type=frame control=sequence
//ff:what 골든 통합 — repo-pass 미니 저장소가 11룰 전부에서 진단 0건인지 검증
package gate

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestRunRepoPass(t *testing.T) {
	dir := filepath.Join("testdata", "repo-pass")
	b, diags, err := blogyaml.Load(filepath.Join(dir, "blog.yaml"))
	if err != nil || len(diags) != 0 {
		t.Fatalf("fixture blog.yaml: %v %v", err, diags)
	}
	got := Run(NewTarget(dir, b, Discover(dir, b)))
	if len(got) != 0 {
		t.Errorf("want 0 diagnostics for repo-pass, got %v", got)
	}
}
