---
title: "Agent-Gated Publishing: Why the Machine Locks PASS"
date: 2026-06-10
lastmod: 2026-06-11
tags: [agents, publishing, quality-gates]
---

![A ratchet mechanism illustrating one-way progress](/images/ratchet-gate.png)

When an agent writes and publishes a blog post end to end, the hard question is
not generation quality but verification authority. Generation is probabilistic;
verification must be deterministic, and only the deterministic side should be
allowed to declare a task done.

## The ratchet model

A quest is a set of items locked by a one-way ratchet. Once an item passes the
gate it can never silently regress: the remaining work only shrinks. The agent
proposes, the gate disposes — see the [Robots Exclusion Protocol](https://www.rfc-editor.org/rfc/rfc9309.html)
gate from an earlier run for how a citation check rejects unreachable sources.

A minimal gate rule looks like this:

```go
var ruleExample = gate.Rule{
	Meta: gate.RuleMeta{ID: "example", Level: gate.LevelFail},
	Check: func(ctx gate.Context) (bool, quest.Fact) {
		// fired == a violation was found

		return false, quest.Fact{}
	},
}
```

## Translation as a quest

Translations are items too: one article times N languages. Each language is
judged independently, so a failing Arabic submission never blocks a passing
Japanese one. The structural contract — headings, paragraph blocks, code
fences, links — is compared against the origin, as described in
[the first post](/posts/agent-gated-publishing/) of this series. Hugo's
multilingual mode is documented in the
[Hugo multilingual guide](https://gohugo.io/content-management/multilingual/).

## Sources

- [RFC 9309: Robots Exclusion Protocol](https://www.rfc-editor.org/rfc/rfc9309.html)
- [Hugo multilingual guide](https://gohugo.io/content-management/multilingual/)
