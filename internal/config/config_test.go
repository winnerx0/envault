package config

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_Save_andLoad(t *testing.T) {
	dir := t.TempDir()

	cfg := &Config{Name: "myproject", Version: "1.0"}
	if err := Save(dir, cfg); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded, err := Load(dir)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if loaded.Name != cfg.Name {
		t.Fatalf("expected name %q, got %q", cfg.Name, loaded.Name)
	}
	if loaded.Version != cfg.Version {
		t.Fatalf("expected version %q, got %q", cfg.Version, loaded.Version)
	}
}

func Test_Load_nonExistentFile(t *testing.T) {
	dir := t.TempDir()

	_, err := Load(dir)
	if err == nil {
		t.Fatal("expected error when loading non-existent config")
	}
}

func Test_Save_createsValidJSON(t *testing.T) {
	dir := t.TempDir()

	cfg := &Config{Name: "test", Version: "2.0"}
	Save(dir, cfg)

	data, err := os.ReadFile(filepath.Join(dir, FileName))
	if err != nil {
		t.Fatal(err)
	}

	content := string(data)
	if content[0] != '{' {
		t.Fatal("expected JSON object")
	}
}

func Test_FindProjectRoot_walksUp(t *testing.T) {
	root := t.TempDir()
	sub := filepath.Join(root, "src", "pkg")
	os.MkdirAll(sub, 0755)

	cfg := &Config{Name: "test", Version: "1.0"}
	Save(root, cfg)

	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(sub)

	found, err := FindProjectRoot()
	if err != nil {
		t.Fatalf("FindProjectRoot failed: %v", err)
	}
	if found != root {
		t.Fatalf("expected %q, got %q", root, found)
	}
}

func Test_FindProjectRoot_notFound(t *testing.T) {
	dir := t.TempDir()

	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(dir)

	_, err := FindProjectRoot()
	if err == nil {
		t.Fatal("expected error when no envault.json exists")
	}
}
