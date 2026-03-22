package app

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const configFileName = "envault.json"

type Config struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize envault in the current directory",
	Long:  "Creates an envault.json config file in the current directory, marking it as the project root",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		configPath := filepath.Join(dir, configFileName)

		if _, err := os.Stat(configPath); err == nil {
			return fmt.Errorf("envault.json already exists in this directory")
		}

		config := Config{
			Name:    filepath.Base(dir),
			Version: "1.0",
		}

		data, err := json.MarshalIndent(config, "", "  ")
		
		if err != nil {
			return err
		}

		if err := os.WriteFile(configPath, data, 0644); err != nil {
			return err
		}

		fmt.Println("Initialized envault project in", dir)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
