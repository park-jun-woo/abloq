// Package cluster embeds the cluster-quest task tree (tasks.md) and the
// curation conventions (context.md) so pkg/quests/cluster can render them
// into the `next` prompt of any blog instance without filesystem coupling.
// (quests/ is .ffignore'd like template/ — embed payload, not app code.)
package cluster

import _ "embed"

// TasksMD is the T1~T3 task tree shown in the cluster prompt.
//
//go:embed tasks.md
var TasksMD string

// ContextMD is the curation conventions document (internal-link anchors,
// lastmod rules, cheese-defense prohibitions) shown in the prompt.
//
//go:embed context.md
var ContextMD string
