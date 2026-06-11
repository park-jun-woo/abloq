---
title: "robots.txt Is a Signal, Not a Lock — Thirty Years of the Robots Exclusion Protocol"
date: 2026-06-11
lastmod: 2026-06-11
tags: ["robots.txt", "GEO", "crawler"]
---

Every crawler that lands on your site makes the same first move: it asks for
`/robots.txt`. The Robots Exclusion Protocol (REP) is the convention behind that
move — a crawler fetches the robots.txt file at the site root and decides for
itself which URIs it may access. The decision happens entirely on the client
side. Your server serves a text file; the crawler interprets it.

## A 1994 convention that became a standard in 2022

The practice is far older than the standard. As the RFC itself credits, the
protocol was originally defined by Martijn Koster in 1994 for service owners
to control how crawlers access their sites. For almost three decades
robots.txt ran the web as a de facto convention with no normative
specification. Only in September 2022 did the IETF publish
[RFC 9309: Robots Exclusion Protocol](https://www.rfc-editor.org/rfc/rfc9309.html),
turning the 1994 practice into a Proposed Standard — and codifying details the
convention had left vague, such as precedence rules and the 500 KiB parsing
limit crawlers may apply.

## What REP is not: access control

RFC 9309 itself is blunt about the protocol's limits: the rules are
["not a form of access authorization"](https://www.rfc-editor.org/rfc/rfc9309.html).
A robots.txt file is a request, and compliance is voluntary. A well-behaved
crawler honors it; a hostile one reads it as a map of what you consider
sensitive. If something must not be fetched, enforcement belongs on the server
side — authentication, IP-level blocking, signed URLs — never in robots.txt.

This is the single most common operational mistake: treating robots.txt as a
lock. It is a signal. The lock is your server.

## The AI crawler era makes the signal strategic

What changed recently is who is asking. Alongside classic search indexers, an
AI crawler may be collecting training corpora, feeding a search-grounded
answer engine, or fetching a page live to answer one user's question. Each
category carries a different trade-off between visibility and control, which
is why robots.txt policy has become part of GEO (Generative Engine
Optimization) decision-making rather than a one-time ops chore: which agents
you admit determines whether generative engines can see — and cite — your
content at all.

The protocol gives you exactly one honest lever: declare, per user agent, what
a voluntary, well-behaved crawler should skip. Design your policy around that
honesty — signal openly, enforce server-side — and the thirty-year-old text
file keeps doing its job in the AI era.

## Sources

- [RFC 9309: Robots Exclusion Protocol](https://www.rfc-editor.org/rfc/rfc9309.html) — IETF, September 2022
- The Web Robots Pages (robotstxt.org) — the 1994 original "A Standard for Robot
  Exclusion"; the site answers HTTP 403 to automated fetchers, so it is listed
  without a hyperlink (see worklog)
