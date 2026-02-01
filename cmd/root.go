package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
	cfgDir  string
	outDir  string
)

var rootCmd = &cobra.Command{
	Use:     "beautifi",
	Short:   "Batch logo generation CLI",
	Long:    `beautifi v` + version + ` — Generate logos and icons using AI image generation.`,
	Version: version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	home, _ := os.UserHomeDir()
	defaultCfg := home + "/.config/beautifi"
	defaultOut := home + "/output/beautifi"

	rootCmd.PersistentFlags().StringVar(&cfgDir, "config-dir", defaultCfg, "config directory")
	rootCmd.PersistentFlags().StringVar(&outDir, "output-dir", defaultOut, "output directory")
}
