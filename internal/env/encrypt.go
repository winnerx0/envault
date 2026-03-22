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

func EncryptFile(passphrase string) error {

	dir, err := config.FindProjectRoot()
	if err != nil {
		return err
	}

	regex := regexp.MustCompile(`^\.env.*`)

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

		outfile, err := os.OpenFile(path+".enc", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}

		outfile.Write(salt)
		outfile.Write(nonce)
		outfile.Write(ciphertext)

		outfile.Close()

		return nil
	})
	return err
}
