//ff:func feature=blogyaml type=parser control=sequence
//ff:what defaultBlog가 스키마 v1 기본값(geo 임계값/llms_txt/jsonld/deploy)을 채우는지 검증
package blogyaml

import (
	"reflect"
	"testing"
)

func TestDefaultBlog(t *testing.T) {
	b := defaultBlog()
	if b.Geo.LlmsTxt != "auto" {
		t.Errorf("want llms_txt auto, got %q", b.Geo.LlmsTxt)
	}
	if !reflect.DeepEqual(b.Geo.JSONLD, []string{"Article", "Person"}) {
		t.Errorf("want jsonld [Article Person], got %v", b.Geo.JSONLD)
	}
	if b.Geo.FreshnessDays != 90 {
		t.Errorf("want freshness_days 90, got %d", b.Geo.FreshnessDays)
	}
	if b.Geo.MinSources != 1 {
		t.Errorf("want min_sources 1, got %d", b.Geo.MinSources)
	}
	if b.Geo.MinInternalLinks != 2 {
		t.Errorf("want min_internal_links 2, got %d", b.Geo.MinInternalLinks)
	}
	if b.Deploy.Provider != "s3-cloudfront" {
		t.Errorf("want provider s3-cloudfront, got %q", b.Deploy.Provider)
	}
	if b.Deploy.Terraform {
		t.Errorf("want terraform false")
	}
	if !b.Deploy.IndexNow {
		t.Errorf("want indexnow true")
	}
}
