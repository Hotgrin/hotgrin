// Command sscheck runs the Watcher over a SimpleScript program and prints what
// it found. Use -af to see the messages in Afrikaans.
//
//	go run ./cmd/sscheck            # check the built-in demo (has mistakes)
//	go run ./cmd/sscheck app.ss     # check your own program
//	go run ./cmd/sscheck -af app.ss # messages in Afrikaans
package main

import (
	"fmt"
	"os"

	"github.com/hotgrin/simplescript/internal/lexer"
	"github.com/hotgrin/simplescript/internal/parser"
	"github.com/hotgrin/simplescript/internal/watcher"
)

// The demo deliberately contains several provable mistakes.
const demo = `action greet with name
    say "Hello " plus name
    give back name
    say "this line can never run"
end action

set total to 100
set unusedThing to 5

say greet with "AJ", "extra"

say discountt

set price to total divided by 0

if 5 is greater than 3
    say "always"
end if
`

func main() {
	lang := "en"
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "-af" {
		lang = "af"
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
	for _, e := range perrs {
		fmt.Printf("  parse  line %d: %s\n", e.Line, e.Message)
	}

	findings := watcher.New(prog).Check()
	if len(findings) == 0 {
		fmt.Println("Nothing to report — the Watcher found no provable problems.")
		return
	}

	for _, f := range findings {
		fmt.Printf("  %-10s line %d: %s\n", f.Severity, f.Line, f.Message(lang))
	}
}
