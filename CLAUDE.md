# beautifi

> Batch logo and asset generation CLI — elevate project aesthetics programmatically.

Generate professional logos, icons, and visual assets for multiple projects using AI image generation APIs. Designed for developers who want their repos to look as good as their code.

---

## Quick Orient

```bash
beautifi generate bosun              # Generate logos for a project
beautifi batch projects.yaml         # Batch generate for multiple projects
beautifi preview ~/output/           # Open gallery to review
beautifi readme bosun --logo best    # Generate README with selected logo
```

---

## Architecture

```
beautifi (Go CLI, Cobra)
├── cmd/
│   ├── root.go           # Global flags, config loading
│   ├── generate.go       # Single project generation
│   ├── batch.go          # Multi-project batch mode
│   ├── preview.go        # HTML gallery generator
│   └── readme.go         # README template application
├── internal/
│   ├── backends/         # Image generation backends
│   │   ├── gemini.go     # Google Imagen API
│   │   ├── openai.go     # DALL-E (future)
│   │   └── backend.go    # Interface
│   ├── prompts/          # Prompt template engine
│   ├── gallery/          # HTML gallery generator
│   └── config/           # YAML config handling
└── templates/
    ├── prompts/          # Per-project prompt templates
    └── readme/           # README templates
```

---

## Workflow

### 1. Define Project Prompts

```yaml
# ~/.config/beautifi/projects/bosun.yaml
project: bosun
tagline: "Orchestration CLI for agentic ecosystems"
themes:
  - nautical (anchor, ship's wheel, compass)
  - command center (dashboard, controls)
  - conductor (orchestra, baton)
styles:
  - flat-minimal     # Simple, modern, flat design
  - gradient-glass   # Glassmorphism with gradients
  - line-art         # Single-weight line illustrations
  - mascot           # Character-based logo
  - abstract-geo     # Abstract geometric shapes
```

### 2. Generate Variants

```bash
beautifi generate bosun --variants 3 --styles all
# Generates: 5 styles × 3 variants × 3 themes = 45 images
```

### 3. Review & Select

```bash
beautifi preview ~/output/bosun/
# Opens HTML gallery for quick review
```

### 4. Apply to README

```bash
beautifi readme bosun --logo bosun-nautical-flat-1.png
# Generates professional README with logo, badges
```

---

## Configuration

```yaml
# ~/.config/beautifi/config.yaml
backend: gemini
api_key_env: GEMINI_API_KEY
output_dir: ~/output/beautifi
defaults:
  variants_per_prompt: 3
  prompts_per_project: 5
  parallel: 4
image:
  width: 512
  height: 512
  format: png
```

---

## Target Projects

Initial batch:
1. **bosun** — Orchestration CLI (nautical theme)
2. **polecat** — Sandboxed Claude runner (mustelid/ferret theme)
3. **hurtlocker** — Agent journaling (vault/lock theme)
4. **noodle** — Exocortex (brain/noodle theme)
5. **jobsworth** — Job automation (briefcase/robot theme)
6. **halhq** — HAL operations (red eye/AI theme)

---

## Current Focus

**Active:** Core CLI framework + Gemini integration
**Next:** Prompt template system, batch mode
**Blocked:** —

---

## Dependencies

- Go 1.21+
- `github.com/spf13/cobra` — CLI framework
- `gopkg.in/yaml.v3` — YAML config
- Google Imagen API (via Gemini API key)
