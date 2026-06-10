package template

import (
	"io/fs"
	"testing"
)

func TestFiles(t *testing.T) {
	fsys := Files()
	must := []string{
		"README.md",
		".gitignore",
		"layouts/_default/baseof.html",
		"layouts/_default/single.html",
		"layouts/partials/jsonld.html",
		"layouts/partials/hooks/social.html",
		"assets/css/main.css",
		"assets/js/particles.js",
		"deploy/archiver.md",
		"deploy/terraform/main.tf",
		"deploy/terraform/modules/site/main.tf",
		"deploy/terraform/modules/waf/main.tf",
		"deploy/terraform/modules/logs/main.tf",
	}
	for _, p := range must {
		if _, err := fs.Stat(fsys, p); err != nil {
			t.Errorf("template payload missing %s: %v", p, err)
		}
	}
}
