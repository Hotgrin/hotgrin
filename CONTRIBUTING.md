# Contributing to SimpleScript

Thanks for your interest — SimpleScript is young and friendly contributions are
very welcome.

## Getting set up

You need **Go 1.22+** ([go.dev/dl](https://go.dev/dl/)). Then:

```bash
git clone https://github.com/hotgrin/simplescript
cd simplescript
go build ./...
go test ./...
```

## The shape of the project

SimpleScript is a transpiler: SimpleScript source → tokens → tree → Go → a real
program. Each stage is its own package under `internal/`:

- `lexer` — text to tokens
- `parser` — tokens to the tree (AST)
- `transpiler` — the tree to Go
- `loader` — resolves `use` libraries into one program
- `watcher` — the checker that reports provable problems

A new language feature usually touches the lexer (if it needs a keyword), the
parser (an AST node + a parse function), the transpiler (how it becomes Go), and
the watcher (so mistakes are caught kindly).

## Two house rules

1. **Every change ships with a test.** We keep the suite green.
2. **The Watcher never raises a false alarm.** If you add a rule, it must only
   fire on something that is genuinely, provably wrong. When in doubt, don't flag.

## Before you open a PR

```bash
go vet ./...
go test ./...
```

Keep code clear and beginner-readable over clever — that is the whole point of
SimpleScript. Dankie!
