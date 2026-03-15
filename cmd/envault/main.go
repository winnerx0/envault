package main

import (
	_ "log"
	_ "os"

	"github.com/winnerx0/envault/app"
	_ "github.com/winnerx0/envault/internal/env"
)

func main() {

	// err := utils.OpenBrowser("http://localhost:8080")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// passphrase := "super-secret-passphrase" // do NOT hardcode in production
	// inputFile := ".env"
	// outputFile := ".env.enc"

	// if err := env.EncryptFile(inputFile, outputFile, passphrase); err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("File encrypted to", outputFile)
	// }

	// bytes, err := os.ReadFile("./.env.enc")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := env.DencryptFile(bytes, passphrase); err != nil {
	// 	fmt.Println("Error:", err)
	// }

	app.Execute()

}
