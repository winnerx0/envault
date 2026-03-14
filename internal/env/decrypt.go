package env

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func DencryptFile(encrypted []byte, passphase string) error {

	saltSize := 16

	salt := encrypted[:saltSize]

	key := DeriveKey(passphase, salt)

	block, err := aes.NewCipher(key)

	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()

	nonce := encrypted[saltSize : saltSize+nonceSize]

	cipherText := encrypted[saltSize+nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)

	if err != nil {
		return err
	}

	fmt.Println(string(plaintext))

	return nil
}
