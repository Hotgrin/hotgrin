# Changelog

All notable changes to SimpleScript are recorded here. This project follows
[Semantic Versioning](https://semver.org/) loosely while it is pre-1.0.

## [0.1.0] - 2026-06-30

The first public release: a clean rebuild from a full specification. The whole
pipeline works end to end — near-English source compiles to a real native
program (or a Windows `.exe`).

### Language
- Variables (`set`), output (`say`), arithmetic, comparisons, and boolean logic.
- `plus` concatenates when either side is text; `divided by` always gives a decimal.
- Conditionals (`if` / `else`) and loops (`repeat N times`, `repeat while`,
  `repeat for each`).
- Actions with inferred parameter and return types (`action ... give back ...`).
- Records via `describe ... end describe`, with `field of record` access.
- Lists with `list of ...`, `put X into list`, `item N of list`, `count of list`.
- Concurrency: `at the same time` (with safe collection) and `start` (background).
- Error handling: `give back problem`, `try` / `if it fails`, and `the problem`.
- First-class tests: `test "..." ... end test` with `expect` assertions.
- Command-line inputs: `input name as text default "..."`, with an auto `--help`.
- Local libraries: `use "path"` merges another file's actions (whole-program).
- Unicode identifiers and strings; English/Afrikaans messages.

### Tooling
- `simplescript` command: `run`, `test`, `build` (`--windows` for a `.exe`),
  `check` (`--af` for Afrikaans), `reveal`, `help`, `version`.
- The **Watcher**: a deterministic checker that reports only provable problems
  (unknown names, wrong argument counts, divide-by-zero, unreachable code,
  duplicate definitions, unhandled fallible calls, and more) — with the iron
  rule that a flag always means a real issue.
- Developer tools: `sslex`, `ssparse`, `ssrun`, `sscheck`.

### Notes
- Requires Go 1.22+ installed; the toolchain is used under the hood.
- Known limitations: tiny standard library, no remote libraries yet, no
  interactive `ask` yet, cross-file line numbers in some messages are per-file.
