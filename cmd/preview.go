package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rickhallett/beautifi/internal/config"
	"github.com/rickhallett/beautifi/internal/generator"
	"github.com/spf13/cobra"
)

var (
	previewFormat string
	previewLimit  int
)

var previewCmd = &cobra.Command{
	Use:   "preview <project>",
	Short: "Preview prompts without generating images",
	Long: `Preview the prompts that would be generated for a project.
	
Useful for reviewing and tweaking your project config before spending API credits.`,
	Args: cobra.ExactArgs(1),
	RunE: runPreview,
}

func init() {
	rootCmd.AddCommand(previewCmd)

	previewCmd.Flags().StringVarP(&previewFormat, "format", "f", "text", "output format (text, json, markdown)")
	previewCmd.Flags().IntVarP(&previewLimit, "limit", "l", 0, "limit number of prompts shown (0 = all)")
	previewCmd.Flags().IntVarP(&variants, "variants", "n", 1, "number of variants per combination")
	previewCmd.Flags().StringSliceVarP(&styles, "styles", "s", []string{"all"}, "styles to preview")
}

func runPreview(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	cfgPath := filepath.Join(cfgDir, "projects", projectName+".yaml")
	proj, err := config.LoadProject(cfgPath)
	if err != nil {
		return fmt.Errorf("failed to load project config: %w", err)
	}

	// Check for existing outputs
	projectOutDir := filepath.Join(outDir, proj.Project)
	existingCount := countExistingImages(projectOutDir)

	// Filter styles
	activeStyles := proj.Styles
	if len(styles) > 0 && styles[0] != "all" {
		activeStyles = filterStyles(proj.Styles, styles)
	}

	prompts := generator.GeneratePrompts(proj, activeStyles, variants)

	if previewLimit > 0 && previewLimit < len(prompts) {
		prompts = prompts[:previewLimit]
	}

	switch previewFormat {
	case "json":
		printPromptsJSON(proj, prompts)
	case "markdown":
		printPromptsMarkdown(proj, prompts, existingCount)
	default:
		printPromptsText(proj, prompts, existingCount)
	}

	return nil
}

func countExistingImages(dir string) int {
	count := 0
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(path, ".png") {
			count++
		}
		return nil
	})
	return count
}

func printPromptsText(proj *config.Project, prompts []generator.PromptSpec, existing int) {
	fmt.Printf("Project: %s\n", proj.Project)
	fmt.Printf("Tagline: %s\n", proj.Tagline)
	fmt.Printf("Themes:  %v\n", proj.Themes)
	fmt.Printf("Styles:  %v\n", proj.Styles)
	if existing > 0 {
		fmt.Printf("Existing: %d images\n", existing)
	}
	fmt.Printf("\n%d prompts to generate:\n", len(prompts))
	fmt.Println(strings.Repeat("─", 60))

	for i, p := range prompts {
		fmt.Printf("\n[%d] %s\n", i+1, p.Filename)
		fmt.Printf("Prompt:\n%s\n", p.Prompt)
	}
}

func printPromptsJSON(proj *config.Project, prompts []generator.PromptSpec) {
	fmt.Println("{")
	fmt.Printf("  \"project\": \"%s\",\n", proj.Project)
	fmt.Printf("  \"count\": %d,\n", len(prompts))
	fmt.Println("  \"prompts\": [")
	for i, p := range prompts {
		comma := ","
		if i == len(prompts)-1 {
			comma = ""
		}
		fmt.Printf("    {\"filename\": \"%s\", \"theme\": \"%s\", \"style\": \"%s\", \"variant\": %d}%s\n",
			p.Filename, p.Theme, p.Style, p.Variant, comma)
	}
	fmt.Println("  ]")
	fmt.Println("}")
}

func printPromptsMarkdown(proj *config.Project, prompts []generator.PromptSpec, existing int) {
	fmt.Printf("# %s Logo Generation\n\n", proj.Project)
	fmt.Printf("**Tagline:** %s\n\n", proj.Tagline)
	fmt.Printf("| Metric | Value |\n")
	fmt.Printf("|--------|-------|\n")
	fmt.Printf("| Themes | %d |\n", len(proj.Themes))
	fmt.Printf("| Styles | %d |\n", len(proj.Styles))
	fmt.Printf("| Prompts | %d |\n", len(prompts))
	if existing > 0 {
		fmt.Printf("| Existing | %d |\n", existing)
	}
	fmt.Println("\n## Prompts\n")

	for i, p := range prompts {
		fmt.Printf("### %d. %s\n\n", i+1, p.Filename)
		fmt.Printf("```\n%s\n```\n\n", p.Prompt)
	}
}
