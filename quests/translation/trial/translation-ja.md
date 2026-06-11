---
title: "エージェントゲート型公開: なぜ機械がPASSをロックするのか"
date: 2026-06-10
lastmod: 2026-06-11
tags: [agents, publishing, quality-gates]
---

![一方向の前進を示すラチェット機構](/images/ratchet-gate.png)

エージェントがブログ記事を最初から最後まで書いて公開するとき、難しい問題は生成の
品質ではなく検証の権限である。生成は確率的であり、検証は決定的でなければならず、
タスク完了を宣言する権限は決定的な側にのみ与えられるべきだ。

## ラチェットモデル

クエストは一方向ラチェットでロックされるアイテムの集合である。アイテムが一度
ゲートを通過すれば、静かに退行することはない: 残りの作業は減るだけだ。エージェントは
提案し、ゲートが決定する — 到達不能なソースを引用チェックが拒否する例は、以前の回の
[Robots Exclusion Protocol](https://www.rfc-editor.org/rfc/rfc9309.html) ゲートを参照。

最小のゲートルールは次のような形だ:

```go
var ruleExample = gate.Rule{
	Meta: gate.RuleMeta{ID: "example", Level: gate.LevelFail},
	Check: func(ctx gate.Context) (bool, quest.Fact) {
		// fired == a violation was found

		return false, quest.Fact{}
	},
}
```

## クエストとしての翻訳

翻訳もアイテムである: 記事1本 × N言語。言語ごとに独立して判定されるため、アラビア語の
提出の失敗が日本語の提出の合格を妨げることはない。構造契約 — 見出し、段落ブロック、
コードフェンス、リンク — は原文と照合される。このシリーズの
[最初の記事](/ja/posts/agent-gated-publishing/)で説明した。Hugoの多言語モードは
[Hugo multilingual guide](https://gohugo.io/content-management/multilingual/)に
文書化されている。

## 出典

- [RFC 9309: Robots Exclusion Protocol](https://www.rfc-editor.org/rfc/rfc9309.html)
- [Hugo multilingual guide](https://gohugo.io/content-management/multilingual/)
