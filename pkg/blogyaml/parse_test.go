//ff:func feature=blogyaml type=parser control=sequence
//ff:what 최소 blog.yaml에 스키마 v1 기본값(llms_txt/jsonld/임계값/deploy)이 주입되는지 검증
package blogyaml

import (
	"reflect"
	"testing"
)

func TestParseDefaults(t *testing.T) {
	src := []byte("site: {baseURL: https://example.com, title: T, author: A}\nlanguages: [ko]\nsections: [tech]\n")
	b, _, diags := Parse("blog.yaml", src)
	if len(diags) != 0 {
		t.Fatalf("want 0 diagnostics, got %v", diags)
	}
	if b.Geo.LlmsTxt != "auto" {
		t.Errorf("want default llms_txt auto, got %q", b.Geo.LlmsTxt)
	}
	if !reflect.DeepEqual(b.Geo.JSONLD, []string{"Article", "Person"}) {
		t.Errorf("want default jsonld [Article Person], got %v", b.Geo.JSONLD)
	}
	if b.Geo.FreshnessDays != 90 {
		t.Errorf("want default freshness_days 90, got %d", b.Geo.FreshnessDays)
	}
	if b.Geo.MinSources != 1 {
		t.Errorf("want default min_sources 1, got %d", b.Geo.MinSources)
	}
	if b.Geo.MinInternalLinks != 2 {
		t.Errorf("want default min_internal_links 2, got %d", b.Geo.MinInternalLinks)
	}
	if b.Deploy.Provider != "s3-cloudfront" {
		t.Errorf("want default provider s3-cloudfront, got %q", b.Deploy.Provider)
	}
	if !b.Deploy.IndexNow {
		t.Errorf("want default indexnow true")
	}
	if b.Deploy.Terraform {
		t.Errorf("want default terraform false")
	}
}
