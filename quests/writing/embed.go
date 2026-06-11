// Package writing embeds the writing-quest task tree (tasks.md) and the
// authoring conventions (context.md) so pkg/quests/writing can render them
// into the `next` prompt of any blog instance without filesystem coupling.
// (quests/ is .ffignore'd like template/ — embed payload, not app code.)
package writing

import _ "embed"

// TasksMD is the T1~T4 task tree shown in the authoring prompt.
//
//go:embed tasks.md
var TasksMD string

// ContextMD is the authoring conventions document (citation format, worklog,
// cheese-defense prohibitions) shown in the authoring prompt.
//
//go:embed context.md
var ContextMD string
