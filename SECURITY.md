# Security Policy

## Reporting a vulnerability

If you find a security issue, please report it privately by opening a GitHub
security advisory on this repository, or by contacting the maintainer directly,
rather than filing a public issue. We will acknowledge your report and work with
you on a fix and disclosure timeline.

## What SimpleScript does on your machine

SimpleScript is a transpiler and command-line tool. Being honest about what it
does helps you reason about its security posture:

- It reads your `.ss` source files and any local libraries you `use`.
- To run or build, it writes generated Go to a temporary directory and shells
  out to your installed `go` toolchain. The temporary directory is removed
  afterwards.
- The core language has no networking and reads no secrets. It does not fetch
  remote libraries (that feature is not built yet).

Because `run`/`build`/`test` invoke the Go compiler on generated code, only run
SimpleScript programs you trust, exactly as you would with any source you
compile. Libraries you `use` are compiled as part of your program.
