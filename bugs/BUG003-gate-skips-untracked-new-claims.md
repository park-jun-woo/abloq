# BUG003 — gate가 untracked·committed 신규 글의 근거 룰을 조용히 스킵한다

> 상태: OPEN · 심각도: High(검증 공백, PASS 사칭) · 발견: abloq 소개 글 도그푸드 중 실측

## 한 줄 요약

repo 레벨 `abloq gate`의 신규-주장/신규-인용 기준선이 **`git diff --name-only HEAD`(추적 파일만)**라, 새 글이 **untracked-unstaged**이거나 이미 **committed**이면 그 글의 `citation-exists`·`numeric-claim-sourced`가 **조용히 검사를 건너뛴다.** 게이트는 "N article(s) pass"를 출력해 *검사한 것처럼* 보이지만 실제로는 신규 글의 근거를 한 건도 보지 않는다 — PASS 사칭.

## 증상 (실측 — abloq 소개 글 도그푸드)

죽은 인용(`https://github.com/park-jun-woo/abloq`, 당시 HTTP 404)을 포함한 새 글 `content/ko/tech/abloq.md`를 인스턴스에 두고 게이트:

```
# (1) 글이 untracked-unstaged 상태
$ abloq gate --rule citation-exists .
/tmp/.../hugo: 733 article(s) pass the gate          # ← 404 링크가 있는데 PASS

# (2) git add 로 스테이징
$ git add content/ko/tech/abloq.md
$ abloq gate --rule citation-exists .
content/ko/tech/abloq.md:124 [citation-exists]
  citation URL https://github.com/park-jun-woo/abloq is not reachable (HTTP 404)
abloq: 1 gate violation(s) found                     # ← 비로소 잡음
```

스테이징 전에는 죽은 링크가 통과했다. 사용자가 새 글을 `git add` 하지 않고 게이트를 돌리면(흔한 일), 또는 CI가 **이미 커밋된** 글에 게이트를 돌리면, 근거 룰이 공허 통과한다.

## 근본 원인

- `pkg/gate/git_changed_set.go:13` — `git diff --relative --name-only HEAD`. 이건 **추적 파일 중 HEAD와 다른 것**만 낸다. `git diff HEAD`는 **untracked 파일을 포함하지 않는다.**
- `pkg/gate/attach_baselines.go` — changed set에 든 글에만 HEAD 원본(`Base`)을 파싱해 부착. 그 밖의 글은 현재본을 원본으로 공유 → `Base == Doc`.
- `pkg/gate/new_citations.go` · `new_claims.go` — `Base == Doc`면 신규 목록은 **nil**(미변경 글은 신규 0). `citation-exists`·`numeric-claim-sourced`는 이 신규 목록만 판정한다.

git 상태별 귀결:

