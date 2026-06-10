package archive

import (
	"encoding/json"
	"testing"
)

func TestPlanDeploy(t *testing.T) {
	resp, err := PlanDeploy(PlanDeployRequest{
		DeployID:     "dep-1",
		Changed:      []string{"https://a/", "https://b/"},
		ReceiptsJSON: `[{"kind":"wayback","target":"https://a/"}]`,
	})
	if err != nil {
		t.Fatalf("PlanDeploy: %v", err)
	}
	if resp.Planned != 5 {
		t.Errorf("Planned = %d, want 5 (6 pairs − 1 existing)", resp.Planned)
	}
	var items []map[string]any
	if err := json.Unmarshal(resp.ItemsJSON, &items); err != nil || len(items) != 5 {
		t.Errorf("ItemsJSON = %s (err=%v), want 5-item JSON array", resp.ItemsJSON, err)
	}
	for _, item := range items {
		if item["deploy_id"] != "dep-1" || item["status"] != "pending" {
			t.Errorf("item %v: want deploy_id dep-1, status pending", item)
		}
	}

	if _, err := PlanDeploy(PlanDeployRequest{DeployID: "d", ReceiptsJSON: "not json"}); err == nil {
		t.Error("invalid receipts JSON must fail")
	}
}
