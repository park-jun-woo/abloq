// Package evidence embeds the evidence-quest task tree (tasks.md) and the
// sourcing conventions (context.md) so pkg/quests/evidence can render them
// into the `next` prompt of any blog instance without filesystem coupling.
// (quests/ is .ffignore'd like template/ — embed payload, not app code.)
package evidence

import _ "embed"

// TasksMD is the T1~T3 task tree shown in the evidence prompt.
//
//go:embed tasks.md
var TasksMD string

// ContextMD is the sourcing conventions document (claim-scope, citation
// rules, lastmod policy, cheese-defense prohibitions) shown in the prompt.
//
//go:embed context.md
var ContextMD string
