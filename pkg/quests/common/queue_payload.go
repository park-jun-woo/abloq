//ff:type feature=quest type=schema topic=queue
//ff:what 큐 소비 Item.Payload — 인스턴스 루트(절대), 대상 글 경로, Key 부품, 전 언어 키, Seed 시점 고정된 큐 payload 사본
//ff:why 큐 payload는 Seed 시점에 Item으로 고정한다 — 게이트가 작업트리 큐 파일을 다시 읽으면 에이전트가 파일을 변조해 검사 범위를 넓히는 치즈에 열린다 (Phase018 계획)
package common

// QueuePayload is the persisted payload every queue-consuming quest item
// carries: the blog instance root (absolute, where blog.yaml lives), the
// root-relative target article path, the key parts (lang/section/slug), the
// per-language join keys, and a frozen copy of the queue file's payload map
// taken at Seed time — the gate never re-reads the working-tree queue file.
type QueuePayload struct {
	Root    string            `json:"root"`
	Article string            `json:"article"`
	Lang    string            `json:"lang"`
	Section string            `json:"section"`
	Slug    string            `json:"slug"`
	Keys    []string          `json:"keys"`
	Queue   map[string]string `json:"queue"`
}
