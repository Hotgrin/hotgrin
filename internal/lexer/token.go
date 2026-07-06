// Package lexer turns hotgrin source text into a stream of tokens.
//
// This file defines the tokens themselves: the small, fixed vocabulary of
// "connector words" that act as the grammar's rails, plus the literal and
// structural tokens (names, numbers, text, punctuation, newlines).
package lexer

import "fmt"

// TokenType identifies what kind of thing a token is.
type TokenType int

const (
	// Structural / literal tokens
	ILLEGAL TokenType = iota // something the lexer did not understand
	EOF                      // end of the source
	NEWLINE                  // end of a line (hotgrin is line-oriented)
	IDENT                    // a name, which may contain spaces (e.g. "cart total")
	NUMBER                   // 42, 3.14, -7
	STRING                   // "in quotes" or 'in quotes'

	// Punctuation (the little we allow)
	LPAREN // (
	RPAREN // )
	COMMA  // ,

	// Value literals
	TRUE    // true / yes
	FALSE   // false / no
	NOTHING // nothing

	// Statement keywords
	SET
	GIVE_BACK
	SAY
	IF
	IF_IT_FAILS
	ELSE
	END
	REPEAT
	WHILE
	FOR_EACH
	IN
	ACTION
	USE
	TRY
	START
	DO
	AT_THE_SAME_TIME
	WAIT_FOR_ALL
	STOP_REPEATING
	SKIP_TO_NEXT
	TEST
	EXPECT
	INPUT
	ASK
	STOP_WITH_ERROR
	DESCRIBE
	INCREASE
	DECREASE
	PUT

	// Connectors (the rails)
	TO
	OF
	WITH
	AND
	OR
	INTO
	FROM
	BY
	AS
	DEFAULT

	// Comparison and math words
	PLUS
	MINUS
	TIMES
	DIVIDED_BY
	IS
	IS_NOT
	IS_GREATER_THAN
	IS_LESS_THAN
	IS_AT_LEAST
	IS_AT_MOST
	CONTAINS
	TO_BE
	TO_BE_AT_LEAST
	TO_BE_AT_MOST
	TO_BE_LESS_THAN
	TO_BE_GREATER_THAN
	ROUNDED_TO
)

// Token is a single lexical unit, with the line it came from so that every
// later stage can point error messages back at the user's hotgrin line.
type Token struct {
	Type    TokenType
	Literal string // the original text (names keep their casing and script)
	Line    int
}

func (t Token) String() string {
	return fmt.Sprintf("%s(%q)@%d", t.Type, t.Literal, t.Line)
}

// tokenNames gives each token type a readable name for debugging and tests.
var tokenNames = map[TokenType]string{
	ILLEGAL: "ILLEGAL", EOF: "EOF", NEWLINE: "NEWLINE", IDENT: "IDENT",
	NUMBER: "NUMBER", STRING: "STRING", LPAREN: "LPAREN", RPAREN: "RPAREN",
	COMMA: "COMMA", TRUE: "TRUE", FALSE: "FALSE", NOTHING: "NOTHING",
	SET: "SET", GIVE_BACK: "GIVE_BACK", SAY: "SAY", IF: "IF",
	IF_IT_FAILS: "IF_IT_FAILS", ELSE: "ELSE", END: "END", REPEAT: "REPEAT",
	WHILE: "WHILE", FOR_EACH: "FOR_EACH", IN: "IN", ACTION: "ACTION",
	USE: "USE", TRY: "TRY", START: "START", DO: "DO",
	AT_THE_SAME_TIME: "AT_THE_SAME_TIME", WAIT_FOR_ALL: "WAIT_FOR_ALL",
	STOP_REPEATING: "STOP_REPEATING", SKIP_TO_NEXT: "SKIP_TO_NEXT",
	TEST: "TEST", EXPECT: "EXPECT", INPUT: "INPUT", ASK: "ASK",
	STOP_WITH_ERROR: "STOP_WITH_ERROR", DESCRIBE: "DESCRIBE",
	INCREASE: "INCREASE", DECREASE: "DECREASE", PUT: "PUT",
	TO: "TO", OF: "OF", WITH: "WITH", AND: "AND", OR: "OR", INTO: "INTO",
	FROM: "FROM", BY: "BY", AS: "AS", DEFAULT: "DEFAULT",
	PLUS: "PLUS", MINUS: "MINUS", TIMES: "TIMES", DIVIDED_BY: "DIVIDED_BY",
	IS: "IS", IS_NOT: "IS_NOT", IS_GREATER_THAN: "IS_GREATER_THAN",
	IS_LESS_THAN: "IS_LESS_THAN", IS_AT_LEAST: "IS_AT_LEAST",
	IS_AT_MOST: "IS_AT_MOST", CONTAINS: "CONTAINS", TO_BE: "TO_BE",
	TO_BE_AT_LEAST: "TO_BE_AT_LEAST", TO_BE_AT_MOST: "TO_BE_AT_MOST",
	TO_BE_LESS_THAN: "TO_BE_LESS_THAN", TO_BE_GREATER_THAN: "TO_BE_GREATER_THAN",
	ROUNDED_TO: "ROUNDED_TO",
}

func (tt TokenType) String() string {
	if name, ok := tokenNames[tt]; ok {
		return name
	}
	return fmt.Sprintf("TokenType(%d)", int(tt))
}

// reserved maps every canonical connector word or phrase to its token type.
// Keys are lowercase and single-space-joined. yes/no alias true/false.
//
// The lexer matches these on WHOLE-WORD boundaries and always prefers the
// LONGEST phrase ("is greater than" beats "is"; "if it fails" beats "if"),
// which is what keeps multi-word names and multi-word keywords unambiguous.
var reserved = map[string]TokenType{
	// values
	"true": TRUE, "yes": TRUE, "false": FALSE, "no": FALSE, "nothing": NOTHING,
	// statement keywords
	"set": SET, "give back": GIVE_BACK, "say": SAY,
	"if it fails": IF_IT_FAILS, "if": IF, "else": ELSE, "end": END,
	"repeat": REPEAT, "while": WHILE, "for each": FOR_EACH, "in": IN,
	"action": ACTION, "use": USE, "try": TRY, "start": START, "do": DO,
	"at the same time": AT_THE_SAME_TIME, "wait for all": WAIT_FOR_ALL,
	"stop repeating": STOP_REPEATING, "skip to next": SKIP_TO_NEXT,
	"test": TEST, "expect": EXPECT, "input": INPUT, "ask": ASK,
	"stop with error": STOP_WITH_ERROR, "describe": DESCRIBE,
	"increase": INCREASE, "decrease": DECREASE, "put": PUT,
	// connectors
	"to": TO, "of": OF, "with": WITH, "and": AND, "or": OR, "into": INTO,
	"from": FROM, "by": BY, "as": AS, "default": DEFAULT,
	// comparison and math
	"plus": PLUS, "minus": MINUS, "times": TIMES, "divided by": DIVIDED_BY,
	"is not": IS_NOT, "is greater than": IS_GREATER_THAN,
	"is less than": IS_LESS_THAN, "is at least": IS_AT_LEAST,
	"is at most": IS_AT_MOST, "is": IS, "contains": CONTAINS,
	"to be at least": TO_BE_AT_LEAST, "to be at most": TO_BE_AT_MOST,
	"to be less than": TO_BE_LESS_THAN, "to be greater than": TO_BE_GREATER_THAN,
	"rounded to": ROUNDED_TO,
	"to be":      TO_BE,
}

// maxPhraseWords is the longest reserved phrase, in words
// ("at the same time", "to be greater than" = 4).
const maxPhraseWords = 4
