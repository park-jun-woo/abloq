//ff:func feature=archive type=client control=iteration dimension=1
//ff:what 배포 1건의 pending 영수증 계획 — 변경 URL × kind 전개에서 이미 영수증 있는 페어 제외 (deploy_id 멱등)
//ff:why 멱등 키는 (deploy_id, kind, target)이다 — URL만으로 거르면 다음 배포의 같은 URL 재변경도 걸러져 "변경 시점마다 증거"가 깨진다 (Phase008)
package archive

// PlanDeploy returns the receipts to record as pending for one deploy:
// every (kind, target) pair of the changed URLs minus the pairs that
// already hold a receipt for this deploy_id. Re-webhooking the same deploy
// therefore plans nothing new, while a later deploy of the same URL plans
// fresh evidence.
func PlanDeploy(deployID string, changed []string, existing []Existing) []Item {
	seen := existingSet(existing)
	planned := []Item{}
	for _, item := range planItems(deployID, changed) {
		if seen[item.Kind+"\n"+item.Target] {
			continue
		}
		planned = append(planned, item)
	}
	return planned
}
