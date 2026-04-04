package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string

var rootCmd = &cobra.Command{
	Use:   "envault",
	Short: "An env tool",
	Long:  "An environment variable syncing library",
	RunE: func(cmd *cobra.Command, args []string) error {
		
		if args[0] == "-v" || args[0] == "--version" {
			fmt.Println(version)
			return nil
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}