//ff:func feature=quest type=rule control=sequence
//ff:what section-order 어댑터 발동 검증 — changelog가 sources보다 앞서는 글에서 Fact 매핑
package writing

import "testing"

const outOfOrderMD = `---
title: "Test Article"
date: 2026-06-01
lastmod: 2026-06-02
tags: [test]
---

![main](cover.png)

*Image: by Tester*

Body text.

## Changelog

- 2026-06-02 first version

## Sources

- Internal style guide
`

func TestAdaptSectionOrder(t *testing.T) {
	root := writeInstance(t)
	fired, fact := fireRule(t, adaptRule("section-order"), subWith(t, root, outOfOrderMD, ""))
	if !fired {
		t.Fatal("section-order: want fired when changelog precedes sources")
	}
	if fact.Where == "" || fact.Actual == "" {
		t.Errorf("Fact incomplete: %+v", fact)
	}
}
