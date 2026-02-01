package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rickhallett/beautifi/internal/api"
	"github.com/rickhallett/beautifi/internal/config"
	"github.com/rickhallett/beautifi/internal/generator"
	"github.com/spf13/cobra"
)

var (
	variants   int
	styles     []string
	dryRun     bool
	verbose    bool
	promptOnly bool
)

var generateCmd = &cobra.Command{
	Use:   "generate <project>",
	Short: "Generate logos for a project",
	Long: `Generate logos using AI image generation.
	
Reads project config from ~/.config/beautifi/projects/<project>.yaml
Generates images to ~/output/beautifi/<project>/`,
	Args: cobra.ExactArgs(1),
	RunE: runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntVarP(&variants, "variants", "n", 1, "number of variants per combination")
	generateCmd.Flags().StringSliceVarP(&styles, "styles", "s", []string{"all"}, "styles to generate (all, flat, gradient, etc.)")
	generateCmd.Flags().BoolVar(&dryRun, "dry-run", false, "show what would be generated without calling API")
	generateCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	generateCmd.Flags().BoolVar(&promptOnly, "prompts-only", false, "only output prompts, no images")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	// Load project config
	cfgPath := filepath.Join(cfgDir, "projects", projectName+".yaml")
	proj, err := config.LoadProject(cfgPath)
	if err != nil {
		return fmt.Errorf("failed to load project config: %w\n\nCreate config at: %s", err, cfgPath)
	}

	if verbose {
		fmt.Printf("Project: %s\n", proj.Project)
		fmt.Printf("Tagline: %s\n", proj.Tagline)
		fmt.Printf("Themes: %v\n", proj.Themes)
		fmt.Printf("Styles: %v\n", proj.Styles)
		fmt.Println()
	}

	// Filter styles if specified
	activeStyles := proj.Styles
	if len(styles) > 0 && styles[0] != "all" {
		activeStyles = filterStyles(proj.Styles, styles)
	}

	// Generate prompts
	prompts := generator.GeneratePrompts(proj, activeStyles, variants)

	if verbose || dryRun || promptOnly {
		fmt.Printf("Generated %d prompts:\n\n", len(prompts))
		for i, p := range prompts {
			fmt.Printf("[%d] %s\n", i+1, p.Filename)
			fmt.Printf("    Theme: %s, Style: %s, Variant: %d\n", p.Theme, p.Style, p.Variant)
			fmt.Printf("    Prompt: %s\n\n", truncate(p.Prompt, 100))
		}
	}

	if dryRun || promptOnly {
		fmt.Printf("Dry run complete. Would generate %d images.\n", len(prompts))
		return nil
	}

	// Check API key
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("GEMINI_API_KEY environment variable not set")
	}

	// Create output directory
	projectOutDir := filepath.Join(outDir, proj.Project)
	if err := os.MkdirAll(projectOutDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Generate images
	client, err := api.NewGeminiClient(apiKey)
	if err != nil {
		return fmt.Errorf("failed to create Gemini client: %w", err)
	}
	defer client.Close()
	
	results, err := generator.GenerateImages(client, prompts, projectOutDir, verbose)
	if err != nil {
		return fmt.Errorf("generation failed: %w", err)
	}

	// Summary
	success := 0
	for _, r := range results {
		if r.Success {
			success++
		}
	}
	fmt.Printf("\nComplete: %d/%d images generated\n", success, len(results))
	fmt.Printf("Output: %s\n", projectOutDir)

	return nil
}

func filterStyles(available, requested []string) []string {
	requestMap := make(map[string]bool)
	for _, s := range requested {
		requestMap[s] = true
	}
	var filtered []string
	for _, s := range available {
		if requestMap[s] {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
