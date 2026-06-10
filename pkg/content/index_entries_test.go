//ff:func feature=content type=parser control=sequence
//ff:what IndexEntries가 검증된 Blog로 IndexRepo와 동일한 인덱스를 내는지 검증 (재로드 없는 본체 동등성)
package content

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestIndexEntries(t *testing.T) {
	root := "testdata/fixture"
	b, diags, err := blogyaml.Load(root + "/blog.yaml")
	if err != nil || len(diags) > 0 {
		t.Fatalf("fixture blog.yaml: %v %v", err, diags)
	}
	got := IndexEntries(root, b)
	want, err := IndexRepo(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 || !reflect.DeepEqual(got, want) {
		t.Errorf("IndexEntries must equal IndexRepo:\n got %+v\nwant %+v", got, want)
	}
}
