package main

import (
	_ "log"
	_ "os"

	"github.com/winnerx0/envault/app"
	_ "github.com/winnerx0/envault/internal/env"
)

func main() {

	app.Execute()

}
