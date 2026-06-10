//ff:type feature=queueio type=schema
//ff:what exporter 설정 — 블로그 저장소 origin URL, 전용 작업 클론 경로, 커밋 author (env에서 주입)
package queueio

// Config drives the exporter's git work clone. RepoURL is the blog repository
// origin (deploy-key SSH URL in production, file:// bare repo in tests);
// Workdir is the dedicated work clone — never the read-only /blog mount.
type Config struct {
	RepoURL     string
	Workdir     string
	AuthorName  string
	AuthorEmail string
}
