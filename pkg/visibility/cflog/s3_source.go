//ff:type feature=visibility type=client topic=crawl
//ff:what S3 로그 소스 — ListObjectsV2·GetObject 2개 API만, stdlib sigv4 직접 서명, env 자격증명, Endpoint는 테스트 스텁 오버라이드
//ff:why 신규 의존성 0 유지(GSC SA JWT RS256 직접 구현 선례 — Phase008): AWS SDK 없이 필요한 두 호출만 서명한다. 자격증명은 전부 백엔드 env에만 — 에이전트에는 없다 (설계서 §3.3, Phase012)
package cflog

import "net/http"

// S3Source serves log objects from an S3 bucket under Prefix. Endpoint
// overrides the regional virtual-host URL for tests (a local stub); empty
// means https://<bucket>.s3.<region>.amazonaws.com. Credentials are the
// standard AWS env values, read once at OpenSource time.
type S3Source struct {
	Bucket       string
	Prefix       string
	Region       string
	AccessKey    string
	SecretKey    string
	SessionToken string
	Endpoint     string
	Client       *http.Client
}
