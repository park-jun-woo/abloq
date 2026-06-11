// Package translation embeds the translation-quest task tree (tasks.md) and
// the translation conventions (context.md) so pkg/quests/translation can
// render them into the `next` prompt of any blog instance without filesystem
// coupling. (quests/ is .ffignore'd like template/ — embed payload, not app
// code.)
package translation

import _ "embed"

// TasksMD is the T1~T3 task tree shown in the translation prompt.
//
//go:embed tasks.md
var TasksMD string

// ContextMD is the translation conventions document (invariants, heading map
// usage, RTL handling, cheese-defense prohibitions) shown in the prompt.
//
//go:embed context.md
var ContextMD string
