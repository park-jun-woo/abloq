//ff:func feature=scan type=client control=sequence topic=evidence
//ff:what URL 1건 점검 — UA 명시 HEAD, HEAD 비허용 응답(405/501)이면 GET 폴백, 상태/오류를 ok·hard·soft로 환원
package evidence

import "net/http"

// checkURL probes one citation URL: a HEAD first (cheap, polite), then a GET
// when the server rejects HEAD as a method. Bodies are never read — only the
// status code matters, and rot needs persistence across scans anyway.
func (c *Checker) checkURL(raw string) string {
	target := c.overrideURL(raw)
	code, err := c.probe(http.MethodHead, target)
	if err != nil {
		return classifyErr(err)
	}
	if code == http.StatusMethodNotAllowed || code == http.StatusNotImplemented {
		code, err = c.probe(http.MethodGet, target)
		if err != nil {
			return classifyErr(err)
		}
	}
	return classify(code)
}
