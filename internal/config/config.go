package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Project represents a beautifi project configuration
type Project struct {
	Project  string   `yaml:"project"`
	Tagline  string   `yaml:"tagline"`
	Themes   []string `yaml:"themes"`
	Styles   []string `yaml:"styles"`

	// Optional overrides
	AspectRatio string            `yaml:"aspect_ratio,omitempty"` // e.g., "1:1", "16:9"
	BasePrompt  string            `yaml:"base_prompt,omitempty"`  // Custom base prompt
	Extras      map[string]string `yaml:"extras,omitempty"`       // Additional template vars
}

// StylePreset defines a reusable style configuration
type StylePreset struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Keywords    []string `yaml:"keywords"`
}

// DefaultStyles returns built-in style presets
func DefaultStyles() map[string]StylePreset {
	return map[string]StylePreset{
		"flat-minimal": {
			Name:        "flat-minimal",
			Description: "Clean flat design with minimal details",
			Keywords:    []string{"flat design", "minimal", "clean lines", "no shadows", "solid colors"},
		},
		"gradient-glass": {
			Name:        "gradient-glass",
			Description: "Modern glassmorphism with gradients",
			Keywords:    []string{"gradient", "glassmorphism", "translucent", "modern", "soft shadows"},
		},
		"neon-glow": {
			Name:        "neon-glow",
			Description: "Cyberpunk neon aesthetic",
			Keywords:    []string{"neon", "glow", "cyberpunk", "dark background", "bright colors"},
		},
		"hand-drawn": {
			Name:        "hand-drawn",
			Description: "Sketchy hand-drawn style",
			Keywords:    []string{"hand-drawn", "sketch", "organic lines", "imperfect", "artistic"},
		},
		"3d-render": {
			Name:        "3d-render",
			Description: "3D rendered with depth",
			Keywords:    []string{"3D render", "depth", "realistic lighting", "shadows", "perspective"},
		},
		"retro-pixel": {
			Name:        "retro-pixel",
			Description: "8-bit pixel art style",
			Keywords:    []string{"pixel art", "8-bit", "retro", "limited palette", "nostalgic"},
		},
		"watercolor": {
			Name:        "watercolor",
			Description: "Soft watercolor painting",
			Keywords:    []string{"watercolor", "soft edges", "paint texture", "artistic", "flowing"},
		},
		"geometric": {
			Name:        "geometric",
			Description: "Bold geometric shapes",
			Keywords:    []string{"geometric", "shapes", "bold", "abstract", "mathematical"},
		},
	}
}

// LoadProject loads a project configuration from YAML file
func LoadProject(path string) (*Project, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var proj Project
	if err := yaml.Unmarshal(data, &proj); err != nil {
		return nil, fmt.Errorf("parse yaml: %w", err)
	}

	// Validate required fields
	if proj.Project == "" {
		return nil, fmt.Errorf("missing required field: project")
	}
	if len(proj.Themes) == 0 {
		return nil, fmt.Errorf("missing required field: themes")
	}
	if len(proj.Styles) == 0 {
		// Use defaults
		proj.Styles = []string{"flat-minimal", "gradient-glass"}
	}

	return &proj, nil
}

// SaveProject writes a project configuration to YAML file
func SaveProject(path string, proj *Project) error {
	data, err := yaml.Marshal(proj)
	if err != nil {
		return fmt.Errorf("marshal yaml: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
