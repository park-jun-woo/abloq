//ff:func feature=quest type=parser control=sequence
//ff:what 테스트 헬퍼 — 전 룰 통과 픽스처: 게이트 클린 글 원문과 insight.yaml(anchors 전부 본문 출현) 상수 제공
package writing

const passArticleMD = `---
title: "Test Article"
date: 2026-06-01
lastmod: 2026-06-02
tags: [test]
---

![main](cover.png)

*Image: by Tester*

This body mentions the alpha anchor.

It also covers the bravo anchor.

## Sources

- Internal style guide

## Changelog

- 2026-06-02 first version
`

const passInsightYAML = `topic: "test topic"
stance: "test stance"
audience: "testers"
section: posts
tags: [test]
claims:
  - id: c1
    text: "claim one"
    kind: claim
    requires_source: true
    anchors: ["alpha anchor"]
  - id: c2
    text: "claim two"
    kind: claim
    requires_source: false
    anchors: ["bravo anchor"]
non_goals: ["benchmarks"]
tone: "neutral"
`

func passFixtures() (article, insightSpec string) { return passArticleMD, passInsightYAML }
