//ff:func feature=quest type=parser control=sequence
//ff:what 테스트 헬퍼 — 패리티 통과 픽스처: 원문(en)과 클린 번역(ko) 상수 제공 (자유 헤딩·코드블록·내부/외부 링크·이미지·sources 포함)
package translation

const originMD = `---
title: "Hello Gate"
date: 2026-06-01
lastmod: 2026-06-03
tags: [test]
---

![cover](/images/cover.png)

Intro paragraph with an [external spec](https://example.org/spec) link.

## Setup

Read the [first post](/posts/first-post/) before you begin.

` + "```bash" + `
echo "do not translate"

still the same block
` + "```" + `

## Sources

- [Example Spec](https://example.org/spec)
`

const cleanKoMD = `---
title: "안녕 게이트"
date: 2026-06-01
lastmod: 2026-06-03
tags: [test]
---

![cover](/images/cover.png)

[외부 명세](https://example.org/spec) 링크가 있는 도입 문단.

## 준비

시작 전에 [첫 글](/ko/posts/first-post/)을 읽어라.

` + "```bash" + `
echo "do not translate"

still the same block
` + "```" + `

## 출처

- [Example Spec](https://example.org/spec)
`

func passPair() (origin, ko string) { return originMD, cleanKoMD }
