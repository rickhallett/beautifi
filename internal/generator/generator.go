package generator

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rickhallett/beautifi/internal/api"
	"github.com/rickhallett/beautifi/internal/config"
)

// PromptSpec represents a single generation task
type PromptSpec struct {
	Theme    string `json:"theme"`
	Style    string `json:"style"`
	Variant  int    `json:"variant"`
	Prompt   string `json:"prompt"`
	Filename string `json:"filename"`
}

// GenerationResult captures the outcome of a single generation
type GenerationResult struct {
	Spec     PromptSpec `json:"spec"`
	Success  bool       `json:"success"`
	Error    string     `json:"error,omitempty"`
	FilePath string     `json:"file_path,omitempty"`
}

// GeneratePrompts creates all prompt combinations for a project
func GeneratePrompts(proj *config.Project, styles []string, variants int) []PromptSpec {
	var prompts []PromptSpec
	stylePresets := config.DefaultStyles()

	for _, theme := range proj.Themes {
		for _, style := range styles {
			for v := 1; v <= variants; v++ {
				prompt := buildPrompt(proj, theme, style, stylePresets)
				filename := buildFilename(theme, style, v)

				prompts = append(prompts, PromptSpec{
					Theme:    theme,
					Style:    style,
					Variant:  v,
					Prompt:   prompt,
					Filename: filename,
				})
			}
		}
	}

	return prompts
}

func buildPrompt(proj *config.Project, theme, style string, presets map[string]config.StylePreset) string {
	var parts []string

	// Base: logo description
	parts = append(parts, fmt.Sprintf("A professional logo icon for '%s'", proj.Project))

	if proj.Tagline != "" {
		parts = append(parts, fmt.Sprintf("a %s tool", strings.ToLower(proj.Tagline)))
	}

	// Theme
	parts = append(parts, fmt.Sprintf("with a %s theme", theme))

	// Style keywords
	if preset, ok := presets[style]; ok {
		parts = append(parts, strings.Join(preset.Keywords, ", "))
	} else {
		parts = append(parts, style+" style")
	}

	// Standard quality additions
	parts = append(parts, "high quality", "suitable for app icon", "centered composition", "white or transparent background")

	// Custom base prompt override
	if proj.BasePrompt != "" {
		return proj.BasePrompt + ". " + strings.Join(parts[2:], ", ")
	}

	return strings.Join(parts, ", ")
}

func buildFilename(theme, style string, variant int) string {
	// Sanitize names for filesystem
	theme = sanitizeName(theme)
	style = sanitizeName(style)
	return fmt.Sprintf("%s-%s-%d.png", theme, style, variant)
}

func sanitizeName(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "_", "-")
	return s
}

// GenerateImages calls the API for each prompt and saves results
func GenerateImages(client *api.GeminiClient, prompts []PromptSpec, outDir string, verbose bool) ([]GenerationResult, error) {
	var results []GenerationResult
	ctx := context.Background()

	for i, spec := range prompts {
		if verbose {
			fmt.Printf("[%d/%d] Generating %s...\n", i+1, len(prompts), spec.Filename)
		}

		result := GenerationResult{Spec: spec}
		outPath := filepath.Join(outDir, spec.Filename)

		// Generate image
		imageData, err := client.GenerateImage(ctx, spec.Prompt)
		if err != nil {
			result.Error = err.Error()
			result.Success = false
			results = append(results, result)

			if verbose {
				fmt.Printf("  Error: %v\n", err)
			}
			continue
		}

		// Save image
		if err := os.WriteFile(outPath, imageData, 0644); err != nil {
			result.Error = fmt.Sprintf("save failed: %v", err)
			result.Success = false
			results = append(results, result)
			continue
		}

		// Save metadata
		metaPath := outPath[:len(outPath)-4] + ".json"
		meta := map[string]interface{}{
			"prompt":  spec.Prompt,
			"theme":   spec.Theme,
			"style":   spec.Style,
			"variant": spec.Variant,
		}
		metaData, _ := json.MarshalIndent(meta, "", "  ")
		os.WriteFile(metaPath, metaData, 0644)

		result.Success = true
		result.FilePath = outPath
		results = append(results, result)

		if verbose {
			fmt.Printf("  Saved: %s\n", outPath)
		}
	}

	return results, nil
}