| 새 글의 git 상태 | `git diff HEAD`에 잡힘? | Base | 근거 룰 |
|---|---|---|---|
| **untracked-unstaged** | ❌ | Doc(==자기 자신) | **스킵** (구멍 #1) |
| staged-uncommitted | ✅ | nil (HEAD에 없음) → 전부 신규 | **검사됨** ✓ |
| **committed** (== HEAD) | ❌ | Doc | **스킵** (구멍 #2 — CI no-op) |

→ 근거 룰의 **가시 윈도우가 좁다**: *추적-수정* 또는 *스테이징-미커밋* 글뿐. 이게 정확히 집필/래칫의 더티-워크트리 윈도우(MANUAL §4.3)다. 그 밖 — 미스테이징 신규 초안, **committed 코드에 도는 CI** — 에선 근거 룰이 공허하다. (BUG001에서 관찰한 "evidence 룰이 사실상 off"의 메커니즘적 정체.)

## 가장 나쁜 점 — 침묵

`733 article(s) pass the gate`는 "전부 검사하고 통과"로 읽힌다. 실제로는 **신규 글의 인용을 0건 검사**한 것이다. 검사를 안 했다는 신호가 없다. *PASS를 사칭하는 게이트* — 게이트가 낼 수 있는 가장 나쁜 거짓이다.

## 개선안

1. **untracked를 changed set에 포함** (`git_changed_set.go`):
   `git diff --name-only HEAD` ∪ `git ls-files --others --exclude-standard`. untracked 글은 HEAD 원본이 없으니 `Base = nil` → "신규 글은 전부 신규" 경로로 자연히 검사된다. (구멍 #1 해소)

2. **`--base <ref>` 기준선 옵션** (gate 진입 + `git_show.go`):
   HEAD 대신 임의 ref(PR base 브랜치, `merge-base origin/main HEAD`)를 기준선으로. CI가 **커밋된** 새 콘텐츠를 base 대비 신규로 판정할 수 있다. (구멍 #2 해소)

3. **침묵 금지 — 스코프 보고**:
   유효 신규-주장/신규-인용 스코프가 0이거나 working tree가 HEAD와 동일하면, 게이트가 그 사실을 출력한다 — 예: `evidence rules: 0 new article(s) in scope (baseline=HEAD, tree clean)`. "no silent caps — 드롭한 것을 로그하라" 원칙. 0건 검사를 PASS로 위장하지 않는다.

## 수용 기준

- [ ] untracked-unstaged 새 글의 죽은 인용/무출처 수치 주장이 `git add` 없이도 FAIL한다(위 증상 (1)이 FAIL로 뒤집힘).
- [ ] `abloq gate --base <ref>`가 ref 대비 신규 콘텐츠를 판정한다(committed CI 시나리오).
- [ ] 유효 신규 스코프가 0이면 게이트가 스코프 규모를 명시 출력한다(침묵 PASS 금지).
- [ ] 기존 동작(추적-수정·스테이징 글) 무변경 — 회귀 없음.
- [ ] 골든/단위 테스트: untracked 픽스처가 근거 룰에 걸리는 케이스 추가.

## 변경 파일 (예상)

| 경로 | 변경 |
|---|---|
| `pkg/gate/git_changed_set.go` | untracked 합집합(`ls-files --others --exclude-standard`) |
| `pkg/gate/attach_baselines.go` | untracked → `Base=nil` 부착 |
| `pkg/gate/git_show.go`, `article.go`, gate 진입부 | `--base <ref>` 기준선 ref 플럼빙(기본 HEAD) |
| `cmd/abloq/new_gate_cmd.go`, `run_gate.go` | `--base` 플래그 + 스코프 0건 보고 |
| `MANUAL.md` §4.3/§5 | 기준선·스코프·CI 사용법 |
| 위 각 `*_test.go` | untracked·base-ref·zero-scope 케이스 |

## 관련 발견 → [BUG004](./BUG004-citation-receipt-stale-failure.md)

같은 도그푸드에서 **citation 영수증 캐시 스테일**도 겪었다 — 고쳐진 URL이 24h 캐시 때문에 계속 FAIL. 별건으로 분리해 [BUG004](./BUG004-citation-receipt-stale-failure.md)로 작성했다. 본 버그(검증 *공백*)와 같은 파일군(`citation_diag`/`citation-exists`)의 *반대 방향* 결함(낡은 실패를 못 놓음)이다.

## 참고

- 실측 transcript: abloq 소개 글 도그푸드(2026-06-11). untracked → "733 pass"(404 링크 포함), staged → 404 검출.
- 동류 관찰: `BUG001-llms-txt-not-curation-grade.md`는 다른 표면이지만, 본 버그는 거기서 스친 "evidence 룰 사실상 off"의 git 메커니즘적 원인이다.
- 기준선 룰 설계 의도: `pkg/gate/new_claims.go` ff 주석 — "게이트는 이번 작업이 추가한 주장만 판정, 코퍼스 전수 감사는 스캐너 소관". 의도는 옳다; 다만 "이번 작업"의 git 정의가 너무 좁다.
