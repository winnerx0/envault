package global

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func withTempHome(t *testing.T, fn func(t *testing.T)) {
	t.Helper()
	tmpHome := t.TempDir()
	orig := os.Getenv("HOME")
	os.Setenv("HOME", tmpHome)
	t.Cleanup(func() { os.Setenv("HOME", orig) })
	fn(t)
}

func Test_Save_andLoad(t *testing.T) {
	withTempHome(t, func(t *testing.T) {
		cfg := &GlobalConfig{
			Token:      "ghp_testtoken123",
			PassPhrase: "mysecret",
			Repo:       "user/repo",
		}

		if err := Save(cfg); err != nil {
			t.Fatalf("Save failed: %v", err)
		}

		loaded, err := Load()
		if err != nil {
			t.Fatalf("Load failed: %v", err)
		}

		if loaded.Token != cfg.Token {
			t.Fatalf("expected token %q, got %q", cfg.Token, loaded.Token)
		}
		if loaded.PassPhrase != cfg.PassPhrase {
			t.Fatalf("expected passphrase %q, got %q", cfg.PassPhrase, loaded.PassPhrase)
		}
		if loaded.Repo != cfg.Repo {
			t.Fatalf("expected repo %q, got %q", cfg.Repo, loaded.Repo)
		}
	})
}

func Test_Load_returnsEmptyWhenMissing(t *testing.T) {
	withTempHome(t, func(t *testing.T) {
		cfg, err := Load()
		if err != nil {
			t.Fatalf("Load failed: %v", err)
		}
		if cfg.Token != "" || cfg.PassPhrase != "" || cfg.Repo != "" {
			t.Fatal("expected empty config when file doesn't exist")
		}
	})
}

func Test_Save_createsDirectory(t *testing.T) {
	withTempHome(t, func(t *testing.T) {
		cfg := &GlobalConfig{Token: "tok"}
		if err := Save(cfg); err != nil {
			t.Fatalf("Save failed: %v", err)
		}

		home := os.Getenv("HOME")
		dirPath := filepath.Join(home, configDir)
		info, err := os.Stat(dirPath)
		if err != nil {
			t.Fatalf("expected .envault dir to exist: %v", err)
		}
		if !info.IsDir() {
			t.Fatal("expected .envault to be a directory")
		}
	})
}

func Test_Save_filePermissions(t *testing.T) {
	withTempHome(t, func(t *testing.T) {
		cfg := &GlobalConfig{Token: "tok", PassPhrase: "pass"}
		Save(cfg)

		home := os.Getenv("HOME")
		filePath := filepath.Join(home, configDir, configFile)
		info, err := os.Stat(filePath)
		if err != nil {
			t.Fatal(err)
		}
		perm := info.Mode().Perm()
		if perm != 0600 {
			t.Fatalf("expected file permission 0600, got %o", perm)
		}
	})
}

func Test_Save_omitsEmptyRepo(t *testing.T) {
	withTempHome(t, func(t *testing.T) {
		cfg := &GlobalConfig{Token: "tok", PassPhrase: "pass"}
		Save(cfg)

		home := os.Getenv("HOME")
		data, _ := os.ReadFile(filepath.Join(home, configDir, configFile))
		if strings.Contains(string(data), "repo") {
			t.Fatal("expected repo field to be omitted when empty")
		}
	})
}
