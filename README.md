# SimpleScript

> A general-purpose programming language that reads like plain English, with the
> least punctuation that still keeps it clear — and compiles to a real program.

**Status: v0.1 — early alpha.** The language works end to end and is well tested,
but it is young: the standard library is tiny, error messages are still growing,
and some constructs (interactive `ask`, remote libraries) are not built yet.
Expect rough edges, and please file issues. SimpleScript needs **Go 1.22+**
installed (from [go.dev/dl](https://go.dev/dl/)); it uses the Go toolchain to
compile and run programs.

This is the **clean rebuild**, started from a fully-designed specification. The
**core pipeline now works end to end** — near-English goes in, a real program
comes out:

- the **lexer** — reads source text and breaks it into *tokens*;
- the **parser** — turns those tokens into the *tree* (AST) of the program;
- the **transpiler** — walks the tree and writes the equivalent **Go**, which the
  Go toolchain compiles into a real native executable (including a Windows `.exe`);
- the **watcher** — the always-on checker that reports *provable* problems in plain
  English (or Afrikaans), with the iron rule: if it flags something, it is real.

## Where things are

```
simplescript/
├── go.mod
├── internal/
│   ├── lexer/         # source text -> tokens
│   ├── ast/           # the tree node types
│   ├── parser/        # tokens -> tree
│   ├── transpiler/    # tree -> Go source
│   ├── loader/        # resolves "use" libraries into one program
│   └── watcher/       # tree -> findings (the checker)
└── cmd/
    ├── simplescript/  # the friendly command (run / test / build / check / reveal)
    ├── sslex/         # prints the tokens for a program
    ├── ssparse/       # prints the tree (AST) for a program
    ├── ssrun/         # transpiles + runs a program (use -go to reveal the Go)
    └── sscheck/       # runs the Watcher (use -af for Afrikaans messages)
```

Example programs live in `examples/`.

Example programs live in `examples/`.

## Try it (the friendly way)

Build the one command, then use it — no need to know Go is underneath:

```bash
go build -o simplescript ./cmd/simplescript

./simplescript run     examples/hello.ss     # run a program
./simplescript run     examples/greet.ss --name AJ   # pass inputs to it
./simplescript test    examples/calc.ss      # run a program's tests
./simplescript check   examples/hello.ss     # check it for problems
./simplescript check --af examples/hello.ss  # ...in Afrikaans
./simplescript build   examples/hello.ss     # make a standalone program
./simplescript build --windows examples/hello.ss   # make a Windows .exe
./simplescript reveal  examples/hello.ss     # show the Go it becomes
./simplescript help
```

`run` and `build` quietly run the Watcher first, so you get friendly SimpleScript
messages instead of raw Go errors.

## Under the hood (developer tools)

```bash
go run ./cmd/sslex     <file.ss>   # show the tokens
go run ./cmd/ssparse   <file.ss>   # show the tree (AST)
go run ./cmd/ssrun -go <file.ss>   # show the generated Go
go test ./...                      # run all tests (61 and counting)
```

## Tests, written in English

SimpleScript treats testing as part of the language, not an add-on:

```
action add with a, b
    give back a plus b
end action

test "addition works"
    expect add with 2, 3 to be 5
end test
```

`simplescript test` compiles and runs these for you. Assertions include
`to be`, `to be at least`, `to be at most`, `to be less than`,
`to be greater than`, and `contains`.

## Handling things that can fail

An action signals failure with `give back problem`, and callers handle it with
`try` / `if it fails`:

```
action safe divide with a, b
    if b is 0
        give back problem "cannot divide by zero"
    end if
    give back a divided by b
end action

try
    set answer to safe divide with 10, 0
    say answer
if it fails
    say "Could not do that: " plus the problem
end try
```

Inside `if it fails`, `the problem` holds the message. The Watcher makes sure a
call that can fail is always wrapped in a `try`, so failures are never silently
dropped.

## Inputs and libraries

A program can take command-line inputs, which also gives it a `--help`:

```
input name as text default "world"
say "Hello, " plus name
```

```bash
simplescript run greet.ss --name AJ
```

And it can reuse code from other SimpleScript files:

```
use "lib/textutils"
say greet with "AJ"
```

The library's actions are compiled together with your program, and the Watcher
checks the imported code too. Library paths are local for now (remote/GitHub
fetching is on the roadmap).

## What's next

The core spine, concurrency, the Watcher, first-class testing, error handling,
command-line inputs, local libraries, and the friendly `simplescript` command
all work. On the roadmap: remote libraries, interactive `ask` prompts, and the
optional Assessor and AI Mentor layers.

### Make a real executable

```bash
go run ./cmd/ssrun -go myapp.ss > main.go
printf 'module myapp\n\ngo 1.22\n' > go.mod
go build -o myapp                       # native executable
GOOS=windows GOARCH=amd64 go build -o myapp.exe   # a Windows .exe
```

## How the lexer works (the one clever idea)

SimpleScript lets a **name contain spaces** — `cart total` is one name, not two.
That only works because a small, fixed set of **connector words** act as rails:
`to`, `of`, `with`, `is`, `and`, and friends. A name is simply every word *up to
the next rail*.

So `set cart total to 0` reads as:

```
SET · IDENT("cart total") · TO · NUMBER("0")
```

The lexer always prefers the **longest** reserved phrase, so `is greater than`
beats `is`, and `give back` is a single keyword. Names keep their original casing
and may use letters from **any language** (`prénom`, `имя`, `名前`) — SimpleScript
is international from the very first stage.

## Licence

MIT — see [LICENSE](LICENSE).
