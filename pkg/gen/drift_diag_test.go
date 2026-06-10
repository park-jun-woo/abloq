//ff:func feature=gen type=rule control=sequence topic=drift
//ff:what driftDiag가 첫 불일치 라인 번호와 기대/실제 줄 내용을 담은 진단을 만드는지 검증
package gen

import "testing"

func TestDriftDiag(t *testing.T) {
	d := driftDiag("static/robots.txt", "robots-policy-match",
		[]byte("User-agent: GPTBot\nDisallow:\n"), []byte("User-agent: GPTBot\nDisallow: /\n"))
	checkDriftDiagFields(t, d, "static/robots.txt", 2, "robots-policy-match",
		`want "Disallow:", got "Disallow: /"`)
}
