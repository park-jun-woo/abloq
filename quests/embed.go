// Package quests embeds the shared queue-consumption protocol document
// (_queue-protocol.md) so the queue-consuming quest packages
// (pkg/quests/{refresh,evidence,cluster}) can render it into their `next`
// prompts without filesystem coupling. (quests/ is .ffignore'd like
// template/ — embed payload, not app code.)
package quests

import _ "embed"

// ProtocolMD is the shared queue-consumption protocol (item = queue file,
// the frozen submit/commit order, queue-scope, claim-line conventions,
// cheese-defense principles) shown in every queue quest prompt.
//
//go:embed _queue-protocol.md
var ProtocolMD string
