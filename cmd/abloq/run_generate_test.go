//ff:func feature=cli type=command control=iteration dimension=1
//ff:what abloq generate가 파생물 4종을 기록하고 2회 연속 실행이 no-op(바이트 동일)인지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRunGenerate(t *testing.T) {
	dir := writeBlogFixture(t)
	var out bytes.Buffer
	if err := runGenerate(&out, dir); err != nil {
		t.Fatalf("runGenerate: %v\noutput: %s", err, out.String())
	}
	paths := []string{"hugo.toml", "static/robots.txt", "static/llms.txt", "data/jsonld.json"}
	first := map[string][]byte{}
	for _, p := range paths {
		data, err := os.ReadFile(filepath.Join(dir, p))
		if err != nil {
			t.Fatalf("%s not generated: %v", p, err)
		}
		first[p] = data
	}
	if err := runGenerate(&out, dir); err != nil {
		t.Fatalf("second runGenerate: %v", err)
	}
	for _, p := range paths {
		again, err := os.ReadFile(filepath.Join(dir, p))
		if err != nil || !bytes.Equal(again, first[p]) {
			t.Errorf("%s changed on second generate (err %v)", p, err)
		}
	}
}
