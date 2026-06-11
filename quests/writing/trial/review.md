reviewer: claude-headless-reviewer-20260611

---

## Disposition — missing claim

- planted-claim: excluded — "robots.txt 캐싱 TTL을 늘리면 서버 비용이 유의미하게 준다"는 본문 어디에도 언급되지 않는다. 작업 로그의 의도적 제외 결정이 타당하다: 이 주장은 인프라 비용 최적화 관점의 운용 조언으로, 본문 stance("REP는 신호다")와 topic 흐름에서 이질적이다. non_goals에 명시적으로 열거되지는 않았지만, "특정 AI 업체별 차단 가이드"·"법적 효력 논의" 수준의 스코프 이탈에 해당한다고 판단한다. 다루지 않는 것이 맞다.

---

## 출현 claim 지지 여부 소견

**rep-definition** (anchors: "robots.txt", "crawler"): 본문 도입부 — "a crawler fetches the robots.txt file at the site root and decides for itself which URIs it may access" — 이 정의를 직접 진술하며 클라이언트 측 자기 판정임을 명시한다. Claim을 충분히 지지한다.

**rep-standardized-2022** (anchors: "RFC 9309", "1994"): "originally defined by Martijn Koster in 1994"와 "Only in September 2022 did the IETF publish RFC 9309" 두 사실이 동일 단락에서 대비된다. 출처도 RFC 9309 링크로 인라인 제공된다. Claim을 충분히 지지하며 requires_source 요건도 충족한다.

**rep-not-access-control** (anchors: "not a form of access authorization", "voluntary"): RFC 9309 원문을 직접 인용("not a form of access authorization")하고, "compliance is voluntary"를 명시한다. 이어지는 단락에서 서버 측 강제 수단(인증, IP 차단, signed URL)을 열거해 주장을 실질적으로 보강한다. Claim을 충분히 지지하며 requires_source 요건도 충족한다.

**ai-crawlers-geo** (anchors: "AI crawler", "GEO"): "The AI Crawler Era" 섹션에서 AI 크롤러를 훈련·검색·실시간 fetch 세 범주로 구분하고, 각각이 visibility와 control 사이의 트레이드오프를 달리 함을 설명한다. "GEO (Generative Engine Optimization) decision-making" 구절에서 robots.txt 정책이 GEO 의사결정의 일부임을 명시한다. requires_source 없음으로 설정된 claim이며, 본문 전개로 충분히 지지된다.

---

## non_goals 이탈 여부 소견

**특정 AI 업체별 크롤러 차단 가이드**: 본문은 GPTBot, ClaudeBot, Bingbot 등 개별 user-agent를 일절 언급하지 않는다. AI 크롤러를 범주 수준(훈련·검색·실시간)에서만 다루며 차단 레시피를 제공하지 않는다. 이탈 없음.

**법적 효력(저작권·약관) 논의**: 저작권, 이용약관, 법적 구속력에 대한 언급이 없다. 이탈 없음.

---

## 종합

anchored 4건 모두 본문이 claim을 실질적으로 지지한다. planted-claim 제외는 topic 흐름·stance 정합성 측면에서 적절하다. 톤(단정적, 표준 문서 인용 중심)도 RFC 원문 직접 인용 및 "The single most common operational mistake" 같은 단언 형식과 일치한다. 이의 없음.
