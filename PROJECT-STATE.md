# Project state

The single source of truth for "where are we right now." Update this at the
end of any session that changes it — a stale README section is a bug; a
stale `PROJECT-STATE.md` is a much worse one, because it's the thing meant
to prevent us re-discovering our own history by accident.

This file answers: what's actually shipped (verified against the real
GitHub remote, not a chat summary), what's mid-flight, what's decided, and
what's next. If a chat and this file disagree, **this file wins** — it's
checked against source; a chat summary might describe a sandbox that never
got pushed.

Last verified against remote: **2026-07-21**, commit `f0ce4dc`, tag `v0.5.10`.
This update (Day Two, v0.5.11) prepared but not yet pushed as of this
writing — see Mid-flight below.

---

## Shipped, confirmed on the real remote

- **v0.5.10** is the latest tag, verified via fresh clone + tag-points-at-
  correct-commit check + full build/vet/test + a spot-check of the
  trickiest glossary snippet run on the actual published binary. **The
  living glossary** (`docs/glossary.md`) is live: 45 terms, comprehensive,
  every code example actually run before shipping.
- **v0.5.9**: Record prototypes now work correctly from any scope —
  actions, `try` blocks, `test` blocks, and combinations — a bug two
  earlier sessions worked on without it ever reaching the remote; this is
  the one that's actually fixed. **Day Zero and Day One**
  (`docs/day-zero.md`, `docs/day-one.md`) are both live and wired in,
  telling one continuous story (same tea example throughout).
- **gobug** (`github.com/Hotgrin/gobug`) — separate side project, CI fully
  fixed and confirmed: v0.2.1 release has three real attached binaries
  (Windows, macOS, Linux), verified by checking the actual release page,
  not just a green CI checkmark.
- **v0.5.6** and earlier: core language is stable: units of measure,
  `std/web`, built-in testing, the Watcher, bilingual (English/Afrikaans)
  errors, remote GitHub libraries, `use go` escape hatch.
- 27-lesson learn path (`examples/learn/`), 21-recipe cookbook, browser
  playground, AI prompt pack + `llms.txt` + `.cursorrules` +
  `copilot-instructions.md`.
- Four homepage showcase examples + the 283-line `site-report.hot` flagship
  program (commit `0d58b19`).
- Record-field-writes (`set price of order to 249`), std renames that
  removed reserved-word collisions (`starts with` → `has prefix`, etc).
- **All six audited engine bugs confirmed filed as real GitHub issues**
  (2026-07-21). Current state:
  - #1 CRITICAL — `at the same time` hangs. Open. Narrowing comment
    posted (2026-07-21, confirmed by AJ) — narrows it to calling **two or
    more different actions** in one block; calling the same action
    multiple times is fine. Bug itself still open.
  - #2 HIGH — `list of nothing` collapses list type to `any`. Open.
  - #3 — **Closed (2026-07-21, confirmed by AJ).** Was the
    record-instantiation bug; fixed and explained in v0.5.9.
  - #4 MEDIUM — `give back` inside `try` fails while Watcher says "All
    good." Open.
  - #5 MEDIUM — unknown record field passes Watcher, fails Go compile.
    Open.
  - #6 LOW — Watcher false-positive "unused" on variables used later in
    the same action/test (covers both originally-audited false alarms in
    one issue). Open.

## Mid-flight — needs a decision or re-verification, not assumed done

- **Day Two (v0.5.11)** — `docs/day-two.md`, continues the Day Zero/Day
  One story with the first real decision, reusing Day Zero's umbrella
  rule. All code verified, including every branch of the final `else if`
  example run with every combination of `true`/`false` and confirmed
  against the exact output the lesson claims. Wired into README,
  `getting-started.md`, `examples/learn/README.md`, and Day One's closing
  (as an optional continuation — Day One still also points to Lesson 01).
  **Prepared, not yet pushed to GitHub as of this writing.**

## Decided — house rules, don't relitigate

- **One source of truth for content.** Beginner-education copy lives in the
  repo (`docs/`), not duplicated into WordPress. hotgrin.com is a thin
  front door that *links to* the repo, never a second copy of it.
- **Day Zero is the canonical absolute-beginner entry point**, folding in
  the strongest idea from an earlier, independent draft called "Class 1:
  You Already Know How To Do This" (the "break it on purpose" exercise).
  That earlier draft is retired — don't resurrect it as a separate page.
- **Numbered comment system is now a formal teaching convention**, not
  just a style the Invoice Maker example happened to use: `[1]`–`[N]`
  section index at the top of a file, matching numbered dividers through
  the body. Next beginner lesson (working title "Class 2" / "Day One")
  teaches this explicitly, plus a "before you type anything" planning
  ritual and a comment-first-code-second workflow.
- **Every shipped change gets a version bump, a changelog entry, and a
  pushed tag.** No partial states on the remote.
- **Live verification only.** Nothing ships without being built, `go vet`,
  `go test`, and (for `.hot` code) actually run through the compiled
  binary — never eyeballed or assumed correct from reading source.
- **Check `git tag` before creating a new tag, every time.** A stale local
  tag silently pointing at the wrong commit shipped as v0.5.8's *first*
  attempt — `git tag vX.Y.Z` doesn't overwrite an existing tag, it just
  fails quietly enough to miss, and the wrong commit gets released. Always
  `git tag` first to check for a collision, and after pushing, verify via
  fresh clone that the tag actually points at the commit you meant.

## Beginner-education initiative — status

Sequence agreed: **Day Zero → Day One → living glossary → first
micro-lessons → AI Mentor**, with community-building ("Study Stoep") and
the hotgrin.com homepage redesign running alongside whenever there's room.
Day Zero, Day One, and the glossary are all shipped. Decided (2026-07-21):
first micro-lessons continue as a "Day Two, Day Three..." story rather
than standalone practical mini-projects — Day Two shipping now (see
Mid-flight above), Day Three (loops) and beyond to follow the same
pattern: reuse an earlier example where possible, live-verify every
branch, hand off to the matching numbered lesson at the end. Homepage
redesign was proposed but not started;
first micro-lessons and AI Mentor not started.

## Marketing / launch — status

dev.to article published, low traction (reported ~1 view, no external
distribution). Account restricted from r/learnprogramming and
r/ProgrammingLanguages for AI-assisted development disclosure — a
community-mood issue, not a verdict on hotgrin. **ZATech Slack post
confirmed live (2026-07-21)** — AJ posted it and is monitoring for
replies. Show HN not yet attempted. Current stance: grow the real thing
first (Day Zero, real beginner users), let institutions and wider
marketing follow evidence rather than chase them early.

## Chat hygiene

Six prior chats existed in this project as of 2026-07-21 covering: initial
build (v0.1–v0.5.4), a stale container-reset checkpoint (safe to delete),
a documentation/bug-fix audit + dev.to/Reddit fallout, an off-topic
portfolio-strategy detour (belongs in a different project, not here), the
Class-1/Day-Zero duplicate (retired, see above), and a compiler-bug session
that ended mid-verification. Going forward: **update this file instead of
relying on chat summaries to reconstruct state.** If a new session needs
history a search can't answer, that's a sign this file needs a better
entry, not a sign to go digging through old chats.

## Next up

1. Push Day Two (v0.5.11) — the one blocking step to close that loop.
2. Day Three (loops) — next in the Day-story sequence.
3. hotgrin.com homepage redesign (simple, plain-language nav, one button).
4. Watch for ZATech replies; respond as they come in.
