package app

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/winnerx0/envault/internal/config"
	"github.com/winnerx0/envault/internal/env"
	"github.com/winnerx0/envault/internal/global"
)

var httpClient = http.Client{}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup env files",
	Long:  "Backup encrypted environment variable files to a private github repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		
		root, err := config.FindProjectRoot()
		
		if err != nil {
			return err
		}
		
		cfg, err := config.Load(root)
		
		if err != nil {
			return err
		}

		gcfg, err := global.Load()
		if err != nil {
			return err
		}
		if gcfg.Token == "" || gcfg.Repo == "" {
			return errors.New("not logged in — run 'envault login --token <TOKEN>' first")
		}
		token := gcfg.Token
		
		
		encrypted, err := env.Encrypt(gcfg.PassPhrase)
		if err != nil {
			return err
		}
		if len(encrypted) == 0 {
			return errors.New("no .env files found")
		}

		fullRepo := gcfg.Repo

		type treeEntry struct {
			Path string `json:"path"`
			Mode string `json:"mode"`
			Type string `json:"type"`
			SHA  string `json:"sha"`
		}
		var treeEntries []treeEntry

		for _, ef := range encrypted {
			sha, err := createBlob(token, fullRepo, ef.Data)
			if err != nil {
				return err
			}

			treeEntries = append(treeEntries, treeEntry{
				Path: filepath.Join(filepath.Base(root), ef.Path+".enc"),
				Mode: "100644",
				Type: "blob",
				SHA:  sha,
			})
		}

		// Get latest commit and its tree
		parentSHA, err := getMainRef(token, fullRepo)
		if err != nil {
			return err
		}

		baseTreeSHA, err := getCommitTreeSHA(token, fullRepo, parentSHA)
		if err != nil {
			return err
		}

		// Create tree with base_tree so other project folders are preserved
		treeSHA, err := createTreeWithBase(token, fullRepo, baseTreeSHA, treeEntries)
		if err != nil {
			return err
		}

		// Create commit with parent
		commitSHA, err := createCommit(token, fullRepo, treeSHA, parentSHA, fmt.Sprintf("backup encrypted env files for %s", cfg.Name))
		if err != nil {
			return err
		}

		// Update main branch ref
		if err := updateRef(token, fullRepo, commitSHA); err != nil {
			return err
		}

		fmt.Println("Encrypted env files backed up to private repo:", fullRepo)
		return nil
	},
}

func githubRequest(token, method, url string, body interface{}) ([]byte, error) {
	var buf io.Reader
	if body != nil {
		data, _ := json.Marshal(body)
		buf = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("github API error (%d): %s", resp.StatusCode, respBody)
	}

	return respBody, nil
}

func createRepo(token, name string) (string, error) {
	body := map[string]interface{}{
		"name":      name,
		"private":   true,
		"auto_init": true,
	}

	respBody, err := githubRequest(token, http.MethodPost, "https://api.github.com/user/repos", body)
	if err != nil {
		return "", err
	}

	var resp struct {
		Owner struct {
			Login string `json:"login"`
		} `json:"owner"`
	}
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return "", err
	}

	return resp.Owner.Login, nil
}

func createBlob(token, repo string, content []byte) (string, error) {
	body := map[string]string{
		"content":  base64.StdEncoding.EncodeToString(content),
		"encoding": "base64",
	}

	respBody, err := githubRequest(token, http.MethodPost,
		fmt.Sprintf("https://api.github.com/repos/%s/git/blobs", repo), body)
	if err != nil {
		return "", err
	}

	var resp struct {
		SHA string `json:"sha"`
	}
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return "", err
	}

	return resp.SHA, nil
}

func createTreeWithBase(token, repo, baseTreeSHA string, entries interface{}) (string, error) {
	body := map[string]interface{}{
		"base_tree": baseTreeSHA,
		"tree":      entries,
	}

	respBody, err := githubRequest(token, http.MethodPost,
		fmt.Sprintf("https://api.github.com/repos/%s/git/trees", repo), body)
	if err != nil {
		return "", err
	}

	var resp struct {
		SHA string `json:"sha"`
	}
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return "", err
	}

	return resp.SHA, nil
}

func getCommitTreeSHA(token, repo, commitSHA string) (string, error) {
	respBody, err := githubRequest(token, http.MethodGet,
		fmt.Sprintf("https://api.github.com/repos/%s/git/commits/%s", repo, commitSHA), nil)
	if err != nil {
		return "", err
	}

	var resp struct {
		Tree struct {
			SHA string `json:"sha"`
		} `json:"tree"`
	}
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return "", err
	}

	return resp.Tree.SHA, nil
}

func getMainRef(token, repo string) (string, error) {
	respBody, err := githubRequest(token, http.MethodGet,
		fmt.Sprintf("https://api.github.com/repos/%s/git/ref/heads/main", repo), nil)
	if err != nil {
		return "", err
	}

	var resp struct {
		Object struct {
			SHA string `json:"sha"`
		} `json:"object"`
	}
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return "", err
	}

	return resp.Object.SHA, nil
}

func createCommit(token, repo, treeSHA, parentSHA, message string) (string, error) {
	body := map[string]interface{}{
		"message": message,
		"tree":    treeSHA,
		"parents": []string{parentSHA},
	}

	respBody, err := githubRequest(token, http.MethodPost,
		fmt.Sprintf("https://api.github.com/repos/%s/git/commits", repo), body)
	if err != nil {
		return "", err
	}

	var resp struct {
		SHA string `json:"sha"`
	}
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return "", err
	}

	return resp.SHA, nil
}

func updateRef(token, repo, commitSHA string) error {
	body := map[string]string{
		"sha": commitSHA,
	}

	_, err := githubRequest(token, http.MethodPatch,
		fmt.Sprintf("https://api.github.com/repos/%s/git/refs/heads/main", repo), body)
	return err
}

func randomString() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
