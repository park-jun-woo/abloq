//ff:type feature=visibility type=client topic=crawl
//ff:what 로그 소스 추상화 — 키 목록(prefix·afterKey)과 키 1개의 본문 스트림, 로컬 디렉토리와 S3가 같은 계약을 구현
//ff:why 테스트·Hurl·CLI 단발 분석은 로컬 디렉토리로, 본번은 S3로 — 수집기는 소스를 모른다. afterKey는 리스팅 최적화일 뿐 커서가 아니다: CF 키의 시간 뒤 접미사가 랜덤이라 start-after 커서는 지연 배달 파일을 영구 누락시킨다 (Phase012)
package cflog

import "io"

// Source lists log object keys and opens one object's body. List returns
// keys under prefix, lexicographically ascending, strictly after afterKey
// (empty afterKey = from the start). Get streams one object; the caller
// closes it.
type Source interface {
	List(prefix, afterKey string) ([]string, error)
	Get(key string) (io.ReadCloser, error)
}
