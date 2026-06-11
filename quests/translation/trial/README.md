# 시범 운행 기록 — Phase017 번역 퀘스트 (2026-06-11)

일회용 인스턴스(`abloq init /tmp/abloq-trial-017 --languages en,ko,ja,ar`, posts 단일
섹션, 구조 body→sources, 언어별 sources 헤딩: Sources/출처/出典/المصادر, min_sources
기본 1)에서 원문 1편 × 3언어 번역의 게이트 루프를 완주한 기록. 12언어 전수 대신
**RTL(ar)·CJK(ja) 포함 4언어**로 구조를 검증했다(계획 확정 — 12언어 박제는
context.md 규약). 발행 없음, 인스턴스는 기록 추출 후 폐기.

## 무대

- 원문(`origin-en.md`): 신규 글 1편 — 자유 H2 섹션 2개 + 인식 sources 섹션,
  Go 코드블록(내부 빈 줄 포함 — fence-aware splitter 검증점), 내부 글 링크 1개,
  외부 링크 2개(본문)+2개(출처 목록), 메인 이미지, date 2026-06-10 /
  lastmod 2026-06-11.
- scan: `abloq quest translation scan content/en/posts/agent-gated-publishing.md`
  → `seeded 3 item(s)` — 기본 언어(en)를 제외한 ko·ja·ar 매트릭스.

## 타임라인

1. **ko — submit #1 PASS**. 내부 링크를 `/ko/posts/agent-gated-publishing/`로 치환,
   `## 출처` 헤딩 맵 사용, date·lastmod 원문 이식. 전 룰(translation-parity,
   slug-consistency 스코프드, front-matter-schema, section-order, heading-canonical,
   min-sources, hugo-build) 통과.
2. **ja — submit #1 의도적 FAIL** (출처(出典) 섹션 누락 상태로 제출):
   ```
   ja/posts/agent-gated-publishing -> FAIL (state TODO)
     - translation-parity: content/ja/posts/agent-gated-publishing.md#headings
       expected="origin heading level sequence [2 2 2]" actual="[2 2] (외 3건)"
     - min-sources: content/ja/posts/agent-gated-publishing.md:1
       expected="the sources section lists at least geo.min_sources entries"
       actual="sources section missing — geo.min_sources requires >= 1 source(s)"
   ```
   translation-parity가 구조 훼손을 Fact로 잡았다: 헤딩 레벨 시퀀스 [2 2 2]→[2 2]
   (외 3건 = 인식 섹션 시퀀스 [sources]→[], 문단 블록 수, 외부 링크 multiset 누락 2건).
   RootCause = translation-parity (카탈로그 선두). tries 1 소진.
3. **ja — 수정** — Fact 반영: `## 出典` 섹션 + 출처 목록 2건 복원.
4. **ja — submit #2 PASS**: `ja/posts/agent-gated-publishing -> PASS (state PASS)`
   — FAIL 1회 후 잔여 2회 내 수렴(MaxTries=3). 시도 이력은 `session.json` log에 보존.
5. **ar — submit #1 PASS** (RTL). 내부 링크 `/ar/posts/...` 치환, `## المصادر` 헤딩,
   방향 제어 문자(U+202A~U+202E·U+2066~U+2069) 0건 확인(grep), 마크다운 구조
   문자·URL·코드블록은 LTR 그대로 — 패리티 전 항목 통과.

## 종료 상태

- `status`: TODO 0 / **PASS 3** / DONE 0 — 전 아이템 PASS.
- **hugo 빌드 0 에러**: 게이트 내 hugo-build 룰(매 제출, 인스턴스 전체)에 더해
  독립 실행 `hugo --quiet -d <tmp>` → exit 0. 빌드된 ar 페이지는
  `<html lang="ar" dir="rtl">`로 렌더되고 아랍어 본문이 들어 있음을 확인.
- **재scan 멱등**: 같은 원문으로 다시 scan → `seeded 0 item(s)` — 3개 번역의
  lastmod(원문 이식 = 2026-06-11)가 원문과 같아 갱신 불필요 판정(미생성).
- export: `translation-results.jsonl` — ko tries=0, ja tries=1(FAIL→PASS), ar tries=0.

## 파일

| 파일 | 내용 |
|---|---|
| `blog.yaml` | 시범 인스턴스 SSOT (4언어 + 언어별 sources 헤딩) |
| `origin-en.md` | 원문(기본 언어) — 게이트 비교 기준 |
| `translation-ko.md` / `translation-ja.md` / `translation-ar.md` | 게이트 PASS 최종 번역 3편 (ja는 FAIL 1회 후 수정본) |
| `session.json` | 래칫 세션 — ja의 try 1 FAIL Fact 이력 포함 |
| `translation-results.jsonl` | export 산출물 (emit-once) |
