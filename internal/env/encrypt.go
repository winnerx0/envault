package env

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"golang.org/x/crypto/argon2"
)

func DeriveKey(passphrase string, salt []byte) []byte {
	return argon2.IDKey([]byte(passphrase), salt, 3, 32*1024, 4, 32)
}

func EncryptFile(inputfile string, outputfile string, passphase string) error {

	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	regex := regexp.MustCompile(".env.*")

	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return nil
		}

		if path == dir {
			return nil
		}

		
		if !regex.MatchString(d.Name()) {
			return nil
		}
		
		fmt.Println("name", path)

		plaintext, err := os.ReadFile(path)

		if err != nil {
			return err
		}

		salt := make([]byte, 16)

		if _, err := io.ReadFull(rand.Reader, salt); err != nil {
			return err
		}

		key := DeriveKey(passphase, salt)

		block, err := aes.NewCipher(key)

		if err != nil {
			return err
		}

		gcm, err := cipher.NewGCM(block)

		if err != nil {
			return err
		}

		nonce := make([]byte, gcm.NonceSize())

		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return err
		}

		ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

		outfile, err := os.Create(path + ".enc")

		if err != nil {
			return err
		}

		defer outfile.Close()

		if _, err := outfile.Write(salt); err != nil {
			return err
		}
		if _, err := outfile.Write(nonce); err != nil {
			return err
		}
		if _, err := outfile.Write(ciphertext); err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		return err
	}

	return nil

}
