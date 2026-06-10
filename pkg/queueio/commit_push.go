//ff:func feature=queueio type=client control=sequence
//ff:what quests/queue 변경분 커밋·푸시 — commitPushPath 위임(경로·메시지 고정), 동일 내용이면 no-op
package queueio

// commitPush commits and pushes the quests/queue changes of the work clone —
// the queue-export specialization of commitPushPath. The export contract
// (idempotent no-op commits, one rebase retry) is unchanged from Phase009.
func commitPush(cfg Config) error {
	_, err := commitPushPath(cfg, "quests/queue", "abloqd: queue export")
	return err
}
