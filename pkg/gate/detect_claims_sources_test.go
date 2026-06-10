//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what 출처 섹션 면제 검증 — sources 내부 수치 목록행 면제, 섹션 밖 본문 검출 유지, blog.yaml에 언어 헤딩 미선언이면 면제 없음
//ff:why Phase010 parkjunwoo.com 사본 1회전에서 출처 명기 목록행 자체가 무출처 주장으로 검출되는 오탐 발견 — 보정 회귀 기준
package gate

import "testing"

func TestDetectClaimsSourcesExempt(t *testing.T) {
	b := loadGateBlog(t)
	body := "Conversion improved by 38% after the rollout.\n\n## Sources\n\n" +
		"- Carnegie Mellon MSR 2026 — 41% permanent increase in throughput\n" +
		"- Example Journal 2025 — latency fell 120ms\n"
	t.Run("sources rows exempt, body claim still detected", func(t *testing.T) {
		claims := DetectClaims(ParseArticle(b, "en", body))
		if len(claims) != 1 {
			t.Fatalf("want exactly the body claim (sources rows exempt), got %d: %+v", len(claims), claims)
		}
		if want := "Conversion improved by 38% after the rollout."; claims[0].Text != want {
			t.Errorf("claims[0].Text = %q, want %q", claims[0].Text, want)
		}
	})
	t.Run("no sources heading declared for the language: no exemption", func(t *testing.T) {
		// "ja" is absent from structure.headings.sources, so "## Sources" is
		// not recognized and the conservative default keeps both list rows.
		claims := DetectClaims(ParseArticle(b, "ja", body))
		if len(claims) != 3 {
			t.Fatalf("want 3 claims (body + 2 unexempted rows), got %d: %+v", len(claims), claims)
		}
	})
}
