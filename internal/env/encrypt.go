package env

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"github.com/winnerx0/envault/internal/config"
	"golang.org/x/crypto/argon2"
)

func DeriveKey(passphrase string, salt []byte) []byte {
	return argon2.IDKey([]byte(passphrase), salt, 3, 32*1024, 4, 32)
}

type EncryptedFile struct {
	Path string // relative path from project root
	Data []byte // encrypted content (salt + nonce + ciphertext)
}

func Encrypt(passphrase string) ([]EncryptedFile, error) {
	dir, err := config.FindProjectRoot()
	if err != nil {
		return nil, err
	}

	regex := regexp.MustCompile(`^\.env.*`)
	var results []EncryptedFile

	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !regex.MatchString(d.Name()) {
			return nil
		}

		plaintext, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		salt := make([]byte, 16)
		io.ReadFull(rand.Reader, salt)

		key := DeriveKey(passphrase, salt)

		block, err := aes.NewCipher(key)
		if err != nil {
			return err
		}

		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return err
		}

		nonce := make([]byte, gcm.NonceSize())
		io.ReadFull(rand.Reader, nonce)

		ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

		var buf []byte
		buf = append(buf, salt...)
		buf = append(buf, nonce...)
		buf = append(buf, ciphertext...)

		rel, _ := filepath.Rel(dir, path)
		results = append(results, EncryptedFile{
			Path: rel,
			Data: buf,
		})

		return nil
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}
