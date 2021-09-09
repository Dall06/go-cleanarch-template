package main

import (
	"fmt"

	"github.com/Dall06/go-cleanarch-template/src/infrastructure/app"
)

func main() {
	// INIT APP
	fmt.Println("Hello, Human!\nThis is the pssword-api-rest service")
	app.RunApp()
}
