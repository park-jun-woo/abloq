---
title: "Ratchet Pattern — How to Make an Agent Finish the Job"
date: 2026-05-15T14:00:00+09:00
lastmod: 2026-05-27T12:00:00+09:00
image: "/images/ratchet-pattern.webp"
tags: ["Reins Engineering", "ratchet code", "AI", "agent", "Ratchet Pattern", "Symbolic Feedback Loop", "developer tools"]
summary: "I asked an AI agent to write tests for 527 functions. It stopped at 40 and declared 'done.' The Ratchet Pattern forces completion by delegating the done/not-done decision to a mechanical verifier — so the agent keeps going until the machine says stop."
---

![Ratchet Pattern](/images/ratchet-pattern.webp)
*Image: AI generated*

If your AI agent says "all done" but the work is not actually finished, if the agent keeps skipping hard tasks, if you want the machine -- not a human -- to judge whether something is complete -- this post explains the structure.

## "All Done"

I asked an AI agent to write tests for 527 functions. The agent finished its work and reported back.

"Done."

Functions that actually got tests: 40.

It wasn't lying. After 40 functions, it decided "that's enough." When it hit a difficult function, it skipped it. After a few more, it concluded "the rest follow a similar pattern, so we're good."

LLMs are great at generation. But they cannot be trusted to judge whether they are finished. Huang et al. experimentally demonstrated that LLM self-correction without external feedback can actually degrade performance[1].

---

## The Ratchet

A ratchet wrench has teeth that catch in only one direction. Turn it and it moves forward. Let go and it stops — but it never moves backward.

The Ratchet Pattern applies this mechanism to agent control. Verification code written with this pattern is called ratchet code — code that never allows regression below a previously passed verification level.

```
Item 1: mechanical verification → PASS → next
Item 2: mechanical verification → FAIL → retry (with feedback)
Item 2: mechanical verification → PASS → next
...
Item N: PASS → complete. Stop.
```

Three rules:
- Show only one item at a time.
- An item must pass before the next one opens.
- When all items pass, stop.

Implement these rules as a CLI, and the agent only needs to know one command: `next`. The machine decides the rest.

---

## The Agent That Stopped at 40 vs. the Ratchet That Finished 527

Same model. Same project. Same 527 functions.

```
Autonomous agent:  40 / 527  (7.6%)  — agent declared "done"
Ratchet CLI:      527 / 527 (100%)  — machine declared "487 remaining"
```

The difference is not model performance. It is **who decides when it's over**.

With an autonomous agent, the LLM decides when to stop. LLMs are optimistic. After 40, it feels like enough. In Cemri et al.'s trace analysis of 1,600 agent runs, premature termination — declaring the goal achieved before it actually was — accounted for 6.2% of all failure modes[2]. With a ratchet, the machine decides when to stop. The machine doesn't feel. It declares "not yet" until the remaining count hits zero.

---

## One-Sentence Definition

> Place a probabilistic agent inside a deterministic state machine.

| Role | Owner |
|------|-------|
| Generation | LLM |
| Judgment | verifier |
| Progress control | ratchet |

Many systems hand generation, judgment, and termination decisions all to the LLM. The Ratchet separates them.

---

## Five Principles

**1. The termination condition is mechanical**

pass/fail. Not "looks good." If `go test` passes, it's a PASS. If coverage hits 100%, it's a PASS. There is no room for subjective judgment.

**2. A PASS is immutable**

Once an item passes, it never reopens. No rollback. The remaining work count decreases monotonically.

```
remaining_work(t+1) ≤ remaining_work(t)
```

What you build today doesn't get torn apart tomorrow. Forward only. This is the fundamental difference from a "24-hour agent." An agent running without a termination condition adds an abstraction today, removes it tomorrow, and adds it back the day after. The ratchet does not permit that kind of oscillation.

**3. The LLM only generates**

Generate code, write tests, propose fixes — that is the LLM's role. What to fix, whether it passed, what comes next, whether it's done — the machine decides all of that. The LLM is not a planner; it is a constrained generator.

**4. Strip the agent of its right to declare completion**

If the LLM says "done," it stops at 40. If the machine says "done," it stops at 527. The entire reason the ratchet exists is captured in this single line.

**5. The verifier must be deterministic**

Not everything qualifies as a verifier.

| Can be a verifier | Cannot be a verifier |
|---|---|
| `go test` | "looks cleaner" |
| coverage measurement | "seems better" |
| AST validation | "more scalable" |
| schema diff | "clean architecture" |

A verifier must satisfy four conditions: deterministic, machine-checkable, resumable, localized feedback. If these are not met, the ratchet's teeth have nothing to catch on.

---

## Feedback as a Gradient Signal

