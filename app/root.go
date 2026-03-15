package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "envault",
	Short: "An env tool",
	Long:  "env syncing",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello from root command")
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
