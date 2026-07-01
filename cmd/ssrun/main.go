// Command ssrun is the whole pipeline in one tool: it reads a SimpleScript
// program, transpiles it to Go, and then runs it. With -go it prints the
// generated Go instead ("reveal the Go").
//
//	go run ./cmd/ssrun            # run the built-in demo
//	go run ./cmd/ssrun app.ss     # run your program
//	go run ./cmd/ssrun -go app.ss # show the Go it generates
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hotgrin/simplescript/internal/lexer"
	"github.com/hotgrin/simplescript/internal/parser"
	"github.com/hotgrin/simplescript/internal/transpiler"
)

const demo = `say "Hello from SimpleScript"

set price to 199
set quantity to 3
set total to price times quantity
say total

describe Adriaan
    age is 56
    location is "Johannesburg"
end describe

say location of Adriaan

action grade with name, mark
    if mark is at least 50
        give back name plus " passed with " plus mark
    else
        give back name plus " must retry"
    end if
end action

say grade with "Adriaan", 82
say grade with "Johan", 47

set scores to list of 90, 85, 100
put 75 into scores
say count of scores

repeat 3 times
    say "SimpleScript works"
end repeat

repeat for each s in scores
    say s
end repeat
`

func main() {
	showGo := false
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "-go" {
		showGo = true
		args = args[1:]
	}

	source := demo
	if len(args) > 0 {
		data, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "could not read file:", err)
			os.Exit(1)
		}
		source = string(data)
	}

	tokens := lexer.New(source).Tokenize()
	prog, perrs := parser.New(tokens).Parse()
	if len(perrs) > 0 {
		fmt.Fprintln(os.Stderr, "parse problems:")
		for _, e := range perrs {
			fmt.Fprintln(os.Stderr, "  "+e.String())
		}
		os.Exit(1)
	}

	goSrc, _, terrs := transpiler.New(prog).Transpile()
	if len(terrs) > 0 {
		fmt.Fprintln(os.Stderr, "transpile notes:")
		for _, e := range terrs {
			fmt.Fprintln(os.Stderr, "  "+e)
		}
	}

	if showGo {
		fmt.Print(goSrc)
		return
	}

	// Write the Go to a temp module and run it.
	dir, err := os.MkdirTemp("", "ssrun")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

	if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte(goSrc), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module ssrun\n\ngo 1.22\n"), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cmd := exec.Command("go", "run", ".")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
}
