// Command sslex is a tiny tool that shows the lexer at work: give it a
// SimpleScript file and it prints the token stream. With no file, it lexes a
// built-in demo that touches many language features at once.
//
//	go run ./cmd/sslex            # lex the built-in demo
//	go run ./cmd/sslex myapp.ss   # lex your own file
package main

import (
	"fmt"
	"os"

	"github.com/hotgrin/simplescript/internal/lexer"
)

const demo = `# A small taste of SimpleScript
say "Hello, world"

set cart total to 0
increase cart total by 199

describe Adriaan
    age is 56
    location is "Johannesburg"
end describe

action grade with name, mark
    if mark is at least 50
        give back name plus " passed"
    else
        give back name plus " must retry"
    end if
end action

say grade with "Adriaan", 82
say location of Adriaan
`

func main() {
	source := demo
	if len(os.Args) > 1 {
		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "could not read file:", err)
			os.Exit(1)
		}
		source = string(data)
	}

	for _, tok := range lexer.New(source).Tokenize() {
		// Skip the newline noise for a cleaner view, but show line numbers.
		if tok.Type == lexer.NEWLINE {
			fmt.Println("  --- line", tok.Line, "---")
			continue
		}
		fmt.Printf("  %-18s %q\n", tok.Type, tok.Literal)
	}
}
