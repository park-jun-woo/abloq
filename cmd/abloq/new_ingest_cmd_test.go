//ff:func feature=cli type=command control=iteration dimension=1 topic=crawl
//ff:what ingest 부모 명령이 Use를 선언하고 crawl 서브커맨드를 등록하는지 검증
package main

import "testing"

func TestNewIngestCmd(t *testing.T) {
	cmd := newIngestCmd()
	if cmd.Use != "ingest" {
		t.Errorf("Use = %q, want \"ingest\"", cmd.Use)
	}
	found := false
	for _, sub := range cmd.Commands() {
		if sub.Name() == "crawl" {
			found = true
		}
	}
	if !found {
		t.Error("crawl subcommand not registered")
	}
}
