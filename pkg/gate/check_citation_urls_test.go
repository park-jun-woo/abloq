//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what 인용 추출 결과 검증 헬퍼 — URL 목록과 파일 라인 번호를 비교
package gate

import "testing"

func checkCitationURLs(t *testing.T, got []Citation, wantURLs []string, wantLine int) {
	t.Helper()
	if len(got) != len(wantURLs) {
		t.Fatalf("got %+v, want URLs %v", got, wantURLs)
	}
	for i, c := range got {
		if c.URL != wantURLs[i] || c.Line != wantLine {
			t.Errorf("citation[%d] = %+v, want URL %s at line %d", i, c, wantURLs[i], wantLine)
		}
	}
}
