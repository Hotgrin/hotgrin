// Command hotparse shows the parser at work: it reads a hotgrin program
// and prints the abstract syntax tree (AST) as a readable S-expression.
//
//	go run ./cmd/hotparse            # parse the built-in demo
//	go run ./cmd/hotparse myapp.hot   # parse your own file
package main

import (
	"fmt"
	"os"

	"github.com/hotgrin/hotgrin/internal/lexer"
	"github.com/hotgrin/hotgrin/internal/parser"
)

const demo = `say "Hello, world"

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

put grade with "Adriaan", 82 into results

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

	tokens := lexer.New(source).Tokenize()
	prog, errs := parser.New(tokens).Parse()

	fmt.Println(prog.String())

	if len(errs) > 0 {
		fmt.Fprintln(os.Stderr, "\nparse problems:")
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, "  "+e.String())
		}
		os.Exit(1)
	}
}
