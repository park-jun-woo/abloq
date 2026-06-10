//ff:func feature=cli type=parser control=sequence topic=citation
//ff:what loadQueriesFile이 YAML과 JSON 질의 파일을 같은 {id, query_text} 목록으로 읽고 빈 목록·빈 질의문·결손 파일은 에러인지 검증
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadQueriesFile(t *testing.T) {
	dir := t.TempDir()
	yamlPath := filepath.Join(dir, "queries.yaml")
	if err := os.WriteFile(yamlPath, []byte("- id: 1\n  query_text: best static blog\n- id: 2\n  query_text: agentic blogging\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	queries, err := loadQueriesFile(yamlPath)
	if err != nil {
		t.Fatalf("yaml: %v", err)
	}
	if len(queries) != 2 || queries[0].ID != 1 || queries[1].Text != "agentic blogging" {
		t.Errorf("yaml queries = %+v", queries)
	}

	jsonPath := filepath.Join(dir, "queries.json")
	if err := os.WriteFile(jsonPath, []byte(`[{"id":3,"query_text":"geo basics"}]`), 0o644); err != nil {
		t.Fatal(err)
	}
	queries, err = loadQueriesFile(jsonPath)
	if err != nil {
		t.Fatalf("json: %v", err)
	}
	if len(queries) != 1 || queries[0].ID != 3 || queries[0].Text != "geo basics" {
		t.Errorf("json queries = %+v", queries)
	}

	empty := filepath.Join(dir, "empty.yaml")
	if err := os.WriteFile(empty, []byte("[]\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := loadQueriesFile(empty); err == nil {
		t.Error("empty list accepted")
	}
	noText := filepath.Join(dir, "notext.yaml")
	if err := os.WriteFile(noText, []byte("- id: 9\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := loadQueriesFile(noText); err == nil {
		t.Error("entry without query_text accepted")
	}
	if _, err := loadQueriesFile(filepath.Join(dir, "missing.yaml")); err == nil {
		t.Error("missing file accepted")
	}
	bad := filepath.Join(dir, "bad.yaml")
	if err := os.WriteFile(bad, []byte(":	::not yaml"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := loadQueriesFile(bad); err == nil {
		t.Error("malformed yaml accepted")
	}
}
