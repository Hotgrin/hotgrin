# Day Two: making choices

On [Day Zero](day-zero.md) you already worked out the umbrella rule: if
it's raining, take an umbrella, otherwise leave it. On [Day
One](day-one.md) you planned a program properly before typing it, gave it
a numbered comment skeleton, and ran your first real code.

Today you do both at once: plan a decision, then actually type it.

---

## Before you type anything

Same four questions as always. Quicker this time — you've done this
before.

1. **What do I want this program to do for me?** Tell me whether to take
   an umbrella, based on whether it's raining.
2. **What does it need from me?** Whether it's raining right now.
3. **What should it show me?** A clear instruction — take it, or leave
   it.
4. **What could go wrong?** Not much, for something this small — same
   honest answer as last time. (One real limitation worth naming: today
   the program only knows what *you* tell it. Teaching it to check the
   actual weather is a job for a much later lesson — `ask`ing the person
   running it is the first small step in that direction, and you'll meet
   that soon.)

---

## The labelled jar, for real this time

Remember the labelled jar from Day Zero — a jam jar labelled "sugar,"
where what's inside can change while the label stays the same? That's a
**variable**, and today you're actually going to make one, instead of
just picturing one.

Open the [playground](https://hotgrin.github.io/hotgrin/playground/) and
type this first, on its own:

```
set raining to true
```

That line makes a jar labelled `raining`, and puts `true` inside it —
`true` meaning "yes, it is." Nothing prints yet; you've just told the
program a fact. The next part is where that fact gets *used*.

---

## From plan to skeleton

Same numbered-comment habit as Day One. Two things need to happen: say
what's true, then decide what to do about it. Type this skeleton:

```
# ============================================
#  SHOULD I TAKE AN UMBRELLA?
# ============================================
#  [1]  Say whether it's raining today
#  [2]  Decide what to do about it
# ============================================

# [1] SAY WHETHER IT'S RAINING TODAY

# [2] DECIDE WHAT TO DO ABOUT IT
```

Press Run. Nothing happens — correct, same as Day One. Comments don't do
anything; they're the plan, not the program yet.

---

## Filling it in — the decision itself

Programmers call an "if this, then that" rule a **decision** — the exact
word from Day Zero. Here's the whole thing:

```
# ============================================
#  SHOULD I TAKE AN UMBRELLA?
# ============================================
#  [1]  Say whether it's raining today
#  [2]  Decide what to do about it
# ============================================

# [1] SAY WHETHER IT'S RAINING TODAY
set raining to true

# [2] DECIDE WHAT TO DO ABOUT IT
if raining
    say "Take the umbrella"
else
    say "Leave the umbrella"
end if
```

Run it. You should see `Take the umbrella`.

**Now the important part.** Go back to line 8, change `true` to `false`,
and run it again. You should see `Leave the umbrella` — a completely
different result, from the exact same decision, because you changed one
fact instead of rewriting the plan. That's genuinely most of what
programming is: writing the decision once, correctly, and letting the
facts change what happens.

---

## One more branch — else if

Real days aren't just "raining or not." Add a second question — is it
very sunny instead? — using `else if`:

```
set raining to false
set very sunny to true

if raining
    say "Take the umbrella"
else if very sunny
    say "Take sunscreen and a hat"
else
    say "You're sorted - just go"
end if
```

Run it, then try flipping each `true`/`false` and running it again. Three
possible outcomes, one decision, same habit as before: change a fact, not
the plan.

---

## What you actually did

You planned a decision in plain words before typing anything. You made
your first real variable — a labelled jar, actually holding something
this time. You wrote the decision once, then proved to yourself that
changing a fact changes the outcome without touching the logic itself.
That's the whole idea behind "if this, then that" — you already knew it
from Day Zero; today you just typed it.

When you're ready, **[Lesson 04 — making
choices](../examples/learn/04-making-choices.hot)** has a few more
examples of the same idea. (Lessons
[02](../examples/learn/02-remembering-things.hot) and
[03](../examples/learn/03-maths-that-behaves.hot) are worth a look too,
if you'd like more practice with variables and arithmetic before moving
on — not required, just there if you want it.)

Or, if you'd rather keep going with this story instead of jumping into
the numbered lessons: Day Three continues right from here, with loops.
