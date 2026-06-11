//ff:func feature=quest type=frame control=sequence
//ff:what blog.yaml 로드 + 대상 글 읽기·파싱 → 단일 글 게이트 Target 조립, Base nil(전량 신규 판정) 규약
//ff:why attachBaselines를 쓰지 않는다 — git 작업트리의 untracked 신규 글은 diff에 안 잡혀 기준선 룰이 침묵 통과하는 치즈 구멍. 집필 퀘스트는 항상 전량 신규로 판정한다 (Phase016)
package writing

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// assembleTarget builds the single-article gate target for one submission:
// the instance blog.yaml plus the parsed target article with no baseline
// (Base nil — the writing-quest convention). It also returns the article's
// raw bytes for insight matching.
func assembleTarget(p Payload) (*agate.Target, []byte, error) {
	b, diags, err := blogyaml.Load(filepath.Join(p.Root, "blog.yaml"))
	if err != nil {
		return nil, nil, err
	}
	if len(diags) > 0 {
		return nil, nil, fmt.Errorf("blog.yaml: %s", diags[0].String())
	}
	body, err := os.ReadFile(filepath.Join(p.Root, filepath.FromSlash(p.Article)))
	if err != nil {
		return nil, nil, fmt.Errorf("target article unreadable: %w", err)
	}
	art := &agate.Article{Lang: p.Lang, Section: p.Section, Slug: p.Slug,
		Path: p.Article, Doc: agate.ParseArticle(b, p.Lang, string(body))}
	return agate.NewTarget(p.Root, b, []*agate.Article{art}), body, nil
}
