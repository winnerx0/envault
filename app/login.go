package app

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/winnerx0/envault/internal/env"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to envault",
	Long:  "Login to envault",
	RunE: func(cmd *cobra.Command, args []string) error {
		password, _ := cmd.Flags().GetString("password")

		fmt.Println("password", password)

		env.EncryptFile(".env", ".env.env", password)
		return nil
	},
}

func init(){

	loginCmd.Flags().String("password", "", "Create a password for envault")
	rootCmd.AddCommand(loginCmd)
}