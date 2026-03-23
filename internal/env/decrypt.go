package env

import (
	"crypto/aes"
	"crypto/cipher"
)

func DencryptFile(encrypted []byte, passphrase string) (string, error) {

	saltSize := 16

	salt := encrypted[:saltSize]

	key := DeriveKey(passphrase, salt)

	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()

	nonce := encrypted[saltSize : saltSize+nonceSize]

	cipherText := encrypted[saltSize+nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)

	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
