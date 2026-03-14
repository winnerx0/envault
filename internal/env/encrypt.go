package env

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/argon2"
)

func DeriveKey(passphrase string, salt []byte) []byte {
	return argon2.IDKey([]byte(passphrase), salt, 3, 32*1024, 4, 32)
}

func EncryptFile(inputfile string, outputfile string, passphase string) error {

	plaintext, err := os.ReadFile(inputfile)

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

	fmt.Println("nonce", string(nonce))

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	outfile, err := os.Create(outputfile)

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

}