package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string

var rootCmd = &cobra.Command{
	Use:   "envault <version>",
	Short: "An env tool",
	Long:  "An environment variable syncing library",
	RunE: func(cmd *cobra.Command, args []string) error {
		
		v, _ := cmd.Flags().GetBool("version")
		
		if v{
			fmt.Printf("v%s\n", version)
			return nil
		}
		return nil
	},
}

func init(){
	rootCmd.Flags().BoolP("version", "v", false, "Show version")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}