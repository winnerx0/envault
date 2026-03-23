package env

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/winnerx0/envault/internal/config"
)

func Test_DeriveKey_returns32bytes(t *testing.T) {
	salt := make([]byte, 16)
	io.ReadFull(rand.Reader, salt)

	key := DeriveKey("testpass", salt)
	if len(key) != 32 {
		t.Fatalf("expected 32-byte key, got %d", len(key))
	}
}

func Test_DeriveKey_deterministic(t *testing.T) {
	salt := make([]byte, 16)
	io.ReadFull(rand.Reader, salt)

	key1 := DeriveKey("testpass", salt)
	key2 := DeriveKey("testpass", salt)
	if string(key1) != string(key2) {
		t.Fatal("same passphrase and salt should produce the same key")
	}
}

func Test_DeriveKey_differentPassphrase(t *testing.T) {
	salt := make([]byte, 16)
	io.ReadFull(rand.Reader, salt)

	key1 := DeriveKey("testpass", salt)
	key2 := DeriveKey("otherpass", salt)
	if string(key1) == string(key2) {
		t.Fatal("different passphrases should produce different keys")
	}
}

func Test_DeriveKey_differentSalt(t *testing.T) {
	salt1 := make([]byte, 16)
	io.ReadFull(rand.Reader, salt1)
	salt2 := make([]byte, 16)
	io.ReadFull(rand.Reader, salt2)

	key1 := DeriveKey("testpass", salt1)
	key2 := DeriveKey("testpass", salt2)
	if string(key1) == string(key2) {
		t.Fatal("different salts should produce different keys")
	}
}

// helper to encrypt plaintext the same way Encrypt does
func encryptBytes(t *testing.T, plaintext, passphrase string) []byte {
	t.Helper()
	salt := make([]byte, 16)
	io.ReadFull(rand.Reader, salt)

	key := DeriveKey(passphrase, salt)
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)

	var buf []byte
	buf = append(buf, salt...)
	buf = append(buf, nonce...)
	buf = append(buf, ciphertext...)
	return buf
}

func Test_DencryptFile_roundTrip(t *testing.T) {
	passphrase := "my-secret-password"
	plaintext := "DB_HOST=localhost\nDB_PASS=hunter2\n"

	encrypted := encryptBytes(t, plaintext, passphrase)

	result, err := DencryptFile(encrypted, passphrase)
	if err != nil {
		t.Fatalf("decryption failed: %v", err)
	}
	if result != plaintext {
		t.Fatalf("expected %q, got %q", plaintext, result)
	}
}

func Test_DencryptFile_wrongPassphrase(t *testing.T) {
	encrypted := encryptBytes(t, "SECRET=value\n", "correct-password")

	_, err := DencryptFile(encrypted, "wrong-password")
	if err == nil {
		t.Fatal("expected error when decrypting with wrong passphrase")
	}
}

func Test_DencryptFile_emptyContent(t *testing.T) {
	encrypted := encryptBytes(t, "", "password")

	result, err := DencryptFile(encrypted, "password")
	if err != nil {
		t.Fatalf("decryption failed: %v", err)
	}
	if result != "" {
		t.Fatalf("expected empty string, got %q", result)
	}
}

func Test_DencryptFile_largeContent(t *testing.T) {
	plaintext := ""
	for i := 0; i < 1000; i++ {
		plaintext += "KEY_" + string(rune('A'+i%26)) + "=some_long_value_here\n"
	}

	encrypted := encryptBytes(t, plaintext, "password")

	result, err := DencryptFile(encrypted, "password")
	if err != nil {
		t.Fatalf("decryption failed: %v", err)
	}
	if result != plaintext {
		t.Fatal("decrypted content does not match original")
	}
}

func Test_Encrypt_findsEnvFiles(t *testing.T) {
	dir := t.TempDir()

	// Create envault.json so FindProjectRoot works
	cfg := &config.Config{Name: "testproject", Version: "1.0"}
	config.Save(dir, cfg)

	// Create .env files
	os.WriteFile(filepath.Join(dir, ".env"), []byte("KEY=value"), 0644)
	os.WriteFile(filepath.Join(dir, ".env.local"), []byte("LOCAL=true"), 0644)

	// Create a non-env file that should be ignored
	os.WriteFile(filepath.Join(dir, "main.go"), []byte("package main"), 0644)

	// Create nested .env
	sub := filepath.Join(dir, "config")
	os.MkdirAll(sub, 0755)
	os.WriteFile(filepath.Join(sub, ".env.production"), []byte("PROD=yes"), 0644)

	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(dir)

	results, err := Encrypt("testpass")
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	if len(results) != 3 {
		t.Fatalf("expected 3 encrypted files, got %d", len(results))
	}

	// Verify each can be decrypted back
	for _, ef := range results {
		_, err := DencryptFile(ef.Data, "testpass")
		if err != nil {
			t.Fatalf("failed to decrypt %s: %v", ef.Path, err)
		}
	}
}

func Test_Encrypt_noEnvFiles(t *testing.T) {
	dir := t.TempDir()

	cfg := &config.Config{Name: "empty", Version: "1.0"}
	config.Save(dir, cfg)

	os.WriteFile(filepath.Join(dir, "main.go"), []byte("package main"), 0644)

	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(dir)

	results, err := Encrypt("testpass")
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}
	if len(results) != 0 {
		t.Fatalf("expected 0 encrypted files, got %d", len(results))
	}
}

func Test_Encrypt_preservesRelativePaths(t *testing.T) {
	dir := t.TempDir()

	cfg := &config.Config{Name: "test", Version: "1.0"}
	config.Save(dir, cfg)

	sub := filepath.Join(dir, "nested", "deep")
	os.MkdirAll(sub, 0755)
	os.WriteFile(filepath.Join(sub, ".env"), []byte("DEEP=true"), 0644)

	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(dir)

	results, err := Encrypt("pass")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 file, got %d", len(results))
	}
	expected := filepath.Join("nested", "deep", ".env")
	if results[0].Path != expected {
		t.Fatalf("expected path %q, got %q", expected, results[0].Path)
	}
}
