package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/winnerx0/envault/internal/config"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize envault in the current directory",
	Long:  "Creates an envault.json config file in the current directory, marking it as the project root",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		configPath := filepath.Join(dir, config.FileName)

		if _, err := os.Stat(configPath); err == nil {
			return fmt.Errorf("envault.json already exists in this directory")
		}

		cfg := &config.Config{
			Name:    filepath.Base(dir),
			Version: "1.0",
		}

		if err := config.Save(dir, cfg); err != nil {
			return err
		}

		fmt.Println("Initialized envault project in", dir)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
