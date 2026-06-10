//ff:func feature=visibility type=parser control=sequence topic=report
//ff:what BotTotals가 pkg/bots 분류로 training/search/fetch를 나누고 md를 분류 불문 누적, 사전 밖 봇은 버리는지 검증
package report

import "testing"

func TestBotTotals(t *testing.T) {
	m := BotTotals([]BotSum{
		{Bot: "GPTBot", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 7, MDHits: 2},
		{Bot: "ChatGPT-User", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 3},
		{Bot: "PerplexityBot", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 5, MDHits: 1},
		{Bot: "NotInDictionary", Lang: "ko", Section: "tech", Slug: "post-a", Hits: 99},
	})
	got := m["ko/tech/post-a"]
	want := Tally{Training: 7, Search: 5, Fetch: 3, MD: 3}
	if got != want {
		t.Errorf("want %+v, got %+v", want, got)
	}
	if len(BotTotals(nil)) != 0 {
		t.Error("empty input must yield empty map")
	}
}
