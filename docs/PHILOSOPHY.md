# base-project

> A scaffold for humans who've realized they're not coders anymore — they're managers of tireless digital juniors who occasionally hallucinate.

---

## The Premise

You've heard the hype. AI writes code now. But you've also felt the pain: ChatGPT spits out garbage, you paste it in, it breaks, you paste the error back, it apologizes and generates different garbage.

That's not agentic engineering. That's **vibe coding** — and it's a slur around here.

This scaffold exists because there's a better way. One where:

- The agent runs its own code and sees its own errors
- You review *outcomes*, not every line of generated boilerplate  
- "Good enough" ships, "perfect" emerges through iteration
- You stop being a typist and start being an architect

---

## What's in the Box

```
.
├── .agent/
│   ├── architecture.md      # The rules. Tech stack, patterns, vibe.
│   └── product_vision.md    # The soul. What are we building and why.
├── .claude/commands/        # Slash commands for tactical modes
│   ├── architect.md         # /architect — plan before you build
│   ├── weaver.md            # /weaver — implement until gate passes
│   ├── auditor.md           # /auditor — review against the rules
│   ├── spike.md             # /spike — creative exploration, intentionally vague
│   ├── gate.md              # /gate — verify your work
│   └── bootstrap.md         # /bootstrap — self-heal the environment
├── .prompts/                # Reusable prompt templates
│   ├── feature_spec.md      # Spec it before you build it
│   ├── refactor.md          # How we refactor around here
│   ├── review.md            # The audit checklist
│   └── spike.md             # Under-prompting for creativity
├── bin/
│   ├── bootstrap            # Self-healing setup (deps, env, migrations)
│   ├── gate                  # The verification loop (lint, typecheck, test)
│   ├── db                   # Database CLI wrapper
│   ├── logs                 # Fetch recent logs
│   └── docs                 # Dump project context for the agent
├── AGENTS.md                # How agents should behave here
├── CLAUDE.md                # Project context (the agent reads this first)
└── learning.log             # Structured capture for your own spaced repetition
```

---

## The Philosophy

### Close the Loop

```
Read → Write → Execute → Fix
```

If your agent can't run its own code, you're just a slow copy-paste intermediary. The agent must execute, see errors, and fix them without you in the middle.

`./bin/gate` is the gate. Nothing ships until exit code 0.

### You're a Manager Now

Stop micromanaging. Give goals, not step-by-step instructions.

Bad: "Create a file called UserService.ts with a method called..."
Good: "Refactor auth to a plugin architecture. Follow patterns in `src/plugins/`."

When the agent makes a silly mistake (and it will), don't rewrite it yourself. Just say: "You forgot to check the import path. Try again."

You're the building inspector, not the bricklayer.

### Vibe is a Technical Requirement

LLMs default to Generic Corporate Code™ — soulless, over-engineered, enterprise bloat.

You must explicitly specify the aesthetic:
- *Concise* — every line earns its place
- *Hacky* — ship it, polish later  
- *Robust* — edge cases matter
- *Playful* — personality allowed

If you don't specify, you get the average of the internet. Which is mediocre.

### Anti-Fragile > Robust

Old thinking: Prevent errors at all costs.
New thinking: Recover from errors cheaply.

For most projects, the cost of *fixing* is now so low that the cost of *preventing* becomes hard to justify. Build the recovery harness, not the prevention fortress.

---

## The Workflow

### 1. `/architect` — Plan

Read the vision. Understand the constraints. Output a plan.

No code yet. Just: what files, what order, what risks.

### 2. `/weaver` — Build

Implement the plan. Follow existing patterns. Run the gate.

If gate fails: read output, fix, repeat.
If gate passes: commit.

Do not stop until `./bin/gate` returns 0.

### 3. `/auditor` — Review

Check against `.agent/architecture.md`. Look for security issues. Verify style.

Fix violations before merge.

### 4. `/spike` — Explore (Optional)

When you don't know the approach yet, under-prompt on purpose:

> "Make this work. Surprise me."

Wrong? Revert. Cheap.
Right? Free creativity.

---

## Quick Start

```bash
# Spawn a new project (if you have the alias)
spawn my-project

# Or manually
cp -r /path/to/base-project ./my-project
cd my-project
./bin/bootstrap
```

Then:
1. Fill in `.agent/product_vision.md` — what are you building?
2. Fill in `.agent/architecture.md` — what's the stack and vibe?
3. Fill in `CLAUDE.md` — what's the current focus?
4. Start building.

---

## The Gate

Your definition of "done":

```bash
./bin/gate
```

This runs:
- Type checking (if TypeScript)
- Linting
- Tests
- Build verification

Exit 0 = ship it.
Exit non-zero = fix it first.

The agent knows it must pass the gate. So should you.

---

## The Learning Log

This isn't just for the agent. It's for you.

`learning.log` captures your journey:
- What you set out to learn
- What you predicted
- What you actually found
- Where your mental model was wrong
- Retention hooks for spaced repetition

Because building software is learning in disguise, and vibe learning is just as bad as vibe coding.

---

## Project Types

Not everything belongs here. This scaffold works best for:

| Type | Description |
|------|-------------|
| **Micro-software** | Market of 1, utility of 10. Too niche to be a startup. |
| **Boring plumbing** | Glue code, integrations, data transformation. |
| **Unstructured reality** | Messy input → structured output. Model as UI. |

Avoid: Precision-critical systems where failure = catastrophe. Video codecs, HFT, pacemakers.

---

## The Mantra

> Stop writing code. Start building loops.
>
> Your value is defining the destination and building the guardrails.
> Not driving the car.

---

## Credits

Philosophy extracted from [Peter Steinberger](https://steipete.com)'s agentic engineering practice.

Scaffold assembled by humans and their tireless digital juniors.

---

*Now go build something that matters to exactly one person: you.*
