# beautifi v0.1.0

Batch logo generation CLI using Google's Imagen 3 API.

## Installation

```bash
cd ~/code/beautifi
go build -o beautifi .
# Optional: symlink to PATH
ln -sf $(pwd)/beautifi ~/.local/bin/beautifi
```

## Configuration

Create project configs in `~/.config/beautifi/projects/`:

```yaml
# ~/.config/beautifi/projects/myproject.yaml
project: myproject
tagline: "Your awesome tool"

themes:
  - nature
  - tech
  - abstract

styles:
  - flat-minimal
  - gradient-glass
  - neon-glow
  - geometric

# Optional
aspect_ratio: "1:1"
base_prompt: "Custom prompt override"
```

### Available Styles

| Style | Description |
|-------|-------------|
| `flat-minimal` | Clean flat design, no shadows |
| `gradient-glass` | Modern glassmorphism |
| `neon-glow` | Cyberpunk neon aesthetic |
| `hand-drawn` | Sketchy organic style |
| `3d-render` | Realistic 3D depth |
| `retro-pixel` | 8-bit pixel art |
| `watercolor` | Soft paint texture |
| `geometric` | Bold abstract shapes |

## Usage

```bash
# Preview prompts without API calls
beautifi preview bosun
beautifi preview bosun --format markdown

# Dry run - see what would be generated
beautifi generate bosun --dry-run --verbose

# Generate with variants
beautifi generate bosun --variants 3

# Filter styles
beautifi generate bosun --styles flat-minimal,neon-glow

# Batch multiple projects
beautifi batch bosun wasp clint
beautifi batch  # processes all projects in config dir
```

## Output

Images saved to `~/output/beautifi/<project>/`:
```
bosun/
├── nautical-flat-minimal-1.png
├── nautical-flat-minimal-1.json  (metadata)
├── nautical-gradient-glass-1.png
├── ...
```

## Environment

```bash
export GEMINI_API_KEY="your-api-key"
```

## API Integration

Uses Google's Imagen 3 endpoint:
```
POST https://generativelanguage.googleapis.com/v1beta/models/imagen-3.0-generate-001:predict
```

If you don't have Imagen API access yet, use `--dry-run` or `--prompts-only` to generate the prompts for use with other tools.
