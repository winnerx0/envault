package app

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/winnerx0/envault/internal/config"
	"github.com/winnerx0/envault/internal/global"
)

var password string
var token string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to envault",
	Long:  "Login to envault",
	RunE: func(cmd *cobra.Command, args []string) error {

		root, err := config.FindProjectRoot()
		if err != nil {
			return err
		}
		fmt.Println("Project root:", filepath.Base(root))

		if token != "" {
			gcfg := &global.GlobalConfig{Token: token, PassPhrase: password}

			// Create a private repo if one doesn't exist yet
			repoName := randomString()
			owner, err := createRepo(token, repoName)
			if err != nil {
				return err
			}
			gcfg.Repo = owner + "/" + repoName
			fmt.Println("Created private backup repo:", gcfg.Repo)

			if err := global.Save(gcfg); err != nil {
				return err
			}
			fmt.Println("Config saved to ~/.envault/config.yaml")
		}
		return nil
	},
}

func init() {
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Create a password for envault")
	loginCmd.MarkFlagRequired("password")
	loginCmd.Flags().StringVarP(&token, "token", "t", "", "GitHub personal access token for backups")
	rootCmd.AddCommand(loginCmd)
}