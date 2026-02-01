# Architecture Rules

The physics of this world. Don't guess — follow these.

---

## Vibe

What should this codebase *feel* like? (Delete options that don't apply)

- [ ] **Concise** — minimal, no bloat, every line earns its place
- [ ] **Hacky** — get it working, polish later
- [ ] **Robust** — defensive, handles edge cases, production-ready
- [ ] **Playful** — personality allowed, not corporate
- [ ] **Enterprise** — formal, documented, conventional

If unmarked, you'll get generic code. Be specific.

---

## Tech Stack

- **Language:** 
- **Framework:** 
- **Database:** 
- **Styling:** 

---

## Patterns (Do This)

- Business logic lives in `/src/services/`, not controllers
- Small, composable functions over monolithic ones
- Types for everything — no implicit `any`
- Error handling at boundaries, not scattered throughout

---

## Anti-Patterns (Never Do This)

- Never use `any` in TypeScript
- Never leave comments describing *what* code does — only *why*
- Never hardcode secrets or config values
- Never commit `.env` files
- Never bypass the Gate

---

## File Organization

```
src/
├── api/          # HTTP handlers, routes
├── services/     # Business logic
├── models/       # Data types, schemas
├── utils/        # Pure helper functions
└── config/       # Environment, constants
```

---

## Naming Conventions

- Files: `kebab-case.ts`
- Functions: `camelCase`
- Types/Classes: `PascalCase`
- Constants: `SCREAMING_SNAKE_CASE`

---

## Testing Philosophy

- Test behavior, not implementation
- One assertion per test when possible
- Tests are documentation — name them clearly
