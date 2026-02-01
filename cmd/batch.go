package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	batchParallel int
)

var batchCmd = &cobra.Command{
	Use:   "batch [projects...]",
	Short: "Generate logos for multiple projects",
	Long: `Run generate command for multiple projects in sequence or parallel.

If no projects specified, processes all projects in config directory.`,
	RunE: runBatch,
}

func init() {
	rootCmd.AddCommand(batchCmd)

	batchCmd.Flags().IntVarP(&batchParallel, "parallel", "p", 1, "number of parallel generations")
	batchCmd.Flags().IntVarP(&variants, "variants", "n", 1, "number of variants per combination")
	batchCmd.Flags().BoolVar(&dryRun, "dry-run", false, "show what would be generated")
	batchCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

func runBatch(cmd *cobra.Command, args []string) error {
	var projects []string

	if len(args) > 0 {
		projects = args
	} else {
		// Discover all projects in config dir
		projectsDir := filepath.Join(cfgDir, "projects")
		entries, err := os.ReadDir(projectsDir)
		if err != nil {
			return fmt.Errorf("failed to read projects directory: %w", err)
		}

		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".yaml") {
				name := strings.TrimSuffix(entry.Name(), ".yaml")
				projects = append(projects, name)
			}
		}
	}

	if len(projects) == 0 {
		return fmt.Errorf("no projects found in %s/projects/", cfgDir)
	}

	fmt.Printf("Batch processing %d projects: %v\n\n", len(projects), projects)

	for i, proj := range projects {
		fmt.Printf("━━━ [%d/%d] %s ━━━\n", i+1, len(projects), proj)

		// Re-use generate command logic
		err := runGenerate(cmd, []string{proj})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			// Continue with other projects
		}
		fmt.Println()
	}

	return nil
}
