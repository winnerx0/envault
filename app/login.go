package app

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/winnerx0/envault/internal/config"
	"github.com/winnerx0/envault/internal/env"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to envault",
	Long:  "Login to envault",
	RunE: func(cmd *cobra.Command, args []string) error {
		password, err := cmd.Flags().GetString("password")
		
		if err != nil {
			return err
		}

		root, err := config.FindProjectRoot()
		if err != nil {
			return err
		}
		fmt.Println("Project root:", filepath.Base(root))

		env.EncryptFile(password)
		return nil
	},
}

func init(){

	loginCmd.Flags().String("password", "", "Create a password for envault")
	rootCmd.AddCommand(loginCmd)
}