If the ratchet only returns "pass/fail," the LLM corrects without direction. The more specific the feedback, the more precise the LLM's corrections become.

```
Weak feedback:    "test failed"             → LLM corrects without direction
Medium feedback:  "coverage 65%"            → LLM roughly reinforces
Strong feedback:  "line 41, 44, 70 uncovered" → LLM covers exactly those branches
```

Numbers verified in a real project:

```
Without feedback:  stuck at 60-70% coverage
With feedback:     100% achieved (for reachable functions)
```

Same model. A single line — "line 41 not covered" — acts as a gradient signal. Chen et al.'s Self-Debug research proved that iterative debugging by feeding compiler/runtime error messages back to the model dramatically improves code accuracy[3].

As feedback resolution increases, the LLM's correction accuracy rises, loop iterations decrease, and costs drop.

---

## Agents Die. Progress Survives.

Agents inevitably crash. Token limits, network errors, session disconnects. If the ratchet persists progress to storage, the next agent picks up where the last one left off.

```
Agent A: processes functions 1-200 → dies
Agent B: next → continues from 201
Agent C: next → continues from 401
```

Agents are disposable. Progress accumulates.

---

## Swap the Verifier, Get a Different Tool

The ratchet is not tied to any specific verifier. Change the verifier and you get a different tool.

| Ratchet + Verifier | Use case |
|---|---|
| Ratchet + `go test` + coverage | Per-function test generation |
| Ratchet + structural rule validator | Code structure cleanup |
| Ratchet + hurl pass/fail | API endpoint verification |
| Ratchet + spec cross-validation | SSOT consistency |
| Ratchet + Toulmin verdict | User-defined rule enforcement |

One pattern. The verifier determines the domain.

---

## Questions

How many items did your agent complete before saying "all done"?

Was it truly done?

Who decided "done" — the agent, or the machine?

---

## Related Posts

- [Model IQ Matters Less Than Feedback Topology](/opinion/feedback-topology/) — The theoretical background of the Ratchet Pattern. Why feedback structure matters more than model performance.
- [tsma](https://github.com/park-jun-woo/tsma) — A Go implementation of the Ratchet Pattern. 527 functions, zero TODO.
- [filefunc](https://github.com/park-jun-woo/filefunc) — A code structure implementation of the Ratchet Pattern. typer refactored, all 1155 tests pass.

## References

[1] Jie Huang, Xinyun Chen, Swaroop Mishra, Huaixiu Steven Zheng, Adams Wei Yu, Xinying Song, Denny Zhou. "Large Language Models Cannot Self-Correct Reasoning Yet." ICLR 2024. [arXiv:2310.01798](https://arxiv.org/abs/2310.01798)

[2] Mert Cemri, Melissa Z. Pan, Shuyi Yang, Lakshya A. Agrawal, Tanay Chopra, Aditya Tiwari, Kurt Keutzer, Aditya Parameswaran, et al. "Why Do Multi-Agent LLM Systems Fail?" NeurIPS 2025 Datasets and Benchmarks Track. [arXiv:2503.13657](https://arxiv.org/abs/2503.13657)

[3] Xinyun Chen, Maxwell Lin, Nathanael Scharli, Denny Zhou. "Teaching Large Language Models to Self-Debug." ICLR 2024. [arXiv:2304.05128](https://arxiv.org/abs/2304.05128)

[4] Noah Shinn, Federico Cassano, Ashwin Gopinath, Karthik Narasimhan, Shunyu Yao. "Reflexion: Language Agents with Verbal Reinforcement Learning." NeurIPS 2023. [arXiv:2303.11366](https://arxiv.org/abs/2303.11366)

[5] Aman Madaan, Niket Tandon, Prakhar Gupta, Skyler Hallinan, Luyu Gao, Sarah Wiegreffe, Uri Alon, Nouha Dziri, Shrimai Prabhumoye, Yiming Yang, et al. "Self-Refine: Iterative Refinement with Self-Feedback." NeurIPS 2023. [arXiv:2303.17651](https://arxiv.org/abs/2303.17651)

[6] Yujia Li, David Choi, Junyoung Chung, Nate Kushman, Julian Schrittwieser, Remi Leblond, Tom Eccles, et al. "Competition-Level Code Generation with AlphaCode." Science 378(6624): 1092-1097, 2022. [DOI:10.1126/science.abq1158](https://www.science.org/doi/10.1126/science.abq1158)

[7] Carlos E. Jimenez, John Yang, Alexander Wettig, Shunyu Yao, Kexin Pei, Ofir Press, Karthik R. Narasimhan. "SWE-bench: Can Language Models Resolve Real-World GitHub Issues?" ICLR 2024. [arXiv:2310.06770](https://arxiv.org/abs/2310.06770)

---

## Changelog

- 2026-05-15: Initial release
