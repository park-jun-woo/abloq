# insight.yaml — 인사이트 명세 작성 가이드

글 1편당 1장. 에이전트가 자료수집·집필·퇴고를 대행하기 전에, **사람이 결정한 인사이트**를 기계가 대조할 수 있는 형태로 적는 양식이다. 여기 적은 claims가 집필 게이트(REVIEW)의 기준이 된다 — 명세에 없는 주장은 글의 본론이 될 수 없고, 명세에 있는 주장은 본문에 대응해야 한다.

## 작성 요령 (비개발자용)

1. **주장 하나 = claim 하나.** 글에서 말하고 싶은 문장을 한 줄씩 쪼갠다. 한 claim에 두 가지 말을 넣지 않는다.
2. **text는 내 언어로** 쓴다(한국어든 영어든). **anchors는 글이 작성될 기본 언어의 어휘**로 쓴다 — 둘은 달라도 된다. anchors는 "이 표현이 본문에 실제로 나오면 이 주장이 다뤄졌다고 볼 단서"다. 표기 변형·동의어가 있으면 목록에 모두 적는다(하나만 나와도 인정).
3. **kind를 고른다**: `claim`(주장) · `rebuttal`(반론 대응) · `prediction`(예측) · `definition`(정의). 이 4개 밖의 값은 에러다.
4. **출처가 필요한 주장**(수치·연구 인용 등)은 `requires_source: true`로 표시한다.
5. **non_goals**에 다루지 않을 것을 적는다 — 에이전트의 범위 이탈을 막는다.
6. `tone`은 힌트일 뿐 게이트 대상이 아니다.

## 저장 위치 (규약)

번역본은 명세를 공유하므로 **기본 언어에만** 둔다. 짝짓기는 파일 이름 기준이다(front matter의 slug와 무관).

| 글 형태 | 글 위치 | insight.yaml 위치 |
|---|---|---|
| 번들 | `content/{기본lang}/{section}/{디렉토리}/index.md` | 같은 디렉토리의 `insight.yaml` |
| 플랫 | `content/{기본lang}/{section}/{이름}.md` | 같은 디렉토리의 `{이름}.insight.yaml` |

insight.yaml은 발행 산출물에 포함되지 않는다 — 생성된 hugo.toml의 `ignoreFiles`가 항상 제외한다.

## 대조 (REVIEW 보조)

```bash
abloq insight match content/en/tech/my-post.insight.yaml content/en/tech/my-post.md
```

본문(front matter 제외)을 NFC 정규화 + 케이스 폴딩한 뒤 anchors의 부분문자열 출현을 판정하고, `section`과 글의 실제 위치 일치를 검사한다. **미출현 claim 목록은 REVIEW 단계에 제시하는 보조 자료다** — 출현이 대응을 보장하지 않고, 미출현이 곧 누락도 아니다. 굴절·표기 변형은 anchors 동의어 목록으로 흡수한다(형태소 분석 없음). 번역 글 대조에는 쓰지 않는다(기본 언어 글 전용).

## 스키마

| 키 | 타입 | 필수 | 설명 |
|---|---|---|---|
| `topic` | string | | 글의 주제 한 줄 |
| `stance` | string | | 관점 한 줄 — 이 글이 어느 편에 서는가 |
| `audience` | string | | 누구에게 말하는 글인가 |
| `section` | string | ✅ | 글이 놓일 섹션 — 실위치와 일치해야 한다 |
| `tags` | [string] | | 태그 (taxonomy SSOT 어휘) |
| `claims[]` | list | ✅ (≥1) | 주장 목록 — 아래 필드 참조 |
| `claims[].id` | string | ✅ | 유니크 식별자 (kebab-case 권장) |
| `claims[].text` | string | ✅ | 주장 한 문장 — 사람 언어 |
| `claims[].kind` | string | ✅ | `claim` \| `rebuttal` \| `prediction` \| `definition` |
| `claims[].requires_source` | bool | | 출처 필수 여부 (기본 `false`) |
| `claims[].anchors` | [string] | | 본문 대응 확인용 핵심 어휘 — **글 언어**, 동의어 허용 목록. 비우면 match가 스크리닝하지 못한다(경고) |
| `non_goals` | [string] | | 다루지 않을 것 |
| `tone` | string | | 문체 힌트 (게이트 비대상) |

스키마에 없는 키는 에러다(strict). 검증 룰: `insight-claims-min` · `insight-claim-id-unique` · `insight-claim-kind` (에러), `insight-claim-anchors-empty` (경고).

## 예제 3종

- 기술 글: [`examples/insight-tech.yaml`](examples/insight-tech.yaml)
- 의견 글: [`examples/insight-opinion.yaml`](examples/insight-opinion.yaml)
- 정의 글: [`examples/insight-definition.yaml`](examples/insight-definition.yaml)
