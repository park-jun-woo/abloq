// Package refresh embeds the refresh-quest task tree (tasks.md) and the
// refresh conventions (context.md) so pkg/quests/refresh can render them
// into the `next` prompt of any blog instance without filesystem coupling.
// (quests/ is .ffignore'd like template/ — embed payload, not app code.)
package refresh

import _ "embed"

// TasksMD is the T1~T3 task tree shown in the refresh prompt.
//
//go:embed tasks.md
var TasksMD string

// ContextMD is the refresh conventions document (stale-fact replacement,
// lastmod rules, claim handling, cheese-defense prohibitions) shown in the
// prompt.
//
//go:embed context.md
var ContextMD string
