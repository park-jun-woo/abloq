//ff:func feature=blogyaml type=parser control=sequence
//ff:what defaultBlog가 스키마 v1 기본값(site 서브디렉토리/geo 임계값/llms_txt/jsonld/deploy)을 채우는지 검증
package blogyaml

import (
	"reflect"
	"testing"
)

func TestDefaultBlog(t *testing.T) {
	b := defaultBlog()
	if !b.Site.DefaultLangInSubdir {
		t.Errorf("want default_lang_in_subdir true by default")
	}
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
	if b.Geo.MinMeaningfulDiff != 10 {
		t.Errorf("want min_meaningful_diff 10, got %d", b.Geo.MinMeaningfulDiff)
	}
	w := b.Geo.PriorityWeights
	if w.Fetcher != 3 || w.Train != 1 || w.GSC != 1 || w.Citation != 2 {
		t.Errorf("want priority_weights {3 1 1 2} (fetcher highest), got %+v", w)
	}
	if w.Fetcher <= w.Train || w.Fetcher <= w.GSC || w.Fetcher <= w.Citation {
		t.Errorf("fetcher must carry the highest default weight: %+v", w)
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
