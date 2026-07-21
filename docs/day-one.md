# Day One: your first real program

Yesterday (or five minutes ago, however this finds you) you did
[Day Zero](day-zero.md) and proved something to yourself: you already
think in algorithms, loops, decisions, variables, and actions. You just
never had the words for it.

Today you're going to do two things for the first time: plan a program
properly, on purpose, before typing anything — and then actually type
it, run it, and watch it work. Both matter. Most people skip the first
one and wonder why the second one feels harder than it should.

---

## Open the playground

You'll need it for the second half of this page, so open it now and
leave it in another tab: **[the browser playground](https://hotgrin.github.io/hotgrin/playground/)**.
Type on the left, press Run, see what happens on the right. Nothing to
install. Nothing you type here can damage anything — the worst case is
a friendly message telling you what to fix.

Don't type anything yet. First, the part almost nobody teaches you.

---

## Before you type anything

Here's a habit worth stealing from every experienced programmer, even
though most of us learned it the hard way instead of being told: **you
plan the program in plain words before you write a single line of
code.** Four questions, always in this order:

1. **What do I want this program to do for me?**
2. **What does it need from me to do that?**
3. **What should it show me at the end?**
4. **What could go wrong?**

That's it. That's the whole ritual. Let's use it for real, on the
tea program from Day Zero.

**1. What do I want this program to do for me?**
Walk me through making a cup of tea, in order, so I never forget a
step — especially before my first coffee of the day, when I forget
everything.

**2. What does it need from me?**
Nothing yet. Today it just recites the steps. (Later, in future
lessons, we could ask it how many cups — not today's problem, and
that's fine. A program is allowed to start small.)

**3. What should it show me at the end?**
Each step, printed to the screen, in the right order.

**4. What could go wrong?**
Honestly? Not much, for something this small. That question won't
always have a big answer — sometimes the honest answer is "not much,"
and that's a fine answer too. The habit of *asking* it every time is
what matters. As your programs grow, this question is the one that
saves you — it's the same instinct behind hotgrin's `try` / `if it
fails`, which you'll meet properly in a later lesson.

Four questions, four honest answers, and you haven't typed a single
line of code yet. That's not stalling — that's the actual first step.

---

## Turning your plan into a numbered comment skeleton

Here's where it gets real. Open any serious hotgrin program — the
Invoice Maker example that ships with hotgrin is a good one — and
you'll find every single one starts the same way: a numbered list of
what's coming, right at the top, as a comment.

```
# ============================================
#  SECTION INDEX
# ============================================
#  [1]  Make the tea   - the steps, in order
# ============================================
```

That's not decoration. It's a table of contents for a recipe book —
and it means that six months from now, you (or anyone else) can open
this file and know exactly what it does before reading a single real
line. You're about to build the tiny version of that habit, from your
very first program, so it's never something you have to "add later."

Our plan only has one real section, so our skeleton is short. Type
this into the playground now — just the comments, nothing else yet:

```
# ============================================
#  MAKE THE TEA
# ============================================
#  [1]  Make the tea   - the steps, in order
# ============================================

# [1] MAKE THE TEA
```

Press Run. Nothing happens — and that's correct! Comments are notes
for humans; the computer skips every line starting with `#`. You've
just written a program that does nothing at all, on purpose, and
that's a perfectly good place to start from.

---

## Comment-first, code-second

Now fill in section `[1]`, one line at a time, straight from your plan
above. Add these lines under `# [1] MAKE THE TEA`:

```
# ============================================
#  MAKE THE TEA
# ============================================
#  [1]  Make the tea   - the steps, in order
# ============================================

# [1] MAKE THE TEA
say "Boil the kettle"
say "Warm the pot"
say "Add the tea leaves"
say "Pour in the hot water"
say "Wait four minutes"
say "Pour and serve"
```

Press Run. You should see all six lines printed, in order, on the
right. That's it. **That's your first real, planned, working hotgrin
program** — typed by you, run by you, and you knew exactly what every
line was for before you wrote it, because you planned it first.

---

## A quick word on folders and files

Right now your whole program lives in one file, and for something this
size, that's exactly right — don't overthink it. Later, as programs
grow bigger, you'll split related things into more files, kept
together inside a folder — the same idea as moving from one recipe
card into a labelled recipe box once you've got more than a handful of
recipes. That's a problem for a future lesson, not today. One file is
the correct answer for a program this size.

## A quick word on "what else do I need to install?"

Nothing. This is genuinely good news, and it's worth saying out loud
because most tutorials never do: for programs like the one you just
wrote — and for quite a while after this — hotgrin needs nothing else.
No package manager, no extra libraries, no accounts to create. One
tool, one file, one command. If that ever changes for something more
advanced, a much later lesson will tell you exactly what's needed and
why. Today, the honest answer is: you already have everything you
need.

---

## What you actually just did

You asked four honest questions before typing anything. You turned
your answers into a numbered plan, written as comments, the same
habit every real hotgrin program uses. You filled it in one line at a
time. You ran it, and it worked, because you understood every part of
it before you wrote it.

That loop — plan, skeleton, fill in, run — is the whole job. Everything
from here is just learning new words to put inside that same loop.

When you're ready, **[Lesson 01 — say hello](../examples/learn/01-say-hello.hot)**
picks up from exactly here. And whenever you hit a word in someone else's
hotgrin code that you don't recognize yet, the **[glossary](glossary.md)**
has a plain-language explanation and a working example for every one.

Or, if you'd rather keep going with this story a little longer before the
numbered lessons: **[Day Two](day-two.md)** continues right from here,
with your first real decision.
