---
name: abloq
description: Operate and set up an agent-run blog. Use when working in a repo with a blog.yaml SSOT — scaffolding a blog (abloq init), generating derived files (hugo.toml/robots.txt/llms.txt/JSON-LD), running GEO/structure/evidence gates, or executing the writing/translation/refresh/evidence/cluster quests. Triggers on blog.yaml, insight.yaml, quests/queue, "abloq gate", citation/source verification, hreflang/freshness/cluster checks, and self-hosting the abloqd backend.
license: MIT
metadata:
  author: park-jun-woo
  version: "0.1.0"
---

# abloq — Agentic blog Quest CLI

abloq drives an agent-operable blog from a single `blog.yaml` SSOT. Humans decide insight;
agents do the prose (quests); a deterministic gate locks "done". Generation is probabilistic;
verification is deterministic.

## When to Use This Skill

- The repo has a `blog.yaml` at its root (or you are creating one with `abloq init`).
- You are scaffolding, generating derived files, or running gates on an abloq blog.
- You are executing a quest: writing, translation, refresh, evidence, or cluster curation.
- You are consuming `quests/queue/*.yaml`, fixing gate FAILs, or self-hosting `abloqd`.

## Do NOT Use This Skill

- A plain Hugo/Jekyll/Astro site with **no `blog.yaml`** — abloq commands assume the SSOT.
- General Markdown editing or a generic static site unrelated to abloq's gates.
- Calling external APIs directly from a quest — agents never touch Wayback/IndexNow/GSC by hand
  (the backend or `abloq archive` does, with env credentials).

## Install

```bash
go install github.com/park-jun-woo/abloq/cmd/abloq@latest
```

## Commands

`[dir]` defaults to `.` (the blog root holding `blog.yaml`). Exit 1 = diagnostics/violations.
Diagnostic format: `file:line [rule-id] message`. `--json` emits a JSON array.

| Command | Purpose |
|---|---|
| `abloq init <dir> --title --baseurl --author --languages --sections` | Scaffold a new blog (blog.yaml + template + CLAUDE.md), gate-clean. Non-interactive by default; `--interactive` to prompt |
| `abloq validate [dir]` | Validate blog.yaml schema and rules |
| `abloq generate [dir]` | Generate hugo.toml, robots.txt, llms.txt, jsonld.json from blog.yaml |
| `abloq check [dir]` | Check derived files against a fresh regeneration (exit 1 on drift) |
| `abloq gate [dir] [--rule <id>] [--offline] [--json]` | Run the 14 structure/evidence rules on articles. `--offline` skips citation-exists |
| `abloq claudemd [dir]` | Regenerate CLAUDE.md from blog.yaml |
| `abloq postbuild md [dir]` | Serve a clean `.md` beside every built article (run after `hugo`) |
| `abloq image og <slug> <title> [--brand --font --out]` | Generate a 1200x630 OG WebP |
| `abloq image convert <src> [--slug --max-width --out]` | Convert an image to WebP |
| `abloq scan {freshness\|evidence\|cluster} [dir]` | Detect quest candidates → write `quests/queue/` |
| `abloq ingest crawl --source <dir\|s3://...> [--repo]` | Aggregate AI-bot crawl hits from CloudFront logs (no DB) |
| `abloq ingest gsc [--site --days --repo]` | Fetch recent GSC Search Analytics rows (no DB) |
| `abloq sample citations --queries <yaml\|json> [--repo]` | Probe AI engines for own-domain citations (not gated) |
| `abloq report monthly --ym <YYYY-MM> [--source <logs>] [dir]` | Partial monthly visibility report (no DB) |
| `abloq archive <url>` | Submit a URL to Wayback/IndexNow/GSC (env credentials) |
| `abloq insight match <insight.yaml> <article>` | Screen insight claims against an article (REVIEW aid) |
| `abloq quest <writing\|translation\|refresh\|evidence\|cluster> <scan\|next\|submit\|status\|export\|rules>` | reins-gated agent quests |

> Credentials for `archive`/`ingest`/`sample` come from env only (`AWS_*`,
> `GSC_SA_JSON`/`GSC_SA_JSON_PATH`, IndexNow/Wayback keys) — never in args or blog.yaml.

## Workflow

**Daily loop:**

```bash
abloq validate .     # blog.yaml itself
abloq generate .     # regenerate derived files
abloq check .        # drift check (exit 1 on drift)
abloq gate .         # 14 rules on articles (exit 1 on violations)
```

**Quests** (each has `scan / next / submit --key <k> --in <file> / status / export / rules`):

- **writing**: `scan <insight.yaml>` → write article → `submit` with
  `{"article","worklog","review"}`. The REVIEW record must be authored by a **separate context**
  (a different session) — a `reviewer:` id distinct from the writing context, plus a disposition
  for every match-missing claim. The writing agent may NOT review its own article.
- **translation**: `scan <default-lang article.md>` seeds one item per non-default language;
  `submit {"article"}`. Preserve heading levels, paragraph blocks, image paths, code blocks,
  external URLs (verbatim); rewrite internal links to the target-language prefix; mirror the
  origin's `date`/`lastmod`.
