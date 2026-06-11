---
title: "에이전트 게이트 게시: 왜 기계가 PASS를 잠그는가"
date: 2026-06-10
lastmod: 2026-06-11
tags: [agents, publishing, quality-gates]
---

![일방향 진행을 보여주는 래칫 기구](/images/ratchet-gate.png)

에이전트가 블로그 글을 처음부터 끝까지 쓰고 게시할 때, 어려운 문제는 생성 품질이
아니라 검증 권한이다. 생성은 확률적이고 검증은 결정적이어야 하며, 작업 완료를
선언할 권한은 결정적인 쪽에만 주어져야 한다.

## 래칫 모델

퀘스트는 일방향 래칫으로 잠기는 아이템들의 집합이다. 아이템이 한번 게이트를
통과하면 조용히 퇴행할 수 없다: 남은 작업은 줄어들기만 한다. 에이전트는 제안하고
게이트가 결정한다 — 도달 불가 출처를 인용 검사가 거부하는 사례는 이전 회차의
[Robots Exclusion Protocol](https://www.rfc-editor.org/rfc/rfc9309.html) 게이트를 보라.

최소 게이트 룰은 이런 모양이다:

```go
var ruleExample = gate.Rule{
	Meta: gate.RuleMeta{ID: "example", Level: gate.LevelFail},
	Check: func(ctx gate.Context) (bool, quest.Fact) {
		// fired == a violation was found

		return false, quest.Fact{}
	},
}
```

## 퀘스트로서의 번역

번역도 아이템이다: 글 1편 × N개 언어. 언어마다 독립적으로 판정되므로 아랍어
제출의 실패가 일본어 제출의 통과를 막지 않는다. 구조 계약 — 헤딩, 문단 블록,
코드 펜스, 링크 — 은 원문과 대조되며, 이 시리즈의
[첫 글](/ko/posts/agent-gated-publishing/)에서 설명했다. Hugo의 다국어 모드는
[Hugo multilingual guide](https://gohugo.io/content-management/multilingual/)에
문서화되어 있다.

## 출처

- [RFC 9309: Robots Exclusion Protocol](https://www.rfc-editor.org/rfc/rfc9309.html)
- [Hugo multilingual guide](https://gohugo.io/content-management/multilingual/)
