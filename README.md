# beautifi ✨

Batch logo and asset generation CLI for developers who want their repos to look as good as their code.

## Why

Professional-looking repos signal quality. Good logos, badges, and visual polish make a difference. But creating them manually is tedious. Beautifi automates the creative exploration phase — generate dozens of variants, pick the best, move on.

## Install

```bash
go install github.com/rickhallett/beautifi@latest
```

Requires: `GEMINI_API_KEY` in environment.

## Usage

```bash
# Generate logos for a project
beautifi generate bosun --variants 3

# Batch generate for multiple projects
beautifi batch projects.yaml

# Preview generated images
beautifi preview ~/output/bosun/

# Apply logo to README
beautifi readme bosun --logo best
```

## How It Works

1. **Define themes** — What visual concepts fit your project?
2. **Define styles** — Flat, gradient, line-art, mascot, abstract?
3. **Generate variants** — N images per prompt, M prompts per project
4. **Review** — HTML gallery for quick selection
5. **Apply** — Insert into professional README template

## Example

```yaml
# ~/.config/beautifi/projects/polecat.yaml
project: polecat
tagline: "Sandboxed Claude CLI runner"
themes:
  - mustelid (ferret, polecat, weasel)
  - burrow (underground, contained)
  - sandbox (playful, safe)
styles:
  - flat-minimal
  - mascot
  - abstract-geo
```

```bash
beautifi generate polecat --variants 3
# → 3 themes × 3 styles × 3 variants = 27 images
# → ~/output/polecat/
```

## Configuration

```yaml
# ~/.config/beautifi/config.yaml
backend: gemini
output_dir: ~/output/beautifi
defaults:
  variants_per_prompt: 3
  parallel: 4
```

## License

MIT
