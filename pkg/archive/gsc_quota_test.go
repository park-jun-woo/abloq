//ff:func feature=archive type=client control=sequence
//ff:what gscQuota가 GSC_DAILY_QUOTA env를 읽고 미설정·0 이하·비정수면 200으로 폴백하는지 검증
package archive

import "testing"

func TestGscQuota(t *testing.T) {
	t.Setenv("GSC_DAILY_QUOTA", "")
	if got := gscQuota(); got != 200 {
		t.Errorf("default = %d, want 200", got)
	}
	t.Setenv("GSC_DAILY_QUOTA", "5")
	if got := gscQuota(); got != 5 {
		t.Errorf("env 5 = %d, want 5", got)
	}
	t.Setenv("GSC_DAILY_QUOTA", "-1")
	if got := gscQuota(); got != 200 {
		t.Errorf("negative = %d, want fallback 200", got)
	}
	t.Setenv("GSC_DAILY_QUOTA", "abc")
	if got := gscQuota(); got != 200 {
		t.Errorf("non-integer = %d, want fallback 200", got)
	}
}
