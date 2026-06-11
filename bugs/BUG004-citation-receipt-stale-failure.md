# BUG004 — 고친 인용 URL이 영수증 캐시 때문에 최대 24h 동안 계속 FAIL

> 상태: OPEN · 심각도: Low(자가 해소 가능, 단 혼란·발행 지연 유발) · 발견: BUG003 디버깅 중 실측

## 한 줄 요약

`citation-exists`는 URL 검증 결과를 24h 캐시(`.abloq/citation-receipts.json`)하는데, **실패(`broken`·`meta-mismatch`) verdict도 캐시한다.** 그래서 죽었던 URL을 고쳐도(404→200) 게이트는 캐시된 옛 실패를 재사용해 **최대 24h 동안 FAIL을 반복**한다. "고쳤는데 왜 아직 빨갛지?"

## 증상 (실측 — abloq repo public 전환)

1. abloq repo가 비공개일 때 `https://github.com/park-jun-woo/abloq` → HTTP 404 → 영수증에 `{verdict:"broken"}` 기록.
2. repo를 **public 전환** → `curl`로 GET·HEAD 모두 **200** 확인.
3. 그런데 게이트는 **여전히 404 FAIL**:
   ```
   content/ko/tech/abloq.md:124 [citation-exists]
     citation URL https://github.com/park-jun-woo/abloq is not reachable (HTTP 404)
   ```
4. `.abloq/citation-receipts.json` **삭제 후** 재실행 → **733/733 PASS.**

URL은 살아 있었다. 게이트만 24h 묵은 "닫힘" 팻말을 붙여 두고 있었다.

## 재현

```bash
# 1) 죽은 URL을 인용한 글로 게이트 → broken 캐시됨
abloq gate --rule citation-exists .
# 2) URL을 고친다(서버에서 200이 되게)
# 3) 다시 게이트 → 여전히 FAIL (캐시 TTL 24h 내)
abloq gate --rule citation-exists .
# 4) rm .abloq/citation-receipts.json 후 → PASS
```

## 근본 원인

`pkg/gate/citation_diag.go`:

```go
r, cached := rcpts[c.URL]
if !cached || now.Sub(r.CheckedAt) >= citationTTL {   // citationTTL = 24h
    verdict, detail := verifyCitation(client, c)
    r = receipt{CheckedAt: now, Verdict: verdict, Detail: detail}
    if verdict != "retry" {                            // ← ok·broken·meta-mismatch 전부 캐시
        rcpts[c.URL] = r
    }
}
```

- `retry`(네트워크 오류·5xx 일시 실패)만 캐시 제외하고, **`broken`(404 등)·`meta-mismatch`(제목 불일치)는 `ok`와 동일하게 24h 캐시**한다.
- 캐시 목적은 정당하다 — 매 게이트 실행마다 모든 URL을 다시 HTTP 치면 느리고 서버를 두드린다. 문제는 **성공과 실패의 수명을 같게 둔 것**이다. 성공은 잘 안 바뀌지만, **실패는 "고치려고" 존재하는 상태**다 — 짧게 살아야 한다.

## 개선안

1. **실패는 장기 캐시하지 않는다** (권장):
   `ok`만 `citationTTL`(24h)로 캐시. `broken`·`meta-mismatch`는 캐시 안 하거나 짧은 TTL(예: 수 분). 고친 URL이 다음 실행에서 곧장 재검증된다.
2. **`--refresh` / `--no-cache` 플래그**: 캐시 무시하고 강제 재검증. 수동 탈출구.
3. (보강) FAIL 진단에 캐시 시각 노출 — `(cached 2026-06-11T16:13, --refresh to recheck)` — "스테일일 수 있음"을 사용자가 인지.

## 수용 기준

- [ ] 죽었다가 200으로 고쳐진 URL이 **캐시 파일 수동 삭제 없이** 다음 게이트 실행에서 PASS한다.
- [ ] `ok` 캐시는 유지돼 반복 실행 성능이 보존된다(매번 전 URL 재요청 안 함).
- [ ] `abloq gate --refresh`가 캐시를 무시하고 전건 재검증한다.
- [ ] `retry`(일시 실패) 비캐시 동작은 무변경.
- [ ] 테스트: broken→fixed 전이가 캐시 삭제 없이 PASS로 뒤집히는 케이스.

## 변경 파일 (예상)

| 경로 | 변경 |
|---|---|
| `pkg/gate/citation_diag.go` | verdict별 캐시 정책 — `ok`만 24h, 실패는 단기/무캐시 |
| `pkg/gate/rule_citation_exists.go` | `citationTTL` 분리(okTTL vs failTTL) 또는 캐시 쓰기 게이팅 |
| `cmd/abloq/new_gate_cmd.go`, `run_gate.go` | `--refresh`/`--no-cache` 플래그 |
| 해당 `*_test.go` | broken→fixed 전이 케이스 |

## 영향 / 우선순위

- **Low** — `.abloq/citation-receipts.json` 삭제로 즉시 자가 해소, 데이터 손실 없음.
- 다만 **신뢰 비용**이 있다: 고친 뒤에도 FAIL이 뜨면 "내 수정이 안 먹었나"와 "캐시가 묵었나"를 구분할 수 없어, 멀쩡한 발행이 최대 24h 막히거나 불필요한 디버깅을 부른다.
- BUG003(검증 *공백*, High)과 짝을 이루는 같은 파일군(`citation_diag`/`citation-exists`)의 *반대 방향* 결함 — 한쪽은 검사를 안 하고, 한쪽은 낡은 실패를 못 놓는다.

## 참고

- 실측 transcript: abloq repo public 전환(2026-06-11). 게이트 404 반복 → `.abloq/citation-receipts.json` 삭제 후 733/733 PASS.
- 설계 ff 주석: `pkg/gate/citation_diag.go` — "24h 내 영수증은 재검증 생략, retry는 캐시하지 않고 RETRY로 보고". 의도는 옳다; 실패 verdict를 성공과 같은 수명으로 둔 것만 문제.
- 연관: [BUG003](./BUG003-gate-skips-untracked-new-claims.md) "관련 발견"에서 분리.
