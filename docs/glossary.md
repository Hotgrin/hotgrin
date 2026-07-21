# Glossary

Every word hotgrin uses for something, explained in plain language, with
a tiny working example. This page doesn't teach you to code — [Day
Zero](day-zero.md) and [Day One](day-one.md) do that. This page is here
for the moment you're reading a program, hit a word you don't recognize,
and just want a quick, honest answer before you carry on.

Jargon is never the starting point on this page — every entry leads with
what the word means in plain English, then shows the code. If you want
the precise, exhaustive technical rules instead of the plain-English
version, the [language reference](language-reference.md) is the
authoritative source; this page is its friendlier front door.

**Jump to:** [A](#action) · [C](#comment) · [D](#decrease) · [E](#error-handling) · [F](#field) · [G](#give-back) · [I](#if--else-if--else) · [L](#list) · [O](#of) · [P](#parameter) · [R](#record) · [S](#say) · [T](#test) · [U](#unit-of-measure) · [V](#variable) · [W](#watcher) · [W](#with)

---

### action

A named, reusable set of steps — the same idea as a recipe you've taught
someone so well they can follow it on their own, whenever you say its
name. Other languages call this a "function."

```
action make tea with cups
    say "Boil the kettle"
    say "Pour " plus cups plus " cups"
end action

make tea with 2
```

See also: [parameter](#parameter), [give back](#give-back).

### and / or

Combine two true-or-false conditions. `and` needs both sides true; `or`
needs at least one side true.

```
if raining and cold
    say "Take a coat and an umbrella"
end if
```

### ask ... into

Prints a question and waits for the person using the program to type an
answer, which always comes back as text.

```
ask "What is your name?" into name
say "Hello, " plus name
```

### at the same time

Runs several actions at once instead of one after another — like putting
the kettle on *and* setting the table simultaneously, rather than doing
one, then the other. hotgrin waits for everything inside the block to
finish before moving on.

```
action knock on a door
    say "Knock knock!"
end action

at the same time
    do knock on a door
    do knock on a door
end at the same time
```

See also: [start](#start).

### comment

A note left for a human reading the code, not an instruction for the
computer — it starts with `#` and the computer skips the rest of that
line entirely. hotgrin programs use a numbered-comment habit for
organizing sections; see [Day One](day-one.md) for how and why.

```
# this whole line is a comment, ignored when the program runs
say "this line actually runs"   # so is this bit, after the #
```

### concatenation

Joining pieces of text together into one longer piece. In hotgrin this
is just `plus` — the same word used for adding numbers, except when
either side is text, in which case it joins them instead of adding them.

```
set name to "Thandi"
say "Hello, " plus name plus "!"
```

### conditional

See [if / else if / else](#if--else-if--else).

### contains

Checks whether a list or piece of text includes something specific.

```
set groceries to list of "bread", "milk", "eggs"
if groceries contains "milk"
    say "Don't forget you already have milk on the list"
end if
```

### decrease

Lowers a number by a given amount — the opposite of [increase](#increase).

```
set stock to 10
decrease stock by 3
say stock   # 7
```

### describe

See [record](#record) — `describe` is the word that starts one.

### error handling

The plan for what happens when something goes wrong — a full meal
running out of an ingredient, a file that isn't where it's supposed to
be. hotgrin makes you handle this on purpose: an action that can fail
must be called inside a [`try`](#try--if-it-fails) block, or the Watcher
won't let the program run. See [give back
problem](#give-back-problem) and [try / if it fails](#try--if-it-fails).

### expression

Anything that produces a value: a number, a piece of text, a variable, or
several of those combined with words like `plus` or `times`. `2 plus 2`
is an expression; so is a single variable name on its own.

### field

One named piece of information that belongs to a [record](#record) — the
same idea as one line on an index card, like "price" on a card that
describes a product.

```
describe order
    item is "Wireless mouse"
    price is 299
end describe

say price of order   # reads the field
```

### give back

Sends a value back out of an [action](#action) to whoever called it —
the answer the recipe hands you once it's done. Other languages call
this "return."

```
action double with n
    give back n times 2
end action

set result to double with 21
say result   # 42
```

### give back problem

Sends back a *failure* instead of a value, from inside an action — this
is what makes an action "fallible," meaning it must always be called
inside a [`try`](#try--if-it-fails) block.

```
action check age with age
    if age is less than 0
        give back problem "age can't be negative"
    end if
    give back age
end action

try
    set result to check age with 25
if it fails
    say "Something went wrong: " plus the problem
end try
```

### if / else if / else

Runs one block of steps or another, depending on whether something is
true — the "if it's raining, take an umbrella, otherwise leave it" rule
from [Day Zero](day-zero.md).

```
if temperature is greater than 30
    say "It's hot"
else if temperature is less than 10
    say "It's cold"
else
    say "It's mild"
end if
```

### increase

Raises a number by a given amount.

```
set stock to 7
increase stock by 3
say stock   # 10
```

### input

Lets a program accept options typed on the command line when it's run,
instead of always doing the exact same thing.

```
input name as text default "world"
say "Hello, " plus name
```
Run with `hotgrin run greet.hot --name AJ` to get `Hello, AJ` instead of
the default.

### is / is not / is greater than / is less than / is at least / is at most

The comparison words — how hotgrin asks "is this true?" about two
values.

```
if age is at least 18
    say "You can vote"
end if
```

### item _N_ of

Reads one specific thing out of a [list](#list) by its position, counting
from 0.

```
set groceries to list of "bread", "milk", "eggs"
say item 0 of groceries   # "bread" - the first one
```

### library

A `.hot` file whose actions you can borrow in your own program with
`use`, instead of writing them yourself — local, from hotgrin's own
standard library, or fetched straight from a GitHub repository. See the
[library guide](library-guide.md).

```
use "std/text"
```

### list

An ordered collection of things — a grocery list, exactly like the one
from [Day Zero](day-zero.md).

```
set groceries to list of "bread", "milk", "eggs"
put "butter" into groceries
say count of groceries   # 4
```
See also: [item _N_ of](#item-n-of), [contains](#contains), [put ... into](#put--into).

### loop

Repeats a block of steps — either a fixed number of times, while
something stays true, or once for every item in a list. The grocery-list
idea from [Day Zero](day-zero.md): pick an item, deal with it, next item,
until you're done.

```
repeat 3 times
    say "Knock!"
end repeat

repeat for each thing in list of "bread", "milk", "eggs"
    say "Buying: " plus thing
end repeat
```

### of

The small connecting word that means "belonging to" — used both to read
a record's [field](#field) and to work with a [list](#list). Which one
happens depends on what's on the right: a record gives you field access,
a list gives you a list operation.

```
say price of order        # field access - order is a record
say count of groceries    # list operation - groceries is a list
```

### operator

A word that combines two values into one result — `plus`, `times`, `is
greater than`, and so on are all operators.

### parameter

A named input an [action](#action) expects when you call it — the
"cups" in "make tea with 2 cups," standing in for whatever number gets
passed in each time.

```
action greet with name
    say "Hello, " plus name
end action

greet with "Thandi"
greet with "Sipho"
```

### plus / minus / times / divided by

The arithmetic words. `plus` also joins text together (see
[concatenation](#concatenation)); `divided by` always gives you a
decimal answer, even if the numbers divide evenly.

```
say 7 divided by 2   # 3.5
```

### put ... into

Adds a new item onto the end of a [list](#list).

```
set groceries to list of "bread", "milk"
put "eggs" into groceries
```

### record

A named bundle of related information — the "index card" for one thing,
like one customer or one order, with a fixed set of labelled fields. You
declare its shape once with `describe`, then read and write its fields
by name.

```
describe order
    item is "Wireless mouse"
    price is 299
end describe

say item of order              # field read
set price of order to 249      # field write
```

### repeat

See [loop](#loop) — `repeat` is the word that starts one.

### reserved word

A word hotgrin already has a job for — `to`, `of`, `with`, `is`, `times`,
and a small fixed set of others — which means it can't also be used as
part of a variable's own name. If you see a confusing error about a name
you were sure was fine, a reserved word hiding inside it is a common
cause; the [prompt pack](ai-prompt-pack.md) has the full list.

### rounded to

Rounds a number to a given number of decimal places, always giving back
a decimal.

```
say 10 divided by 3 rounded to 2   # 3.33
```

### say

Prints something to the screen — the very first word most people learn,
in [Lesson 01](../examples/learn/01-say-hello.hot).

```
say "Hello, world!"
```

### set

Creates a new [variable](#variable), or changes the value of one that
already exists.

```
set name to "Thandi"     # creates it
set name to "Sipho"       # changes it
```

### start

Runs an [action](#action) in the background and moves on immediately,
without waiting for it to finish — the program only waits for it right
at the very end, before exiting.

```
action announce
    say "This runs in the background"
end action

start announce
```
See also: [at the same time](#at-the-same-time).

### stop with error

Ends the program immediately with a message printed to the error output
and a failure exit code — for situations a program genuinely can't
recover from.

```
stop with error "no file given"
```

### test

A small, separate check that proves a piece of your program behaves the
way you expect — written once, then run automatically any time you want
reassurance nothing broke.

```
action double with n
    give back n times 2
end action

test "double doubles a number"
    expect double with 21 to be 42
end test
```
Run with `hotgrin test file.hot`.

### try / if it fails

Wraps a call to an action that might fail, so the failure is handled on
purpose instead of crashing the program unexpectedly.

```
action check age with age
    if age is less than 0
        give back problem "age can't be negative"
    end if
    give back age
end action

try
    set result to check age with -5
if it fails
    say "Something went wrong: " plus the problem
end try
```

### type

What *kind* of value something is — text, a whole number, a decimal,
true-or-false, a list, or a record. hotgrin figures this out on its own
from context; you never have to state it yourself.

### unit of measure

A measurement hotgrin understands natively — kilograms, metres, hours,
and so on — so mismatched measurements (adding a weight to a length) get
caught as an error before your program even runs, and correct
conversions happen automatically.

```
set weight to 129 kg
say weight in g          # 129000 g
```

### use

Borrows another `.hot` file's actions into your own program. See
[library](#library).

### use go

An escape hatch for writing real Go code directly inside a hotgrin
program, for the rare moment hotgrin itself doesn't have what you need
yet. Not something you'll need starting out.

```
use go
import "strings"
func shoutCase(s string) string { return strings.ToUpper(s) + "!" }
end go

say shout case with "howzit"
```

### variable

A labelled space that holds a value, where the value is allowed to
change while the label stays the same — the labelled jar from [Day
Zero](day-zero.md).

```
set sugar to "half a jar"
say sugar
```

### Watcher

hotgrin's always-on checker. It reads your program before running it and
only ever reports things that are genuinely real problems — typos,
unhandled failures, unreachable code — never a false alarm. If the
Watcher finds nothing, it says so plainly: "All good - I found no
problems."

### with

The word used both to pass arguments into an [action](#action) and to
combine measurements of the same kind.

```
greet with "Thandi"
set total to 2 km plus 500 m
```