- **refresh / evidence / cluster** (queue consumers): `scan [instance-dir]` reads
  `quests/queue/`. Follow the **commit order** strictly (below).

**Queue consumer commit order (do NOT reorder — reordering is blocked):**

1. Edit the article in the **working tree** (do not commit yet).
2. `submit` → PASS. The gate compares the **dirty working tree vs git HEAD**. Committing the
   article first makes tree==HEAD, voiding the baseline rules — forbidden.
3. Commit the article edit (only after PASS).
4. If `lastmod` advanced: re-sync all languages via the translation quest, then commit.
5. Commit the queue-file deletion (consumption signal) **last**.

## Key Concepts

- **SSOT**: `blog.yaml` declares site/languages/sections/structure/geo/deploy. All derived files
  and gate parameters come from it. No article bypasses the gate unless blog.yaml changes.
- **Gate = violation detectors**: a rule fires with a `Fact{Where, Expected, Actual}`. Fix that
  Fact, resubmit. `MaxTries=3` → a 4th FAIL locks the item to DONE permanently.
- **Authority asymmetry**: only the machine (L1) locks PASS; the agent (L2) does REVIEW only;
  the human (L3) does the rest.
- **Insight**: a human-authored `insight.yaml` (one per article, default language only) is the
  writing gate's REVIEW basis. See `docs/insight-spec.md`.

## Common Errors and Fixes

| FAIL rule | Cause | Fix |
|---|---|---|
| `front-matter-schema` | missing title/date/lastmod/tags | fill required front matter fields |
| `image-first` / `image-attribution` | body starts with text / no attribution line | put the main image first, italic source line right below |
| `section-order` / `heading-canonical` | sections out of `structure.order` / wrong heading level | reorder per blog.yaml; use `##` for recognized sections |
| `min-sources` | sources section below `geo.min_sources` | add source entries |
| `numeric-claim-sourced` | a numeric claim added since HEAD has no source | add an inline source link in the same paragraph |
| `citation-exists` | URL not 200 / title mismatch | replace with a real URL whose title overlaps the anchor text (or `--offline` to skip) |
| `honest-lastmod` | lastmod advanced without a meaningful body diff or queue membership | make a real `>= geo.min_meaningful_diff` change, or do not touch lastmod |
| `front-matter-intact` | changed a field other than lastmod | revert all front matter except lastmod |
| `slug-consistency` / `hreflang-complete` | slug differs across languages / missing hreflang | use one valid slug across all languages; build all-language alternates |
| `review-record` | missing REVIEW record / no `reviewer:` line / undisposed claim | add a separate-context review with a disposition per match-missing claim |
| `translation-parity` | structure drift from origin | match headings, paragraph blocks, images, code, links, date/lastmod |
| `lastmod-advance` (refresh) | empty refresh (lastmod not advanced) | make real changes and advance lastmod |
| `claim-preserved` (refresh) | numeric claim count dropped | replace stale figures, do not delete claims |
| `claims-resolved` / `rot-resolved` (evidence) | queued unsourced claim still unsourced / rot URL still cited | add a source in-paragraph; replace dead URLs |
| `claim-scope` (evidence) | a claim outside the queue payload changed | leave non-queued claims byte-identical; do not re-wrap claim lines |
| `cluster-resolved` (cluster) | queued cluster violation still present on rescan | fix tags / add internal links from payload candidates |
| `queue-scope` (queue quests) | edited a file outside the allowed set, or touched the queue file pre-gate | restrict edits to the target article + its translations + insight sidecar (+ cluster candidates); delete the queue file only in the final commit |

## Conventions (cheese defenses)

- **No self-REVIEW** — the writing context cannot author its own REVIEW (`review-record`).
- **No out-of-scope edits** — never touch blog.yaml, layouts, or unrelated articles
  (`queue-scope`).
- **No fake lastmod** — body-less lastmod bumps fail `honest-lastmod`; translations mirror the
  origin's lastmod.
- **No fabricated sources** — `citation-exists` verifies HTTP 200 + title overlap.
- **No claim deletion / re-wording to dodge detection** — `claim-preserved`, `claim-scope`,
  `numeric-claim-sourced` catch it.
- **No gate voiding** — never commit the article before `submit`; the gate needs a dirty tree vs
  HEAD baseline.
- **Agents never call external APIs directly** — archive/index side-effects are backend receipts
  (or `abloq archive`).

## Full Documentation

- **`MANUAL.md`** — full operating manual (CLI reference, quest procedures, gate diagnostics,
  cheese defenses, backend integration).
- `docs/blog-yaml.md` — blog.yaml schema reference.
- `docs/insight-spec.md` — insight.yaml authoring guide.
- `docs/operations.md` — abloqd self-host operations (cron profiles, failure runbooks).
- `README.md` — concepts and motivation.
- Gate engine: [reins](https://github.com/park-jun-woo/reins) (MIT).
