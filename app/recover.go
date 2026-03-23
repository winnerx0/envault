package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/winnerx0/envault/internal/config"
	"github.com/winnerx0/envault/internal/env"
	"github.com/winnerx0/envault/internal/global"
)

type Entry struct {
	Path         string `json:"path"`
	Name         string `json:"name"`
	Download_url string `json:"download_url"`
	ContentType  string `json:"type"`
}

type GithubError struct {
	Message string `json:"message"`
}

type ResponseBody struct {
	Entries []Entry `json:"entries"`
}

var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "Recover encrypted env files",
	RunE: func(cmd *cobra.Command, args []string) error {

		root, err := config.FindProjectRoot()

		if err != nil {
			return err
		}

		gcfg, err := global.Load()

		if err != nil {
			return err
		}

		cfg, err := config.Load(root)

		if err != nil {
			return err
		}

		entries, err := fetchEntries(gcfg.Repo, cfg.Name, gcfg.Token)

		if err != nil {
			return err
		}

		fmt.Println("Environment variables fully recovered")

		for _, entry := range entries {
			resp, err := http.Get(entry.Download_url)
			if err != nil {
				return err
			}

			relPath, err := filepath.Rel("envault", entry.Path)

			if err != nil {
				return err
			}

			out, err := os.Create(strings.Replace(relPath, ".enc", "", 1))

			if err != nil {
				return err
			}

			defer out.Close()

			respBody, err := io.ReadAll(resp.Body)

			if err != nil {
				return err
			}

			defer resp.Body.Close()

			plainText, err := env.DencryptFile(respBody, gcfg.PassPhrase)

			if err != nil {
				return err
			}

			_, err = io.Copy(out, strings.NewReader(plainText))

			if err != nil {
				return err
			}

		}
		return nil
	},
}

func fetchEntries(repo string, path string, token string) ([]Entry, error) {

	var respBody ResponseBody

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repo, path), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")
	req.Header.Set("Accept", "application/vnd.github.object")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var gitubError GithubError
		err = json.NewDecoder(resp.Body).Decode(&gitubError)
		return nil, errors.New(gitubError.Message)
	}

	err = json.NewDecoder(resp.Body).Decode(&respBody)

	if err != nil {
		return nil, err
	}

	var entries []Entry

	for _, entry := range respBody.Entries {

		if entry.ContentType == "dir" {

			nestedEntries, err := fetchEntries(repo, entry.Path, token)

			if err != nil {
				return nil, err
			}

			entries = append(entries, nestedEntries...)

		} else {
			entries = append(entries, entry)
		}

	}

	return entries, nil
}

func init() {
	rootCmd.AddCommand(recoverCmd)
}
