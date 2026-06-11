//ff:type feature=insight type=schema
//ff:what 인사이트 명세 1장 — 글 1편의 topic/stance/audience/section/tags/claims/non_goals/tone, 집필 게이트의 사람 입력
package insight

// Insight is the structured human input for one article (insight.yaml).
// Claims are the machine-checkable core; tone is a hint only (no gate).
type Insight struct {
	Topic    string   `yaml:"topic" json:"topic"`
	Stance   string   `yaml:"stance" json:"stance"`
	Audience string   `yaml:"audience" json:"audience"`
	Section  string   `yaml:"section" json:"section"`
	Tags     []string `yaml:"tags" json:"tags"`
	Claims   []Claim  `yaml:"claims" json:"claims"`
	NonGoals []string `yaml:"non_goals" json:"non_goals"`
	Tone     string   `yaml:"tone" json:"tone"`
}
