package main

import (
	"fmt"
	"log"
	_ "log"
	"os"

	"github.com/winnerx0/envault/internal/env"
	_ "github.com/winnerx0/envault/internal/utils"
)

func main() {

	// err := utils.OpenBrowser("http://localhost:8080")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	passphrase := "super-secret-passphrase" // do NOT hardcode in production
	inputFile := ".env"
	outputFile := ".env.enc"

	if err := env.EncryptFile(inputFile, outputFile, passphrase); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File encrypted to", outputFile)
	}

	bytes, err := os.ReadFile("./.env.enc")

	if err != nil {
		log.Fatal(err)
	}

	if err := env.DencryptFile(bytes, passphrase); err != nil {
		fmt.Println("Error:", err)
	}
}
